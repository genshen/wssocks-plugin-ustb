package vpn

import (
	"bufio"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	"github.com/genshen/cmds"
	"github.com/genshen/wssocks-plugin-ustb/plugins/vpn/passwd"
	"github.com/genshen/wssocks-plugin-ustb/plugins/vpn/qrcode"
	"github.com/genshen/wssocks-plugin-ustb/plugins/vpn/webview"
	plugin "github.com/genshen/wssocks/client"
	"github.com/genshen/wssocks/cmd/client"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	VpnAuthMethodPasswd = iota
	VpnAuthMethodQRCode
	VpnAuthMethodWebview
)

type UstbVpn struct {
	Enable      bool
	AuthMethod  int // value of VpnAuthMethodPasswd or VpnAuthMethodQRCode
	PasswdAuth  passwd.UstbVpnPasswdAuth
	QrCodeAuth  qrcode.QrCodeAuth
	WebviewAuth webview.WebviewAuth
	TargetVpn   string
	HostEncrypt bool
	ForceLogout bool
	ConnOptions plugin.Options // normal connection options
}

// create a UstbVpn instance, and add necessary command options to client sub-command.
func NewUstbVpnCli() *UstbVpn {
	vpn := UstbVpn{}
	// add more command options for client sub-command.
	if ok, clientCmd := cmds.Find(client.CommandNameClient); ok {
		clientCmd.FlagSet.BoolVar(&vpn.Enable, "vpn-enable", false, `enable USTB vpn feature.`)
		clientCmd.FlagSet.StringVar(&vpn.PasswdAuth.Username, "vpn-username", "", `username to login vpn.`)
		clientCmd.FlagSet.StringVar(&vpn.PasswdAuth.Password, "vpn-password", "", `password to login vpn.`)
		clientCmd.FlagSet.StringVar(&vpn.TargetVpn, "vpn-host", passwd.USTBVpnHost, `hostname of vpn server.`)
		clientCmd.FlagSet.BoolVar(&vpn.ForceLogout, "vpn-force-logout", false,
			`force logout account on other devices.`)
		clientCmd.FlagSet.BoolVar(&vpn.HostEncrypt, "vpn-host-encrypt", true,
			`encrypt proxy host using aes algorithm.`)
		vpn.AuthMethod = VpnAuthMethodPasswd // todo: for cli, only support password auth.
	}
	return &vpn
}

// BeforeRequest is implementation of interface RequestPlugin
// In the UstbVpn plugin, we use it for vpn auth (password auth and QR code auth).
func (v *UstbVpn) BeforeRequest(hc *http.Client, transport *http.Transport, url *url.URL, header *http.Header) error {
	if !v.Enable {
		return nil
	}

	if v.AuthMethod == VpnAuthMethodPasswd {
		return v.PasswordAuthForCookie(hc, transport, url)
	} else if v.AuthMethod == VpnAuthMethodQRCode {
		return v.QrCodeAuthForCookie(hc, transport, url)
	} else if v.AuthMethod == VpnAuthMethodWebview {
		return v.WebviewAuthForCookie(hc, transport, url)
	}
	return fmt.Errorf("unknown auth method")
}

// PasswordAuthForCookie send password to vpn server for auth,
// and keep cookie for websocket request.
// It can support cli and gui client.
func (v *UstbVpn) PasswordAuthForCookie(hc *http.Client, transport *http.Transport, url *url.URL) error {
	// read username and password if they are empty.
	if v.PasswdAuth.Username == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter username: ")
		if text, err := reader.ReadString('\n'); err != nil {
			return fmt.Errorf("error while reading username, %w", err)
		} else {
			v.PasswdAuth.Username = strings.TrimSuffix(text, "\n")
		}
	}
	if v.PasswdAuth.Password == "" {
		fmt.Print("Enter Password: ")
		if bytePassword, err := terminal.ReadPassword(int(os.Stdin.Fd())); err != nil { // error
			return fmt.Errorf("error while parsing password, %w", err)
		} else {
			v.PasswdAuth.Password = string(bytePassword)
		}
	}

	// add cookie
	al := passwd.AutoLogin{Host: v.TargetVpn, ForceLogout: v.ForceLogout, SkipTLSVerify: v.ConnOptions.SkipTLSVerify}
	if cookies, err := al.VpnLogin(v.PasswdAuth.Username, v.PasswdAuth.Password); err != nil {
		return fmt.Errorf("error vpn login: %w", err)
	} else {
		return v.SetWebSocketCookies(al.SSLEnabled, hc, transport, url, cookies)
	}
}

func (v *UstbVpn) SetWebSocketCookies(SSLEnabled bool, hc *http.Client, transport *http.Transport, url *url.URL, cookies []*http.Cookie) error {
	// In vpnLogin, we can test https support.
	// If the vpn support https, we can set transport.SkipTLSVerify if necessary.
	if SSLEnabled && v.ConnOptions.SkipTLSVerify {
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	// change target url.
	vpnUrl(v.HostEncrypt, v.TargetVpn, SSLEnabled, url)
	log.Infof("real url: %s, ssl enabled:%t", url.String(), SSLEnabled)

	if jar, err := cookiejar.New(nil); err != nil {
		return err
	} else {
		cookieUrl := *url
		// replace url scheme "wss" to "https" and "ws"to "http"
		cookieUrl.Scheme = strings.Replace(cookieUrl.Scheme, "ws", "http", 1)
		jar.SetCookies(&cookieUrl, cookies)
		hc.Jar = jar
		return nil
	}
}

func (v *UstbVpn) QrCodeAuthForCookie(hc *http.Client, transport *http.Transport, url *url.URL) error {
	if v.QrCodeAuth == nil {
		return fmt.Errorf("QrCodeAuth is not configed")
	}
	// Note: todo: check https enabled for the vpn host
	// currently, it only support https schema.
	authHttpClient := http.Client{}
	var cookies []*http.Cookie

	// step1: send request to get a frame and SID in the frame.
	var qr qrcode.QrImg
	if err := qr.ParseQRCodeImgUrl(&authHttpClient, &cookies); err != nil {
		return err
	}

	// step2: pass qr code content to show qr code in ui and wait for scan status.
	if _, err := v.QrCodeAuth.ShowQrCodeAndWait(&authHttpClient, cookies, qr); err != nil {
		return err
	} else {
		// pass cookie to websocket
		return v.SetWebSocketCookies(true, hc, transport, url, cookies)
	}
}

func (v *UstbVpn) WebviewAuthForCookie(hc *http.Client, transport *http.Transport, u *url.URL) error {
	if v.WebviewAuth == nil {
		return fmt.Errorf("WebviewAuth is not configed")
	}

	// fixme: v.TargetVpn's schema must bu checked and gnerated
	if cookies, err := v.WebviewAuth.GetCookie(hc, "https://"+v.TargetVpn); err != nil {
		return err
	} else {
		if err := v.WebviewAuth.WaitAuthFinished(); err != nil {
			return err
		} else {
			// pass cookies to websocket
			return v.SetWebSocketCookies(true, hc, transport, u, cookies)
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
			u.Scheme = passwd.USTBVpnWSSScheme
		} else {
			u.Scheme = passwd.USTBVpnWSScheme
		}
	} else { // http or https
		if ssl {
			u.Scheme = passwd.USTBVpnHttpsScheme
		} else {
			u.Scheme = passwd.USTBVpnHttpScheme
		}
	}
}
