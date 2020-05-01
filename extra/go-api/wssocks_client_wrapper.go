package main

import "C"
import (
	"github.com/genshen/wssocks-plugin-ustb/extra"
	"github.com/genshen/wssocks-plugin-ustb/plugins/vpn"
)

//export StartClientWrapper
func StartClientWrapper(localAddr, remoteAddr, httpLocalAddr *C.char,
	httpEnable, vpnEnable, vpnForceLogout, vpnHostEncrypt C._Bool,
	vpnHostInput, vpnUsername, vpnPassword *C.char) (err *C.char) {
	options := extra.Options{
		LocalSocks5Addr: C.GoString(localAddr),
		RemoteAddr:      C.GoString(remoteAddr),
		HttpEnable:      bool(httpEnable),
		LocalHttpAddr:   C.GoString(httpLocalAddr),
		UstbVpn: vpn.UstbVpn{
			Enable:      bool(vpnEnable),
			ForceLogout: bool(vpnForceLogout),
			HostEncrypt: bool(vpnHostEncrypt),
			TargetVpn:   C.GoString(vpnHostInput),
			Username:    C.GoString(vpnUsername),
			Password:    C.GoString(vpnPassword),
		},
	}
	var handles extra.Handles
	if err := handles.StartWssocks(options); err != nil {
		return C.CString(err.Error())
	}
	return C.CString("")
}

//export StopClientWrapper
func StopClientWrapper() *C.char {
	return C.CString("")
}

func main() {}
