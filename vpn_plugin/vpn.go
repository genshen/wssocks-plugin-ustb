package vpn_plugin

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/genshen/cmds"
	"github.com/genshen/wssocks/client"
	"github.com/gorilla/websocket"
	"github.com/howeyc/gopass"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)

const USTBVpnHost = "n.ustb.edu.cn"
const USTBVpnHttpScheme = "http"
const USTBVpnLoginUrl = USTBVpnHttpScheme + "://" + USTBVpnHost + "/do-login"
const USTBVpnWSScheme = "ws"

type UstbVpn struct {
	enable    bool
	username  string
	password  string
	targetUrl string
	hostEncrypt bool
}

// create a UstbVpn instance, and add necessary command options to client sub-command.
func NewUstbVpn() *UstbVpn {
	vpn := UstbVpn{}
	// add more command options for client sub-command.
	if ok, clientCmd := cmds.Find(client.CommandNameClient); ok {
		clientCmd.FlagSet.BoolVar(&vpn.enable, "vpn-enable", false, `enable USTB vpn feature.`)
		clientCmd.FlagSet.StringVar(&vpn.username, "vpn-username", "", `username to login vpn.`)
		clientCmd.FlagSet.StringVar(&vpn.password, "vpn-password", "", `password to login vpn.`)
		clientCmd.FlagSet.StringVar(&vpn.targetUrl, "vpn-login-url", USTBVpnLoginUrl, `address to login vpn.`)
		clientCmd.FlagSet.BoolVar(&vpn.hostEncrypt, "vpn-host-encrypt", true,
			`encrypt proxy host using aes algorithm.`)
	}
	return &vpn
}

func (v *UstbVpn) BeforeRequest(dialer *websocket.Dialer, url *url.URL, header http.Header) error {
	if !v.enable {
		return nil
	}
	// read username and password if they are empty.
	if v.username == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter username: ")
		if text, err := reader.ReadString('\n'); err != nil {
			return errors.New("Whoops! Error while reading username:" + err.Error())
		} else {
			v.username = strings.TrimSuffix(text,"\n")
		}
	}
	if v.password == "" {
		fmt.Print("Enter Password: ")
		if bytePassword, err := gopass.GetPasswd(); err != nil { // error
			return errors.New("Whoops! Error while parsing password:" + err.Error())
		} else {
			v.password = string(bytePassword)
		}
	}

	// change target url.
	vpnUrl(v.hostEncrypt, USTBVpnHost, url)
	fmt.Println("real url:", url.String())

	// add cookie
	if cookies, err := vpnLogin(v.targetVpn, v.username, v.password); err != nil {
		return err
	} else {
		if jar, err := cookiejar.New(nil); err != nil {
			return err
		} else {
			dialer.Jar = jar
			cookieUrl := *url
			// replace url scheme "wss" to "https" and "ws"to "http"
			cookieUrl.Scheme = strings.Replace(cookieUrl.Scheme, "ws", "http", 1)
			dialer.Jar.SetCookies(&cookieUrl, cookies)
			return nil
		}
	}
}

func vpnUrl(hostEncrypt bool, vpnHost string, u *url.URL) {
	// replace https://abc.com to "http://n.ustb.edu.cn/https/abc.com"
	// replace https://abc.com:8080 to "http://n.ustb.edu.cn/https-8080/abc.com"
	// ?wrdrecordrvisit=record

	// split host and port if it could
	port := u.Port()
	if strings.ContainsRune(u.Host, ':') {
		if h, p, err := net.SplitHostPort(u.Host); err != nil {
			panic(err)
		} else {
			u.Host = h
			if port != "" {
				port = p
			}
		}
	}

	schemeWithPort := u.Scheme
	if (u.Scheme == "wss" || u.Scheme == "https") && port != "" && port != "443" {
		schemeWithPort = u.Scheme + "-" + port
	}
	if (u.Scheme == "ws" || u.Scheme == "http") && port != "" && port != "80" {
		schemeWithPort = u.Scheme + "-" + port
	}

	if hostEncrypt {
		const key = "wrdvpnisthebest!"
		var aes_e = newAesEncrypt(key)
		encryptHost, _ := aes_e.Encrypt(u.Host)
		u.Path = schemeWithPort + "/" + hex.EncodeToString([]byte(key)) + hex.EncodeToString(encryptHost) + u.Path
	} else {
		u.Path = schemeWithPort + "/" + u.Host + u.Path
	}
	u.Host = vpnHost

	// set scheme
	if u.Scheme == "wss" || u.Scheme == "ws" {
		u.Scheme = USTBVpnWSScheme
	} else if u.Scheme == "https" {
		u.Scheme = USTBVpnHttpScheme
	} else {
		u.Scheme = USTBVpnHttpScheme
	}
}
