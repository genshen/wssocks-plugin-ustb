package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	resource "github.com/genshen/wssocks-plugin-ustb/client-ui/resources"
	"github.com/genshen/wssocks-plugin-ustb/extra"
	"github.com/genshen/wssocks-plugin-ustb/plugins/vpn"
	pluginversion "github.com/genshen/wssocks-plugin-ustb/wssocks-ustb/version"
	"github.com/genshen/wssocks/client"
	"github.com/genshen/wssocks/version"
	"net/url"
)

const (
	AppName           = "wssocks Client"
	AppId             = "wssocks-ustb.genshen.github.com"
	CoreGithubRepoUrl = "https://github.com/genshen/wssocks"
	GithubRepoUrl     = "https://github.com/genshen/wssocks-plugin-ustb"
	DocumentUrl       = "https://genshen.github.io/wssocks-plugin-ustb/"
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
	wssApp.Settings().SetTheme(&myTheme{})

	w := wssApp.NewWindow(AppName)
	//w.SetFixedSize(true)
	//w.Resize(fyne.NewSize(100, 100))

	// basic input
	uiLocalAddr := &widget.Entry{PlaceHolder: "socks5 listen address", Text: "127.0.0.1:1080"}
	uiRemoteAddr := &widget.Entry{PlaceHolder: "wssocks server address"}
	uiHttpEnable := newCheckbox("", false, nil)
	uiHttpLocalAddr := &widget.Entry{PlaceHolder: "http listen address", Text: "127.0.0.1:1086"}
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
	uiVpnHostInput := &widget.Entry{PlaceHolder: "vpn hostname", Text: "n.ustb.edu.cn"}
	uiVpnUsername := &widget.Entry{PlaceHolder: "vpn username", Text: ""}
	uiVpnPassword := &widget.Entry{PlaceHolder: "vpn password", Text: "", Password: true}

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

	coreRepoUrl, err := url.Parse(CoreGithubRepoUrl)
	if err != nil {
		return
	}

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon(
			"Basic",
			theme.SettingsIcon(),
			&widget.Form{Items: []*widget.FormItem{
				{Text: "socks5 address", Widget: uiLocalAddr},
				{Text: "remote address", Widget: uiRemoteAddr},
				{Text: "http(s) proxy", Widget: uiHttpEnable},
				{Text: "http(s) address", Widget: uiHttpLocalAddr},
				{Text: "skip TSL verify", Widget: uiSkipTSLVerify},
			}},
		),
		container.NewTabItemWithIcon(
			"USTB VPN",
			theme.AccountIcon(),
			&widget.Form{Items: []*widget.FormItem{
				{Text: "enable", Widget: uiVpnEnable},
				{Text: "force logout", Widget: uiVpnForceLogout},
				{Text: "host encrypt", Widget: uiVpnHostEncrypt},
				{Text: "vpn host", Widget: uiVpnHostInput},
				{Text: "username", Widget: uiVpnUsername},
				{Text: "password", Widget: uiVpnPassword},
			}},
		),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	w.SetContent(container.NewVBox(
		widget.NewCard("Settings", "", tabs),
		btnStart,
		&widget.Separator{},
		container.NewGridWithColumns(2,
			container.NewHBox(
				NewHyperlinkIcon(resource.GithubIcon(), coreRepoUrl),
				widget.NewHyperlink("wssocks core: ", coreRepoUrl),
			),
			widget.NewLabel("v"+version.VERSION),
		),
		container.NewGridWithColumns(2,
			container.NewHBox(
				NewHyperlinkIcon(resource.GithubIcon(), repoUrl),
				widget.NewHyperlink("USTB vpn plugin: ", repoUrl),
			),
			container.NewGridWithColumns(2,
				widget.NewLabel("v"+pluginversion.VERSION),
				container.NewHBox(
					layout.NewSpacer(),
					widget.NewToolbar(
						widget.NewToolbarAction(theme.HelpIcon(), func() {
							if err := fyne.CurrentApp().OpenURL(docUrl); err != nil {
								dialog.ShowError(fmt.Errorf("open link %s failed", docUrl), w)
							}
						}),
					),
				),
			),
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
