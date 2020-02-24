package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/genshen/wssocks-plugin-ustb/plugins/vpn"
	"github.com/genshen/wssocks/version"
	"net/url"
)

const (
	AppName       = "wssocks Client"
	AppId         = "wssocks-ustb.genshen.github.com"
	GithubRepoUrl = "https://github.com/genshen/wssocks"
	DocumentUrl   = "https://github.com/genshen/wssocks-plugin-ustb/blob/master/docs/zh-cn/README.md"
)

const (
	btnStopped = iota
	btnStarting
	btnRunning
	btnStopping
)

func main() {
	wssApp := app.NewWithID(AppId)
	wssApp.Settings().SetTheme(theme.LightTheme())

	w := wssApp.NewWindow(AppName)
	//w.SetFixedSize(true)
	//w.Resize(fyne.NewSize(100, 100))

	// basic input
	uiLocalAddr := &widget.Entry{Text: "127.0.0.1:1080"}
	uiRemoteAddr := &widget.Entry{Text: "ws://proxy.gensh.me"}
	uiHttpEnable := &widget.Check{Checked: false}
	uiHttpLocalAddr := &widget.Entry{Text: "127.0.0.1:1086"}

	loadBasicPreference(wssApp.Preferences(), uiLocalAddr, uiRemoteAddr, uiHttpLocalAddr, uiHttpEnable)

	uiHttpEnable.OnChanged = func(checked bool) {
		if checked {
			uiHttpLocalAddr.Enable()
		} else {
			uiHttpLocalAddr.Disable()
		}
	}

	// vpn input
	uiVpnEnable := &widget.Check{Text: "enable ustb vpn", Checked: true}
	uiVpnForceLogout := &widget.Check{Text: "", Checked: true}
	uiVpnHostEncrypt := &widget.Check{Text: "", Checked: true}
	uiVpnHostInput := &widget.Entry{Text: "vpn4.ustb.edu.cn"}
	uiVpnUsername := &widget.Entry{Text: ""}
	uiVpnPassword := &widget.Entry{Text: "", Password: true}

	loadVPNPreference(wssApp.Preferences(), uiVpnEnable, uiVpnForceLogout,
		uiVpnHostEncrypt, uiVpnHostInput, uiVpnUsername, uiVpnPassword)

	uiVpnEnable.OnChanged = func(checked bool) {
		if checked {
			uiVpnForceLogout.Enable()
			uiVpnHostEncrypt.Enable()
			uiVpnHostInput.Enable()
			uiVpnUsername.Enable()
			uiVpnPassword.Enable()
		} else {
			uiVpnForceLogout.Disable()
			uiVpnHostEncrypt.Disable()
			uiVpnHostInput.Disable()
			uiVpnUsername.Disable()
			uiVpnPassword.Disable()
		}
	}

	btnStart := widget.NewButton("Start", nil)
	btnStatus := btnStopped
	var handles Handles
	btnStart.OnTapped = func() {
		if btnStatus == btnRunning { // running can stop
			btnStatus = btnStopping
			btnStart.SetText("Stopping")
			handles.Close()
			btnStart.SetText("Start")
			btnStatus = btnStopped
		} else if btnStatus == btnStopped { // stopped can run
			options := Options{
				localSocks5Addr: uiLocalAddr.Text,
				remoteAddr:      uiRemoteAddr.Text,
				httpEnable:      uiHttpEnable.Checked,
				localHttpAddr:   uiHttpLocalAddr.Text,
				UstbVpn: vpn.UstbVpn{
					Enable:      uiVpnEnable.Checked,
					ForceLogout: uiVpnForceLogout.Checked,
					HostEncrypt: uiVpnHostEncrypt.Checked,
					TargetVpn:   uiVpnHostInput.Text,
					Username:    uiVpnUsername.Text,
					Password:    uiVpnPassword.Text,
				},
			}
			btnStatus = btnStarting
			btnStart.SetText("Loading")
			if err := handles.startWssocks(options); err != nil {
				// log error
				dialog.ShowError(err, w)
				btnStart.SetText("Start")
				btnStatus = btnStopped
				return
			}
			btnStart.SetText("Stop")
			btnStatus = btnRunning
		}
	}

	docUrl, err := url.Parse(DocumentUrl)
	if err != nil {
		return
	}

	repoUrl, err := url.Parse(GithubRepoUrl)
	if err != nil {
		return
	}

	w.SetContent(widget.NewVBox(
		widget.NewGroup("basic",
			&widget.Form{Items: []*widget.FormItem{
				{Text: "socks5 address", Widget: uiLocalAddr},
				{Text: "remote address", Widget: uiRemoteAddr},
				{Text: "http(s) proxy", Widget: uiHttpEnable},
				{Text: "http(s) address", Widget: uiHttpLocalAddr},
			}},
		), // end group
		widget.NewGroup("vpn",
			&widget.Form{Items: []*widget.FormItem{
				{Text: "enable", Widget: uiVpnEnable},
				{Text: "force logout", Widget: uiVpnForceLogout},
				{Text: "host encrypt", Widget: uiVpnHostEncrypt},
				{Text: "vpn host", Widget: uiVpnHostInput},
				{Text: "username", Widget: uiVpnUsername},
				{Text: "password", Widget: uiVpnPassword},
			}},
		), // end group
		btnStart,
		widget.NewHBox(
			widget.NewLabel("v"+version.VERSION),
			widget.NewHyperlink("Github", repoUrl),
			widget.NewHyperlink("Document", docUrl),
		),
	))

	w.SetOnClosed(func() {
		// todo close all and stop if network lost
		if btnStatus == btnRunning { // running can stop
			btnStatus = btnStopping
			btnStart.SetText("Stopping")
			handles.Close()
		}
		saveBasicPreference(wssApp.Preferences(), uiLocalAddr, uiRemoteAddr, uiHttpLocalAddr, uiHttpEnable)
		saveVPNPreference(wssApp.Preferences(), uiVpnEnable, uiVpnForceLogout,
			uiVpnHostEncrypt, uiVpnHostInput, uiVpnUsername, uiVpnPassword)
	})
	//w.SetOnClosed() todo
	w.ShowAndRun()
}
