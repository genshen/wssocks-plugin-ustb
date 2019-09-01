package main

import (
	"github.com/genshen/cmds"
	"github.com/genshen/wssocks-plugin-ustb/plugins/ver"
	"github.com/genshen/wssocks-plugin-ustb/plugins/vpn"
	"github.com/genshen/wssocks/client"
	_ "github.com/genshen/wssocks/server"
	log "github.com/sirupsen/logrus"
	//_ "github.com/genshen/wssocks/version"
	_ "github.com/genshen/wssocks-plugin-ustb/wssocks-ustb/version"
)

// initialize USTB vpn (n.ustb.edu.cn) plugin
func init() {
	vpn := vpn.NewUstbVpn()
	ver := ver.PluginVersionNeg{}
	client.AddPluginRedirect(vpn)
	client.AddPluginVersion(&ver)
}

func main() {
	cmds.SetProgramName("wssocks-ustb")
	if err := cmds.Parse(); err != nil {
		log.Fatal(err)
	}
}
