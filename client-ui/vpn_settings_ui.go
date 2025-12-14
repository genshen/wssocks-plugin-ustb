package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/genshen/wssocks-plugin-ustb/plugins/vpn"
	"github.com/genshen/wssocks-plugin-ustb/plugins/vpn/passwd"
)

type VpnSettingsUI struct {
	uiVpnForceLogout      *widget.Check
	uiVpnHostEncrypt      *widget.Check
	uiVpnHostInput        *widget.Entry
	uiVpnUsername         *widget.Entry
	uiVpnPassword         *widget.Entry
	uiVpnAuthMethod       *widget.RadioGroup
	uiChromePathContainer *widget.Entry
}

func (v *VpnSettingsUI) OpenVpnSettings(wssApp *fyne.App, pref fyne.Preferences) {
	v.uiVpnForceLogout = newCheckbox("", true, nil)
	v.uiVpnHostEncrypt = newCheckbox("", true, nil)
	v.uiVpnHostInput = &widget.Entry{PlaceHolder: "vpn hostname", Text: "n.ustb.edu.cn"}
	v.uiVpnUsername = &widget.Entry{PlaceHolder: "vpn username", Text: ""}
	v.uiVpnPassword = &widget.Entry{PlaceHolder: "vpn password", Text: "", Password: true}

	// select auth method
	v.uiVpnAuthMethod = widget.NewRadioGroup([]string{TextVpnAuthMethodPasswd, TextVpnAuthMethodQrCode, TextVpnAuthMethodWebview}, func(value string) {
		// todo:
	})
	v.uiVpnAuthMethod.Horizontal = true

	v.uiChromePathContainer = widget.NewEntry()
	v.uiChromePathContainer.SetPlaceHolder("browser path")
	v.uiChromePathContainer.Disable()

	// load Preference
	loadVpnPreference(pref, v.uiVpnAuthMethod, v.uiVpnForceLogout, v.uiVpnHostEncrypt, v.uiVpnHostInput, v.uiVpnUsername, v.uiVpnPassword, v.uiChromePathContainer)

	authWindow := (*wssApp).NewWindow("VPN Auth Settings")

	content := container.NewVBox(
		&widget.Form{Items: []*widget.FormItem{
			{Text: "force logout", Widget: v.uiVpnForceLogout},
			{Text: "host encrypt", Widget: v.uiVpnHostEncrypt},
			{Text: "vpn host", Widget: v.uiVpnHostInput},
			{Text: "auth method", Widget: v.uiVpnAuthMethod},
		}},
		&widget.Separator{},
		container.NewAppTabs(
			container.NewTabItem("Password Auth",
				&widget.Form{Items: []*widget.FormItem{
					{Text: "username", Widget: v.uiVpnUsername},
					{Text: "password", Widget: v.uiVpnPassword},
				}}),
			container.NewTabItem("QR Code Auth", widget.NewLabel("World!")),
			container.NewTabItem("Webview Auth", container.NewVBox(
				loadFilePicker(authWindow, v.uiChromePathContainer),
				v.uiChromePathContainer,
			)),
		),
		container.NewVBox(),
	)

	authWindow.SetContent(content)
	authWindow.Resize(fyne.NewSize(400, 0))
	authWindow.SetOnClosed(func() {
		// saveVpnMainPreference store the values from vpn UI to filesystem
		saveVPNPreference(pref, v.uiVpnAuthMethod, v.uiVpnForceLogout, v.uiVpnHostEncrypt, v.uiVpnHostInput, v.uiVpnUsername, v.uiVpnPassword, v.uiChromePathContainer)
	})
	authWindow.Show()
}

func loadFilePicker(win fyne.Window, pathContainer *widget.Entry) *widget.Button {
	selectBtn := widget.NewButton("Select Chrome/Edge/Chromium Browser", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				return
			}
			if reader == nil {
				return
			}
			pathContainer.SetText(reader.URI().Path())
		}, win)

		fd.Show()
		fd.SetOnClosed(func() {})
	})
	return selectBtn
}

func (v *VpnSettingsUI) LoadSettingsValues(values *vpn.UstbVpn) {
	values.ForceLogout = v.uiVpnForceLogout.Checked
	values.HostEncrypt = v.uiVpnHostEncrypt.Checked
	values.TargetVpn = v.uiVpnHostInput.Text
	values.AuthMethod = getAuthMethodInt(v.uiVpnAuthMethod)
	values.PasswdAuth = passwd.UstbVpnPasswdAuth{
		Username: v.uiVpnUsername.Text,
		Password: v.uiVpnPassword.Text,
	}
}

// convert from  selected string to int value
func getAuthMethodInt(uiVpnAuthMethod *widget.RadioGroup) int {
	if uiVpnAuthMethod.Selected == TextVpnAuthMethodPasswd {
		return vpn.VpnAuthMethodPasswd
	} else if uiVpnAuthMethod.Selected == TextVpnAuthMethodQrCode {
		return vpn.VpnAuthMethodQRCode
	} else {
		return vpn.VpnAuthMethodWebview
	}
}
