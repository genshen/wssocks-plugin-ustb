// this file provide api for launching and stopping client

package extra

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/genshen/wssocks-plugin-ustb/plugins/ver"
	"github.com/genshen/wssocks-plugin-ustb/plugins/vpn"
	"github.com/genshen/wssocks/client"
)

type Options struct {
	client.Options
	vpn.UstbVpn
	RemoteAddr string
}

type TaskHandles struct {
	client.Handles
	once *sync.Once
}

func (h *TaskHandles) NotifyCloseWrapper() {
	h.NotifyClose(h.once, false)
}

var pluginRegistered = false

func (h *TaskHandles) StartWssocks(options Options) error {
	if !pluginRegistered {
		client.AddPluginRequest(&options.UstbVpn)
		client.AddPluginVersion(&ver.PluginVersionNeg{})
		pluginRegistered = true
	}

	// check remote url
	if options.RemoteAddr == "" {
		return errors.New("empty remote address")
	}
	u, err := url.Parse(options.RemoteAddr)
	if err != nil {
		return err
	} else {
		options.RemoteUrl = u
	}

	options.RemoteHeaders = make(http.Header)

	h.Handles = *client.NewClientHandles()
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute) // fixme
	defer cancel()

	_, err = h.CreateServerConn(&options.Options, ctx)
	if err != nil {
		return err
	}
	// server connect successfully

	if err := h.NegotiateVersion(ctx, options.RemoteAddr); err != nil {
		return err
	}

	var once sync.Once
	h.once = &once
	h.StartClient(&options.Options, &once)
	return nil
}
