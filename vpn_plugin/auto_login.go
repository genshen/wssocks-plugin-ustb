package vpn_plugin

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const USTBVpnHost = "n.ustb.edu.cn"
const USTBVpnHttpScheme = "http"
const USTBVpnWSScheme = "ws"

// auto login vpn and get cookie
func vpnLogin(loginHost, uname, passwd string) ([]*http.Cookie, error) {
	var loginAddress = USTBVpnHttpScheme + "://" + loginHost + "/do-login"
	form := url.Values{
		"auth_type": {"local"},
		"sms_code":  {""},
		"username":  {uname},
		"password":  {passwd},
	}

	hc := http.Client{
		// disable redirection
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest("POST", loginAddress, strings.NewReader(form.Encode())) // todo missing http.
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if resp, err := hc.Do(req); err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()
		cookies := resp.Cookies()
		// return cookies or error.
		if len(cookies) == 0 {
			return nil, errors.New(fmt.Sprintf("no cookie while auto login to %s ", loginAddress))
		} else {
			return cookies, nil
		}
	}
}
