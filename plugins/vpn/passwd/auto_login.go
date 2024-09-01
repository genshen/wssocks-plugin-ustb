package passwd

import (
	"bufio"
	"crypto/tls"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
)

const USTBVpnHost = "n.ustb.edu.cn"
const USTBVpnHttpScheme = "http"
const USTBVpnHttpsScheme = "https"
const USTBVpnWSScheme = "ws"
const USTBVpnWSSScheme = "wss"

type AutoLoginInterface interface {
	TestAddr() string
	LoginAddr() string
	LogoutAddr() string
}

type AutoLogin struct {
	Host          string
	ForceLogout   bool
	SSLEnabled    bool // the vpn server supports https
	SkipTLSVerify bool // skip tsl verify when setting https connectioon
}

func (al *AutoLogin) TestAddr(ssl bool) string {
	if ssl {
		return USTBVpnHttpsScheme + "://" + al.Host + "/"
	}
	return USTBVpnHttpScheme + "://" + al.Host + "/"
}

func (al *AutoLogin) LoginAddr(ssl bool) string {
	if ssl {
		return USTBVpnHttpsScheme + "://" + al.Host + "/do-login"
	}
	return USTBVpnHttpScheme + "://" + al.Host + "/do-login"
}

func (al *AutoLogin) LogoutAddr(ssl bool) string {
	if ssl {
		return USTBVpnHttpsScheme + "://" + al.Host + "/do-confirm-login"
	}
	return USTBVpnHttpScheme + "://" + al.Host + "/do-confirm-login"
}

// create http request client with SSLEnabled and skipTLSVerify as config
// checkRedirect will be passed into http.Client as CheckRedirect func if it is specified.
// If force is true, it will enable "InsecureSkipVerify" forcely 
// even if current connection is under http (may be redirected to https)
func (al *AutoLogin) NewHttpClient(force bool, checkRedirect func(req *http.Request, via []*http.Request) error) *http.Client {
	hc := http.Client{}
	if (force || al.SSLEnabled) && al.SkipTLSVerify {
		hc.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	if checkRedirect != nil {
		hc.CheckRedirect = checkRedirect
	}
	return &hc
}

// VpnLogin login vpn automatically and get cookie
func (al *AutoLogin) VpnLogin(uname, passwd string) ([]*http.Cookie, error) {
	// send a get request and check whether it is https protocol.
	// and save https enable/disable flag
	if httpsEnabled, err := al.testHttpsEnabled(al.Host); err != nil {
		return nil, err
	} else {
		if httpsEnabled {
			al.SSLEnabled = true
		}
	}

	var loginAddress = al.LoginAddr(al.SSLEnabled)

	form := url.Values{
		"auth_type": {"local"},
		"sms_code":  {""},
		"username":  {uname},
		"password":  {passwd},
	}

	// disable redirection here
	// If login success, it will be redirected to index page
	// and cookie would lost if we enable redirection.
	redirect := func(req *http.Request, via []*http.Request) error {
		// if upgrade http to https
		return http.ErrUseLastResponse
	}
	hc := al.NewHttpClient(false, redirect)

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
			// test connection and logout account if have login.
			if err := al.testConnect(uname, cookies); err != nil {
				return nil, err
			}
			return cookies, nil
		}
	}
}

func (al *AutoLogin) testHttpsEnabled(host string) (bool, error) {
	testUrl, err := url.Parse(USTBVpnHttpScheme + "://" + host + "/")
	if err != nil {
		return false, err
	}
	httpsSupport := false

	// disable redirection
	// if login success, it will be redirected to index page
	// and cookie would lost if we enable redirection.
	redirect := func(req *http.Request, via []*http.Request) error {
		// if upgrade http to https
		if testUrl.Scheme != req.URL.Scheme && testUrl.Path == req.URL.Path { // is http -> https redirection
			httpsSupport = true
			return nil
		}
		return http.ErrUseLastResponse
	}
	hc := al.NewHttpClient(true, redirect)

	req, err := http.NewRequest("GET", testUrl.String(), nil)
	if err != nil {
		return false, err
	}

	if resp, err := hc.Do(req); err != nil {
		return false, err
	} else {
		defer resp.Body.Close()
		return httpsSupport, nil
	}
}

func (al *AutoLogin) testConnect(uname string, cookies []*http.Cookie) error {
	hc := al.NewHttpClient(false, nil)

	req, err := http.NewRequest("GET", al.TestAddr(al.SSLEnabled), nil) // // todo missing http.
	if err != nil {
		return err
	}

	jar, _ := cookiejar.New(nil)
	jar.SetCookies(req.URL, cookies)
	hc.Jar = jar

	if resp, err := hc.Do(req); err != nil {
		return err
	} else {
		defer resp.Body.Close()
		if found, token, err := al.findLogoutToken(resp.Body); err != nil {
			return err
		} else {
			if found {
				if !al.ForceLogout { // if force logout is not enabled, just return an error.
					return errors.New("you account have been signed in on other device")
				}
				log.WithField("token", token).Info("found logout token, we will logout account.")
				log.Info("sending logout request.")
				if err := al.logoutAccount(uname, token, cookies); err != nil {
					return err
				}
			}
			// if we did not found token in http response body, we do nothing.
		}
	}
	return nil
}

func (al *AutoLogin) findLogoutToken(rd io.Reader) (bool, string, error) {
	reader := bufio.NewReader(rd)
	for {
		// read a line
		if line, _, err := reader.ReadLine(); err != nil {
			break
		} else {
			// if matched.
			matched, _ := regexp.Match(`logoutOtherToken[\s]+=[\s]+'[\w]+`, line)
			if matched { // matched
				subString := strings.Split(string(line), `'`)
				if len(subString) >= 2 {
					return true, subString[1], nil
				} else {
					return false, "", errors.New("logout token not fount")
				}
			}
		}
	}
	return false, "", nil
}

func (al *AutoLogin) logoutAccount(uname, token string, cookies []*http.Cookie) error {
	form := url.Values{
		"logoutOtherToken": {token},
		"username":         {uname},
	}

	hc := al.NewHttpClient(false,nil)

	req, err := http.NewRequest("POST", al.LogoutAddr(al.SSLEnabled),
		strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	jar, _ := cookiejar.New(nil)
	jar.SetCookies(req.URL, cookies)
	hc.Jar = jar

	if resp, err := hc.Do(req); err != nil {
		return err
	} else {
		defer resp.Body.Close()
		return nil // ok
	}
}
