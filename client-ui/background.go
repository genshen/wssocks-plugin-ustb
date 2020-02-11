package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/genshen/wssocks-plugin-ustb/plugins/vpn"
	"github.com/genshen/wssocks/wss"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"sync"
)

type Options struct {
	vpn.UstbVpn
	remoteUrl       *url.URL
	localSocks5Addr string
	remoteAddr      string
	httpEnable      bool
	localHttpAddr   string
}

type Handles struct {
	wsc        *wss.WebSocketClient
	hb         *wss.HeartBeat
	httpServer *http.Server
	cl         *wss.Client
	closed     bool
	wg         sync.WaitGroup
}

func (h *Handles) Close() {
	if h.closed {
		return
	}
	h.closed = true
	h.cl.Close(false)
	if h.httpServer != nil {
		h.httpServer.Shutdown(context.TODO())
	}
	if h.hb != nil {
		h.hb.Close()
	}
	if h.wsc != nil {
		h.wsc.Close()
	}
	h.wg.Wait() // wait tasks finishing
}

func (h *Handles) startWssocks(o Options) error {
	// check remote url
	if o.remoteAddr == "" {
		return errors.New("empty remote address")
	}
	if u, err := url.Parse(o.remoteAddr); err != nil {
		return err
	} else {
		o.remoteUrl = u
	}

	dialer := websocket.DefaultDialer
	wsHeader := make(http.Header)

	// we don't register redirect plugin, just call vpn plugin directly.
	if err := o.UstbVpn.BeforeRequest(dialer, o.remoteUrl, wsHeader); err != nil {
		return err
	}

	wsc, err := wss.NewWebSocketClient(websocket.DefaultDialer, o.remoteUrl.String(), wsHeader)
	if err != nil {
		return fmt.Errorf("establishing connection error: %w", err)
	}

	// todo chan for wsc and tcp accept

	if _, err := wss.ExchangeVersion(wsc.WsConn); err != nil {
		return err
	}

	h.wsc = wsc
	// start websocket message listen.
	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		if err := wsc.ListenIncomeMsg(); err != nil {
			log.Println("error websocket read:", err)
		}
	}()
	// send heart beats.
	h.hb = wss.NewHeartBeat(wsc)
	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		if err := h.hb.Start(); err != nil {
			log.Println("heartbeat ending", err)
		}
	}()

	record := wss.NewConnRecord() // connection record
	if o.httpEnable {
		h.wg.Add(1)
		go func() {
			defer h.wg.Done()
			handle := wss.NewHttpProxy(wsc, record)
			h.httpServer = &http.Server{Addr: o.localHttpAddr, Handler: &handle}
			if err := h.httpServer.ListenAndServe(); err != nil {
				log.Println(err)
			}
		}()
	}

	// start listen for socks5 and https connection.
	h.cl = wss.NewClient()
	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		if err := h.cl.ListenAndServe(record, wsc, o.localSocks5Addr, o.httpEnable, func() {
		}); err != nil {
			log.Println(err)
		}
	}()
	h.closed = false // reopen
	return nil
}
