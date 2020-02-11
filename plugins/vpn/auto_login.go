package vpn

import (
	"bufio"
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
const USTBVpnWSScheme = "ws"

type AutoLoginInterface interface {
	TestAddr() string
	LoginAddr() string
	LogoutAddr() string
}

type AutoLogin struct {
	Host        string
	ForceLogout bool
}

func (al *AutoLogin) TestAddr() string {
	return USTBVpnHttpScheme + "://" + al.Host + "/"
}

func (al *AutoLogin) LoginAddr() string {
	return USTBVpnHttpScheme + "://" + al.Host + "/do-login"
}

func (al *AutoLogin) LogoutAddr() string {
	return USTBVpnHttpScheme + "://" + al.Host + "/do-confirm-login"
}

// auto login vpn and get cookie
func (al *AutoLogin) vpnLogin(uname, passwd string) ([]*http.Cookie, error) {
	var loginAddress = al.LoginAddr()
	loginUrl, err := url.Parse(loginAddress)
	if err != nil {
		return nil, err
	}

	form := url.Values{
		"auth_type": {"local"},
		"sms_code":  {""},
		"username":  {uname},
		"password":  {passwd},
	}

	hc := http.Client{
		// disable redirection
		// If login success, it will be redirected to index page
		// and cookie would lost if we enable redirection.
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// if upgrade http to https
			if loginUrl.Scheme != req.URL.Scheme && loginUrl.Path == req.URL.Path { // is http -> https redirection
				return nil
			}
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
			// test connection and logout account if have login.
			if err := al.testConnect(uname, cookies); err != nil {
				return nil, err
			}
			return cookies, nil
		}
	}
}

func (al *AutoLogin) testConnect(uname string, cookies []*http.Cookie) error {
	hc := http.Client{}
	req, err := http.NewRequest("GET", al.TestAddr(), nil) // // todo missing http.
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

	hc := http.Client{}
	req, err := http.NewRequest("POST", al.LogoutAddr(),
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
