package vpn_plugin

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const USTBVpnHost = "n.ustb.edu.cn"
const USTBVpnHttpScheme = "http"
const USTBVpnWSScheme = "ws"

type UstbVpn struct{}

func (v *UstbVpn) BeforeRequest(url *url.URL, header http.Header) {
	if os.Getenv("USTB_VPN") == "" {
		return
	}
	header.Add("Cookie", os.Getenv("VPN_COOKIE"))
	vpnUrl(url)
}

func vpnUrl(u *url.URL) {
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

	u.Path = schemeWithPort + "/" + u.Host + u.Path
	u.Host = USTBVpnHost

	// set scheme
	if u.Scheme == "wss" || u.Scheme == "ws" {
		u.Scheme = USTBVpnWSScheme
	} else if u.Scheme == "https" {
		u.Scheme = USTBVpnHttpScheme
	} else {
		u.Scheme = USTBVpnHttpScheme
	}

	fmt.Println("real url:", u.String())
}
