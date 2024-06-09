package main

import (
	"fyne.io/fyne/v2/canvas"
	"github.com/genshen/wssocks-plugin-ustb/plugins/vpn"
	log "github.com/sirupsen/logrus"
)

// LoadQRImage shows QR image for login
func LoadQRImage() *canvas.Image {
	if qrImgReader, err := vpn.LoadQrAuthImage(); err != nil {
		log.Println(err) // todo
		return nil
	} else {
		image := canvas.NewImageFromReader(qrImgReader, "qr.png")
		image.FillMode = canvas.ImageFillOriginal
		return image
	}
}
