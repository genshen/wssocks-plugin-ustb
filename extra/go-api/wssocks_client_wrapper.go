package main

import "C"
import (
	"github.com/genshen/wssocks-plugin-ustb/extra"
	"github.com/genshen/wssocks-plugin-ustb/plugins/vpn"
	"unsafe"
)

// so it would not be destroyed by garbage collection
var handleInstances map[uintptr]*extra.Handles

//export NewClientHandles
func NewClientHandles() uintptr {
	hd := new(extra.Handles)
	ptr := uintptr(unsafe.Pointer(hd))
	if handleInstances == nil {
		handleInstances = make(map[uintptr]*extra.Handles)
	}
	handleInstances[ptr] = hd
	return ptr
}

//export StartClientWrapper
func StartClientWrapper(handlesPtr uintptr, localAddr, remoteAddr, httpLocalAddr *C.char,
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
	var hp = (*extra.Handles)(unsafe.Pointer(handlesPtr))
	if err := hp.StartWssocks(options); err != nil {
		return C.CString(err.Error())
	}
	return C.CString("")
}

//export StopClientWrapper
func StopClientWrapper(handlesPtr uintptr) *C.char {
	var hp = (*extra.Handles)(unsafe.Pointer(handlesPtr))
	hp.Close()
	return C.CString("")
}

func main() {}
