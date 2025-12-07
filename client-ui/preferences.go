package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/genshen/wssocks-plugin-ustb/plugins/vpn"
	"strings"
)

const (
	PrefHasPreference  = "has_preference"
	PrefLocalAddr      = "local_addr"
	PrefRemoteAddr     = "remote_addr"
	PrefHttpEnable     = "http_enable"
	PrefHttpLocalAddr  = "http_local_addr"
	PrefSkipTSLVerify  = "skip_TSL_verify"
	PrefVpnEnable      = "vpn_enable"
	PrefVpnAuthMethod  = "auth_method"
	PrefVpnForceLogout = "vpn_force_logout"
	PrefVpnHostEncrypt = "vpn_host_encrypt"
	PrefVpnHostInput   = "vpn_host"
	PrefVpnUsername    = "vpn_username"
	PrefVpnPassword    = "vpn_password"
)

func saveBasicPreference(pref fyne.Preferences, uiLocalAddr, uiRemoteAddr,
	uiHttpLocalAddr *widget.Entry, uiHttpEnable *widget.Check,
	uiSkipTSLVerify *widget.Check) {
	pref.SetBool(PrefHasPreference, true)
	pref.SetString(PrefLocalAddr, uiLocalAddr.Text)
	pref.SetString(PrefRemoteAddr, uiRemoteAddr.Text)

	pref.SetBool(PrefHttpEnable, uiHttpEnable.Checked)
	pref.SetString(PrefHttpLocalAddr, uiHttpLocalAddr.Text)
	pref.SetBool(PrefSkipTSLVerify, uiSkipTSLVerify.Checked)
}

func saveVPNMainPreference(pref fyne.Preferences,
	uiVpnEnable *widget.Check) {
	pref.SetBool(PrefVpnEnable, uiVpnEnable.Checked)
}

func saveVPNPreference(pref fyne.Preferences, uiVpnAuthMethod *widget.RadioGroup, uiVpnForceLogout, uiVpnHostEncrypt *widget.Check,
	uiVpnHostInput, uiVpnUsername, uiVpnPassword *widget.Entry) {
	if !pref.Bool(PrefHasPreference) {
		return
	}

	pref.SetBool(PrefVpnForceLogout, uiVpnForceLogout.Checked)
	pref.SetBool(PrefVpnHostEncrypt, uiVpnHostEncrypt.Checked)
	pref.SetString(PrefVpnHostInput, uiVpnHostInput.Text)
	pref.SetString(PrefVpnUsername, uiVpnUsername.Text)
	//pref.SetString(PrefVpnPassword,uiVpnPassword.Text)
	if uiVpnAuthMethod.Selected == TextVpnAuthMethodPasswd {
		pref.SetInt(PrefVpnAuthMethod, vpn.VpnAuthMethodPasswd)
	} else if uiVpnAuthMethod.Selected == TextVpnAuthMethodQrCode {
		pref.SetInt(PrefVpnAuthMethod, vpn.VpnAuthMethodQRCode)
	}
}

func loadBasicPreference(pref fyne.Preferences, uiLocalAddr, uiRemoteAddr,
	uiHttpLocalAddr *widget.Entry, uiHttpEnable *widget.Check,
	uiSkipTSLVerify *widget.Check) {
	if !pref.Bool(PrefHasPreference) {
		uiHttpLocalAddr.Disable()
		return
	}

	// local address
	if localAddr := pref.String(PrefLocalAddr); strings.TrimSpace(localAddr) != "" {
		uiLocalAddr.SetText(strings.TrimSpace(localAddr))
	}
	// remote address
	if remoteAddr := pref.String(PrefRemoteAddr); strings.TrimSpace(remoteAddr) != "" {
		uiRemoteAddr.SetText(strings.TrimSpace(remoteAddr))
	}
	// http enable (default false)
	if pref.Bool(PrefHttpEnable) {
		uiHttpEnable.SetChecked(true)
	}
	// http local address
	if httpAddr := pref.String(PrefHttpLocalAddr); strings.TrimSpace(httpAddr) != "" {
		uiHttpLocalAddr.SetText(strings.TrimSpace(httpAddr))
	}
	// skip TSL verify
	if pref.Bool(PrefSkipTSLVerify) {
		uiSkipTSLVerify.SetChecked(true)
	}

	if !uiHttpEnable.Checked {
		uiHttpLocalAddr.Disable()
	}
}

func loadVPNMainPreference(pref fyne.Preferences, uiVpnEnable *widget.Check) {
	if !pref.Bool(PrefHasPreference) {
		return
	}
	// vpn enable
	if enable := pref.Bool(PrefVpnEnable); !enable {
		uiVpnEnable.SetChecked(enable) // toggle default value
	} // else, default value(true) or preference is true, dont touch it.

}

func loadVpnPreference(pref fyne.Preferences, uiVpnAuthMethod *widget.RadioGroup, uiVpnForceLogout, uiVpnHostEncrypt *widget.Check,
	uiVpnHostInput, uiVpnUsername, uiVpnPassword *widget.Entry) {
	if !pref.Bool(PrefHasPreference) {
		return
	}
	// vpn force logout
	if enable := pref.Bool(PrefVpnForceLogout); !enable {
		uiVpnForceLogout.SetChecked(enable)
	}
	// vpn force logout
	if enable := pref.Bool(PrefVpnHostEncrypt); !enable {
		uiVpnHostEncrypt.SetChecked(enable)
	}

	// vpn auth method
	authMethod := pref.Int(PrefVpnAuthMethod)
	if authMethod == vpn.VpnAuthMethodPasswd {
		uiVpnAuthMethod.SetSelected(TextVpnAuthMethodPasswd)
	} else if authMethod == vpn.VpnAuthMethodQRCode {
		uiVpnAuthMethod.SetSelected(TextVpnAuthMethodQrCode)
	} else {
		// todo error
	}

	// vpn host, username, password
	if host := pref.String(PrefVpnHostInput); strings.TrimSpace(host) != "" {
		uiVpnHostInput.SetText(strings.TrimSpace(host))
	}
	if username := pref.String(PrefVpnUsername); strings.TrimSpace(username) != "" {
		uiVpnUsername.SetText(strings.TrimSpace(username))
	}
	//if password := pref.String(PrefVpnPassword); password != "" {
	//	uiVpnPassword.SetText(password)
	//}
}
