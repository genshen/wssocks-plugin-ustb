package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	qrcode2 "github.com/genshen/wssocks-plugin-ustb/plugins/vpn/qrcode"
	"github.com/skip2/go-qrcode"
	"net/http"
	"time"
)

type FyneQrCodeAuth struct {
	appRef *fyne.App
}

func newQrCodeAuth(app *fyne.App) qrcode2.QrCodeAuth {
	return &FyneQrCodeAuth{
		appRef: app,
	}
}

func (q *FyneQrCodeAuth) ShowQrCodeAndWait(client *http.Client, cookies []*http.Cookie, qr qrcode2.QrImg) ([]*http.Cookie, error) {
	// generate qr code from image
	qrPng, err := qrcode.Encode(qr.GenQrCodeContent(), qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewReader(qrPng)
	QrImage := canvas.NewImageFromReader(buf, "qr.png")
	QrImage.FillMode = canvas.ImageFillOriginal

	scanned := make(chan bool, 1) // signal of qr code scan finished
	// show qr code window
	qrAuthWindow := (*q.appRef).NewWindow("QR Code vpn auth")
	qrAuthWindow.SetContent(container.NewVBox(
		QrImage,
		widget.NewLabel("scan QR code, and then click button `Finish` "),
		widget.NewButton("Finish", func() {
			scanned <- true
		}),
	))
	qrAuthWindow.Show()

	// wait qr code scanned or time out
	// scan the qr code in 30 seconds. Otherwise, an error of Timeout will return.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	select {
	case <-scanned:
		qrAuthWindow.Close()
		return WaitStatus(client, cookies, qr)
	case <-ctx.Done():
		qrAuthWindow.Close()
		return nil, errors.New("scan QR code canceled due to timeout")
	}
}

func WaitStatus(client *http.Client, cookies []*http.Cookie, qr qrcode2.QrImg) ([]*http.Cookie, error) {
	// todo: set http cancel, after timeout
	if state, err := qrcode2.WaitQrState(qr.Sid); err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		if err = qrcode2.RedirectToLogin(client, cookies, qr.Config.AppID, state, qr.Config.RandToken); err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	return nil, nil
}
