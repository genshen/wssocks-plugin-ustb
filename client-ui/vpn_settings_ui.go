package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/genshen/wssocks-plugin-ustb/plugins/vpn"
	"github.com/genshen/wssocks-plugin-ustb/plugins/vpn/passwd"
)

type VpnSettingsUI struct {
	uiVpnForceLogout *widget.Check
	uiVpnHostEncrypt *widget.Check
	uiVpnHostInput   *widget.Entry
	uiVpnUsername    *widget.Entry
	uiVpnPassword    *widget.Entry
	uiVpnAuthMethod  *widget.RadioGroup
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

	// load Preference
	loadVpnPreference(pref, v.uiVpnAuthMethod, v.uiVpnForceLogout, v.uiVpnHostEncrypt, v.uiVpnHostInput, v.uiVpnUsername, v.uiVpnPassword)

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
			container.NewTabItem("Webview Auth", widget.NewLabel("World!")),
		),
		container.NewVBox(),
	)

	authWindow := (*wssApp).NewWindow("VPN Auth Settings")
	authWindow.SetContent(content)
	authWindow.Resize(fyne.NewSize(400, 0))
	authWindow.SetOnClosed(func() {
		// saveVpnMainPreference store the values from vpn UI to filesystem
		saveVPNPreference(pref, v.uiVpnAuthMethod, v.uiVpnForceLogout, v.uiVpnHostEncrypt, v.uiVpnHostInput, v.uiVpnUsername, v.uiVpnPassword)
	})
	authWindow.Show()
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
