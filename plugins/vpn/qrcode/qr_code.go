package qrcode

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

const LoadImgUrl = "https://n.ustb.edu.cn/login/"
const SisAuthPath = "https://sis.ustb.edu.cn"
const FindQrcodeUrlRegex = `"ustb-qrcode",`
const FindQrcodeImgTagRegex = `<img`

// QRCodeImgLoaderConfig is the config to generate an url of sis.ustb.edu.cn for requesting qr code image.
// here we treat js as yaml (because the source string is not standard json, we can not parse it as a json)
type QRCodeImgLoaderConfig struct {
	Id        string `yaml:"id"`
	ApiUrl    string `yaml:"api_url"`
	AppID     string `yaml:"appid"`
	ReturnUrl string `yaml:"return_url"`
	RandToken string `yaml:"rand_token"`
}

// GenUrl generates the url of the iframe.
func (q *QRCodeImgLoaderConfig) genIframeUrl() (string, error) {
	if q.ApiUrl == "" {
		return "", errors.New("api url is empty")
	}
	if q.AppID == "" {
		return "", errors.New("app id is empty")
	}
	if q.ReturnUrl == "" {
		return "", errors.New("return url is empty")
	}
	if q.RandToken == "" {
		return "", errors.New("rand token is empty")
	}
	// todo  encodeURI for return_url
	return fmt.Sprintf("%s?appid=%s&return_url=%s&rand_token=%s&embed_flag=1", q.ApiUrl, q.AppID, q.ReturnUrl, q.RandToken), nil
}

type QrImg struct {
	Config QRCodeImgLoaderConfig
	Sid    string // sis in ustb auth, can be parsed from image url.
}

type QrCodeAuth interface {
	ShowQrCodeAndWait(client *http.Client, cookies []*http.Cookie, qrCode QrImg) ([]*http.Cookie, error)
}

// ParseQRCodeImgUrl uses ParseQRCodeHtmlUrl to get the iframe html,
// and then parse the html file to get final image url (contains SID).
// And set QrImg's fields of config and sid.
func (i *QrImg) ParseQRCodeImgUrl(client *http.Client, cookies *[]*http.Cookie) error {
	qrImgUrlConfig, err := ParseQRCodeHtmlUrl(client, cookies)
	if err != nil {
		return err
	}
	i.Config = qrImgUrlConfig
	// generate iframe url.
	iframeUrl, err := qrImgUrlConfig.genIframeUrl()
	if err != nil {
		return err
	}

	// make a http request of the iframe
	req, err := http.Get(iframeUrl)
	if err != nil {
		return err
	}

	defer req.Body.Close()

	scanner := bufio.NewScanner(req.Body)
	var imgUrl string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, FindQrcodeImgTagRegex) { // first line to match start
			// line e.g. <img id="qrimg" src="/connect/qrimg?sid=3894c5568dd1ef0f6434f426297a678d" height="90%" border="0">
			subStr := strings.SplitN(line, "\"", 5)
			if len(subStr) != 5 {
				return errors.New("invalid format in qr image url parsing")
			} else {
				imgUrl = subStr[3]
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// parse sid in qr image url.
	sidStr := strings.SplitN(imgUrl, "=", 2)
	if len(sidStr) != 2 {
		return errors.New("invalid format in qr image url (sis) parsing")
	} else {
		i.Sid = sidStr[1]
	}
	return nil
}

func ParseQRCodeHtmlUrl(client *http.Client, cookies *[]*http.Cookie) (QRCodeImgLoaderConfig, error) {
	// parse the html of `LOAD_IMG_URL` to get following object text:
	//{
	//  id: "ustb-qrcode",
	//	api_url: "",
	//	appid: "",
	//	return_url: "",
	//	rand_token: "",
	//	width: "",
	//	height: ""
	//}
	// make a http request of the iframe
	req, err := http.NewRequest("GET", LoadImgUrl, nil)
	if err != nil {
		return QRCodeImgLoaderConfig{}, err
	}

	response, err := client.Do(req)
	if err != nil {
		return QRCodeImgLoaderConfig{}, err
	}
	defer response.Body.Close()

	*cookies = response.Cookies() // save cookies
	if len(*cookies) == 0 {
		return QRCodeImgLoaderConfig{}, fmt.Errorf("no cookie found while getting iframe")
	}
	log.Println("COOKIE:", *cookies)

	scanner := bufio.NewScanner(response.Body)
	var findQrMatchStart = false
	var qrConfigBuffer bytes.Buffer
	qrConfigBuffer.WriteString("{")
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, FindQrcodeUrlRegex) { // first line to match start
			findQrMatchStart = true // start match
		}
		// only start matched and end not matched, we can add text to buffer.
		if findQrMatchStart {
			qrConfigBuffer.WriteString(line)
			if strings.Contains(line, "}") {
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return QRCodeImgLoaderConfig{}, err
	}

	fmt.Println("parsed qr code config:", qrConfigBuffer.String())
	var qrImgUrlConfig QRCodeImgLoaderConfig
	if err := yaml.Unmarshal(qrConfigBuffer.Bytes(), &qrImgUrlConfig); err != nil {
		return QRCodeImgLoaderConfig{}, err
	}

	return qrImgUrlConfig, nil
}

func (i *QrImg) GenQrCodeContent() string {
	return fmt.Sprintf(SisAuthPath+"/auth?sid=%s", i.Sid)
}

// GenQrImgUrl generate the url of qr code image
func (i *QrImg) GenQrImgUrl(imgUrl string) (string, error) {
	iframeUrl, err := i.Config.genIframeUrl()
	if err != nil {
		return "", err
	}

	htmlUri, err := url.Parse(iframeUrl)
	if err != nil {
		return "", err
	}
	// use htmlUri's host, schema
	return fmt.Sprintf("%s://%s%s", htmlUri.Scheme, htmlUri.Host, imgUrl), nil
}

type StateResponseAuthData struct {
	State int    `json:"state"`
	Data  string `json:"data"`
}

// WaitQrState waits qr state and get auth code (as return value)
func WaitQrState(sid string) (string, error) {
	stateUrl := fmt.Sprintf(SisAuthPath+"/connect/state?sid=%s", sid)
	response, err := http.Get(stateUrl)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	authData := StateResponseAuthData{}
	if err := json.Unmarshal(body, &authData); err != nil {
		return "", err
	}
	if authData.State != 200 {
		return "", fmt.Errorf("auth status is not 200, but %d", authData.State)
	}
	return authData.Data, nil
}

// RedirectToLogin sends callback request.
func RedirectToLogin(client *http.Client, cookies []*http.Cookie, appid, authCode, randToken string) error {
	loginUrl := fmt.Sprintf(LoadImgUrl+"?ustb_sis=true&appid=%s&auth_code=%s&rand_token=%s", appid, authCode, randToken)
	// todo: generate login url based on return url.
	log.Println("redirect url:", loginUrl)

	req, err := http.NewRequest("GET", loginUrl, nil)
	if err != nil {
		return err
	}

	jar, _ := cookiejar.New(nil)
	jar.SetCookies(req.URL, cookies)
	client.Jar = jar

	response, err := client.Do(req)
	if err != nil {
		return err
	}
	//, err := ioutil.ReadAll(response.Body)
	//fmt.Println(string(b))

	defer response.Body.Close()
	return nil
}
