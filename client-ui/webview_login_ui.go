package main

import (
	"errors"
	"fyne.io/fyne/v2"
	"net/http"

	"github.com/genshen/wssocks-plugin-ustb/plugins/vpn/webview"
)

func NewWebviewAuth(app *fyne.App, chromePathInSettings string) webview.WebviewAuth {
	return &FyneWebviewAuth{
		appRef:         app,
		chromeProxy:    &webview.ChromedpWebview{},
		chromePathHint: chromePathInSettings,
	}
}

type FyneWebviewAuth struct {
	appRef         *fyne.App
	chromeProxy    *webview.ChromedpWebview
	chromePathHint string
}

func (w *FyneWebviewAuth) GetCookie(client *http.Client, loginUrl string) ([]*http.Cookie, error) {
	if w.chromeProxy == nil {
		return nil, errors.New("Chromedp is not created")
	}

	// created ui:
	return w.chromeProxy.ShowWebviewAndSetCookies(client, loginUrl, w.chromePathHint)
}

func (w *FyneWebviewAuth) WaitAuthFinished() error {
	return nil
}
