package vpn

import (
	"net/url"
	"testing"
)

func TestVpnUrl(t *testing.T) {
	// case 1
	u, _ := url.Parse("https://abc.com")
	vpnUrl(false, USTBVpnHost, u)
	if u.String() != "http://n.ustb.edu.cn/https/abc.com" {
		t.Error("error parsing, result is", u)
	}

	// case 2
	u, _ = u.Parse("https://abc.com/path1")
	vpnUrl(false, USTBVpnHost, u)
	if u.String() != "http://n.ustb.edu.cn/https/abc.com/path1" {
		t.Error("error parsing, result is", u)
	}

	// case 3
	u, _ = u.Parse("https://abc.com/path1?ab=1")
	vpnUrl(false, USTBVpnHost, u)
	if u.String() != "http://n.ustb.edu.cn/https/abc.com/path1?ab=1" {
		t.Error("error parsing, result is", u)
	}

	// case 4
	u, _ = u.Parse("wss://abc.com/path1?ab=1")
	vpnUrl(false, USTBVpnHost, u)
	if u.String() != "ws://n.ustb.edu.cn/wss/abc.com/path1?ab=1" {
		t.Error("error parsing, result is", u)
	}

	// case 5 with port
	u, _ = u.Parse("wss://abc.com:8080/path1?ab=1")
	vpnUrl(false, USTBVpnHost, u)
	if u.String() != "ws://n.ustb.edu.cn/wss-8080/abc.com/path1?ab=1" {
		t.Error("error parsing, result is", u)
	}

	// case 6 with port
	u, _ = u.Parse("ws://abc.com:8080/path1?ab=1")
	vpnUrl(false, USTBVpnHost, u)
	if u.String() != "ws://n.ustb.edu.cn/ws-8080/abc.com/path1?ab=1" {
		t.Error("error parsing, result is", u)
	}

	// case7 6 with port
	u, _ = u.Parse("http://abc.com:8080/path1?ab=1")
	vpnUrl(false, USTBVpnHost, u)
	if u.String() != "http://n.ustb.edu.cn/http-8080/abc.com/path1?ab=1" {
		t.Error("error parsing, result is", u)
	}
}
