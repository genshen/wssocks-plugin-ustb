package webview

import (
	"github.com/chromedp/cdproto/network"
	"net/http"
	"time"
)

// networkCookieToHTTPCookie convert chromedp cookie to http cookie.
func convertChromedpCookieToHTTPCookie(nc *network.Cookie) *http.Cookie {
	hc := &http.Cookie{
		Name:     nc.Name,
		Value:    nc.Value,
		Path:     nc.Path,
		Domain:   nc.Domain,
		Secure:   nc.Secure,
		HttpOnly: nc.HTTPOnly,
		SameSite: convertSameSite(nc.SameSite),
	}

	if nc.Expires > 0 { // if nc.Expires < 0, we ignore it.
		// unit of network.Cookie.Expires is second
		hc.Expires = time.Unix(int64(nc.Expires), 0)
	}

	if !hc.Expires.IsZero() {
		hc.MaxAge = int(hc.Expires.Sub(time.Now()).Seconds())
	}

	return hc
}

func convertSameSite(ss network.CookieSameSite) http.SameSite {
	switch ss {
	case network.CookieSameSiteStrict:
		return http.SameSiteStrictMode
	case network.CookieSameSiteLax:
		return http.SameSiteLaxMode
	case network.CookieSameSiteNone:
		return http.SameSiteNoneMode
	default:
		return http.SameSiteDefaultMode
	}
}
