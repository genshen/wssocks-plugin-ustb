package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

type HyperlinkIcon struct {
	widget.Icon
	URL       *url.URL
}

func NewHyperlinkIcon(res fyne.Resource, url *url.URL) *HyperlinkIcon {
	hlIcon := &HyperlinkIcon{URL: url}
	hlIcon.ExtendBaseWidget(hlIcon)
	hlIcon.SetResource(res)
	return hlIcon
}

func (hl *HyperlinkIcon) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}

func (hl *HyperlinkIcon) Tapped(*fyne.PointEvent) {
	if hl.URL != nil {
		err := fyne.CurrentApp().OpenURL(hl.URL)
		if err != nil {
			fyne.LogError("Failed to open url", err)
		}
	}
}
