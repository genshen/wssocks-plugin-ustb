package vpn

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/genshen/cmds"
	"github.com/genshen/wssocks/client"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)

type UstbVpn struct {
	Enable      bool
	Username    string
	Password    string
	TargetVpn   string
	HostEncrypt bool
	ForceLogout bool
}

// create a UstbVpn instance, and add necessary command options to client sub-command.
func NewUstbVpnCli() *UstbVpn {
	vpn := UstbVpn{}
	// add more command options for client sub-command.
	if ok, clientCmd := cmds.Find(client.CommandNameClient); ok {
		clientCmd.FlagSet.BoolVar(&vpn.Enable, "vpn-enable", false, `enable USTB vpn feature.`)
		clientCmd.FlagSet.StringVar(&vpn.Username, "vpn-username", "", `username to login vpn.`)
		clientCmd.FlagSet.StringVar(&vpn.Password, "vpn-password", "", `password to login vpn.`)
		clientCmd.FlagSet.StringVar(&vpn.TargetVpn, "vpn-host", USTBVpnHost, `hostname of vpn server.`)
		clientCmd.FlagSet.BoolVar(&vpn.ForceLogout, "vpn-force-logout", false,
			`force logout account on other devices.`)
		clientCmd.FlagSet.BoolVar(&vpn.HostEncrypt, "vpn-host-encrypt", true,
			`encrypt proxy host using aes algorithm.`)
	}
	return &vpn
}

func (v *UstbVpn) BeforeRequest(dialer *websocket.Dialer, url *url.URL, header http.Header) error {
	if !v.Enable {
		return nil
	}
	// read username and password if they are empty.
	if v.Username == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter username: ")
		if text, err := reader.ReadString('\n'); err != nil {
			return fmt.Errorf("error while reading username, %w", err)
		} else {
			v.Username = strings.TrimSuffix(text, "\n")
		}
	}
	if v.Password == "" {
		fmt.Print("Enter Password: ")
		if bytePassword, err := terminal.ReadPassword(int(os.Stdin.Fd())); err != nil { // error
			return fmt.Errorf("error while parsing password, %w", err)
		} else {
			v.Password = string(bytePassword)
		}
	}

	// add cookie
	al := AutoLogin{Host: v.TargetVpn, ForceLogout: v.ForceLogout}
	if cookies, err := al.vpnLogin(v.Username, v.Password); err != nil {
		return err
	} else {
		// change target url.
		vpnUrl(v.HostEncrypt, v.TargetVpn, al.SSLEnabled, url)
		log.Infof("real url: %s", url.String(), ", ssl enabled:", al.SSLEnabled)

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

// ssl specific the protocol(whether to use ssl) used in the real connection
func vpnUrl(hostEncrypt bool, vpnHost string, ssl bool, u *url.URL) {
	// replace https://abc.com to "http://n.ustb.edu.cn/https/abc.com"
	// replace https://abc.com:8080 to "http://n.ustb.edu.cn/https-8080/abc.com"

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
		u.Path = "/" + schemeWithPort + "/" + hex.EncodeToString([]byte(key)) + hex.EncodeToString(encryptHost) + u.Path
	} else {
		u.Path = "/" + schemeWithPort + "/" + u.Host + u.Path
	}
	if !strings.HasSuffix(u.Path, "/") {
		u.Path = u.Path + "/"
	}
	u.Host = vpnHost

	// set scheme
	if u.Scheme == "wss" || u.Scheme == "ws" {
		if ssl {
			u.Scheme = USTBVpnWSSScheme
		} else {
			u.Scheme = USTBVpnWSScheme
		}
	} else { // http or https
		if ssl {
			u.Scheme = USTBVpnHttpsScheme
		} else {
			u.Scheme = USTBVpnHttpScheme
		}
	}
}
