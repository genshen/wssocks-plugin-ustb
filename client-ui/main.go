package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/genshen/wssocks-plugin-ustb/extra"
	"github.com/genshen/wssocks-plugin-ustb/plugins/vpn"
	"github.com/genshen/wssocks/client"
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

func newEntryWithText(text string) *widget.Entry {
	entry := widget.NewEntry()
	entry.SetText(text)
	return entry
}

func newCheckbox(text string, checked bool, onChanged func(bool)) *widget.Check {
	checkbox := widget.NewCheck(text, onChanged)
	checkbox.SetChecked(checked)
	return checkbox
}

func main() {
	wssApp := app.NewWithID(AppId)
	wssApp.Settings().SetTheme(myTheme{})

	w := wssApp.NewWindow(AppName)
	//w.SetFixedSize(true)
	//w.Resize(fyne.NewSize(100, 100))

	// basic input
	uiLocalAddr := &widget.Entry{Text: "127.0.0.1:1080"}
	uiRemoteAddr := &widget.Entry{Text: "ws://proxy.gensh.me"}
	uiHttpEnable := newCheckbox("", false, nil)
	uiHttpLocalAddr := newEntryWithText("127.0.0.1:1086")
	uiSkipTSLVerify := newCheckbox("", false, nil)

	loadBasicPreference(wssApp.Preferences(), uiLocalAddr, uiRemoteAddr, uiHttpLocalAddr, uiHttpEnable, uiSkipTSLVerify)

	uiHttpEnable.OnChanged = func(checked bool) {
		if checked {
			uiHttpLocalAddr.Enable()
		} else {
			uiHttpLocalAddr.Disable()
		}
	}

	// vpn input
	uiVpnEnable := newCheckbox("enable ustb vpn", true, nil)
	uiVpnForceLogout := newCheckbox("", true, nil)
	uiVpnHostEncrypt := newCheckbox("", true, nil)
	uiVpnHostInput := &widget.Entry{Text: "n.ustb.edu.cn"}
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

	btnStart := widget.NewButtonWithIcon("Start", theme.MailSendIcon(), nil)
	btnStart.Importance = widget.HighImportance

	btnStatus := btnStopped
	var handles extra.TaskHandles
	btnStart.OnTapped = func() {
		if btnStatus == btnRunning { // running can stop
			btnStatus = btnStopping
			btnStart.SetText("Stopping")
			handles.NotifyCloseWrapper()
			btnStart.SetText("Start")
			btnStatus = btnStopped
		} else if btnStatus == btnStopped { // stopped can run
			options := extra.Options{
				Options: client.Options{
					LocalSocks5Addr: uiLocalAddr.Text,
					HttpEnabled:     uiHttpEnable.Checked,
					LocalHttpAddr:   uiHttpLocalAddr.Text,
					SkipTLSVerify:   uiSkipTSLVerify.Checked,
				},
				UstbVpn: vpn.UstbVpn{
					Enable:      uiVpnEnable.Checked,
					ForceLogout: uiVpnForceLogout.Checked,
					HostEncrypt: uiVpnHostEncrypt.Checked,
					TargetVpn:   uiVpnHostInput.Text,
					Username:    uiVpnUsername.Text,
					Password:    uiVpnPassword.Text,
				},
				RemoteAddr: uiRemoteAddr.Text,
			}
			btnStatus = btnStarting
			btnStart.SetText("Loading")
			if err := handles.StartWssocks(options); err != nil {
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

	w.SetContent(container.NewVBox(
		widget.NewCard("", "Basic",
			&widget.Form{Items: []*widget.FormItem{
				{Text: "socks5 address", Widget: uiLocalAddr},
				{Text: "remote address", Widget: uiRemoteAddr},
				{Text: "http(s) proxy", Widget: uiHttpEnable},
				{Text: "http(s) address", Widget: uiHttpLocalAddr},
				{Text: "skip TSL verify", Widget: uiSkipTSLVerify},
			}},
		), // end group
		widget.NewCard("", "USTB VPN",
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
		container.NewHBox(
			widget.NewLabel("core: v"+version.VERSION),
			widget.NewHyperlink("Github", repoUrl),
			widget.NewHyperlink("Document", docUrl),
		),
	))

	w.SetOnClosed(func() {
		// todo close all and stop if network lost
		if btnStatus == btnRunning { // running can stop
			btnStatus = btnStopping
			btnStart.SetText("Stopping")
			handles.NotifyCloseWrapper()
		}
		saveBasicPreference(wssApp.Preferences(), uiLocalAddr, uiRemoteAddr, uiHttpLocalAddr, uiHttpEnable, uiSkipTSLVerify)
		saveVPNPreference(wssApp.Preferences(), uiVpnEnable, uiVpnForceLogout,
			uiVpnHostEncrypt, uiVpnHostInput, uiVpnUsername, uiVpnPassword)
	})
	//w.SetOnClosed() todo
	w.ShowAndRun()
}
