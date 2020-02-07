package main

import (
	"errors"
	"fmt"
	"github.com/genshen/wssocks-plugin-ustb/plugins/vpn"
	"github.com/genshen/wssocks/wss"
	"github.com/genshen/wssocks/wss/term_view"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
)

type Options struct {
	vpn.UstbVpn
	remoteUrl       *url.URL
	localSocks5Addr string
	remoteAddr      string
	httpEnable      bool
	localHttpAddr   string
}

func startWssocks(o Options) error {
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
		return fmt.Errorf("establishing connection error: %s", err.Error())
	}

	// todo chan for wsc and tcp accept
	//defer wsc.WSClose()

	if _, err := wss.ExchangeVersion(wsc.WsConn); err != nil {
		return err
	}

	// start websocket message listen.
	go func() {
		if err := wsc.ListenIncomeMsg(); err != nil {
			log.Println("error websocket read:", err)
		}
	}()
	// send heart beats.
	go func() {
		if err := wsc.HeartBeat(); err != nil {
			log.Println("heartbeat ending", err)
		}
	}()

	if o.httpEnable {
		go func() {
			handle := wss.NewHttpProxy(wsc, nil)
			if err := http.ListenAndServe(o.localHttpAddr, &handle); err != nil {
				log.Fatalln(err)
			}
		}()
	}

	plog := term_view.NewPLog()
	// start listen for socks5 and https connection.
	if err := wss.ListenAndServe(plog, wsc, o.localSocks5Addr, o.httpEnable, func() {

	}); err != nil {
		return err
	}
	return nil
}
