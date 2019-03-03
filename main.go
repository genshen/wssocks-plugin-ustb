package main

import (
	"github.com/genshen/cmds"
	"github.com/genshen/wssocks/client"
	_ "github.com/genshen/wssocks/client"
	_ "github.com/genshen/wssocks/server"
	"log"
	//_ "github.com/genshen/wssocks/version"
	_ "github.com/genshen/wssocks-plugin-ustb/version"
	"github.com/genshen/wssocks-plugin-ustb/vpn_plugin"
)

// initialize USTB vpn (n.ustb.edu.cn) plugin
func init() {
	vpn := vpn_plugin.UstbVpn{}
	client.AddPluginRedirect(&vpn)
}

func main() {
	cmds.SetProgramName("wssocks-ustb")
	if err := cmds.Parse(); err != nil {
		log.Fatal(err)
	}
}
