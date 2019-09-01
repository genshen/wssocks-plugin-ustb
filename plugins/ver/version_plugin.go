package ver

import (
	"errors"
	"github.com/genshen/wssocks/wss"
	log "github.com/sirupsen/logrus"
)

type PluginVersionNeg struct{}

func (v *PluginVersionNeg) OnServerVersion(version wss.VersionNeg) error {
	log.WithFields(log.Fields{
		"compatible version code": version.CompVersion,
		"version code":            version.VersionCode,
		"wssocks version":         version.Version,
	}).Info("server version")

	if version.CompVersion > wss.VersionCode || wss.VersionCode > version.VersionCode {
		log.WithFields(log.Fields{
			"version code":    wss.VersionCode,
			"wssocks version": wss.CoreVersion,
		}).Info("client version")
		return errors.New("incompatible protocol version of client and server")
	}
	if version.Version != wss.CoreVersion {
		log.WithFields(log.Fields{
			"client wssocks version": wss.CoreVersion,
			"server wssocks version": version.Version,
		}).Warning("different versions of client and server wssocks.")
	}
	return nil
}
