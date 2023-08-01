package main

import "C"
import (
	"github.com/genshen/wssocks-plugin-ustb/extra"
	"github.com/genshen/wssocks-plugin-ustb/plugins/vpn"
	"github.com/genshen/wssocks/client"
	"unsafe"
)

// so it would not be destroyed by garbage collection
var handleInstances map[uintptr]*extra.TaskHandles

//export NewClientHandles
func NewClientHandles() uintptr {
	hd := new(extra.TaskHandles)
	ptr := uintptr(unsafe.Pointer(hd))
	if handleInstances == nil {
		handleInstances = make(map[uintptr]*extra.TaskHandles)
	}
	handleInstances[ptr] = hd
	return ptr
}

//export StartClientWrapper
func StartClientWrapper(handlesPtr uintptr, localAddr, remoteAddr, httpLocalAddr *C.char,
	httpEnable, skipTSLVerify, vpnEnable, vpnForceLogout, vpnHostEncrypt C._Bool,
	vpnHostInput, vpnUsername, vpnPassword *C.char) (err *C.char) {
	options := extra.Options{
		Options: client.Options{
			LocalSocks5Addr: C.GoString(localAddr),
			HttpEnabled:     bool(httpEnable),
			LocalHttpAddr:   C.GoString(httpLocalAddr),
			SkipTLSVerify:   bool(skipTSLVerify),
		},
		UstbVpn: vpn.UstbVpn{
			Enable:      bool(vpnEnable),
			ForceLogout: bool(vpnForceLogout),
			HostEncrypt: bool(vpnHostEncrypt),
			TargetVpn:   C.GoString(vpnHostInput),
			Username:    C.GoString(vpnUsername),
			Password:    C.GoString(vpnPassword),
		},
		RemoteAddr: C.GoString(remoteAddr),
	}
	var hp = (*extra.TaskHandles)(unsafe.Pointer(handlesPtr))
	if err := hp.StartWssocks(options); err != nil {
		return C.CString(err.Error())
	}
	return C.CString("")
}

//export WaitClientWrapper
func WaitClientWrapper(handlesPtr uintptr) *C.char {
	var hp = (*extra.TaskHandles)(unsafe.Pointer(handlesPtr))
	if err := hp.Wait(); err != nil {
		return C.CString(err.Error())
	}
	return C.CString("")
}

//export StopClientWrapper
func StopClientWrapper(handlesPtr uintptr) *C.char {
	var hp = (*extra.TaskHandles)(unsafe.Pointer(handlesPtr))
	hp.NotifyCloseWrapper()
	return C.CString("")
}

func main() {}
