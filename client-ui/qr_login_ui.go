package main

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
	"github.com/genshen/wssocks-plugin-ustb/plugins/vpn"
	log "github.com/sirupsen/logrus"
)

// LoadQRImage shows QR image for login
func LoadQRImage() *canvas.Image {
	if qrCodeImgUrl, err := vpn.ParseQRCodeImgUrl(); err != nil {
		log.Println(err) // todo
		return nil
	} else {
		if uri, err := storage.ParseURI(qrCodeImgUrl); err != nil {
			log.Println(err) // todo:
			return nil
		} else {
			image := canvas.NewImageFromURI(uri)
			image.FillMode = canvas.ImageFillOriginal
			return image
		}
	}
}
