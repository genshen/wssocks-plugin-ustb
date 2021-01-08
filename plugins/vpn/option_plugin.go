package vpn

import "github.com/genshen/wssocks/client"

// implementation of interface OptionPlugin
func (v *UstbVpn) OnOptionSet(options client.Options) error {
	v.ConnOptions = options
	return nil
}
