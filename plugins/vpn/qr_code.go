package vpn

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/skip2/go-qrcode"
	"gopkg.in/yaml.v3"
	"io"
	"net/http"
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
	imgUrl string
	QrImg  []byte
	Sid    string // sis in ustb auth, can be parsed from image url.
}

// ParseQRCodeImgUrl uses ParseQRCodeHtmlUrl to get the iframe html,
// and then parse the html file to get final image url (contains SID).
// And set QrImg's imgUrl and sid.
func (i *QrImg) ParseQRCodeImgUrl() (string, error) {
	iframeUrl, err := ParseQRCodeHtmlUrl()
	if err != nil {
		return "", err
	}
	htmlUri, err := url.Parse(iframeUrl)
	if err != nil {
		return "", err
	}

	// make a http request of the iframe
	response, err := http.Get(iframeUrl)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, FindQrcodeImgTagRegex) { // first line to match start
			// line e.g. <img id="qrimg" src="/connect/qrimg?sid=3894c5568dd1ef0f6434f426297a678d" height="90%" border="0">
			subStr := strings.SplitN(line, "\"", 5)
			if len(subStr) != 5 {
				return "", errors.New("invalid format in qr image url parsing")
			} else {
				i.imgUrl = subStr[3]
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	// parse sid in qr image url.
	sidStr := strings.SplitN(i.imgUrl, "=", 2)
	if len(sidStr) != 2 {
		return "", errors.New("invalid format in qr image url (sis) parsing")
	} else {
		i.Sid = sidStr[1]
	}

	// use htmlUri's host, schema
	return fmt.Sprintf("%s://%s%s", htmlUri.Scheme, htmlUri.Host, i.imgUrl), nil
}

func ParseQRCodeHtmlUrl() (string, error) {
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
	response, err := http.Get(LoadImgUrl)
	if err != nil {
	}

	defer response.Body.Close()

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
		return "", err
	}

	fmt.Println("parsed qr code config:", qrConfigBuffer.String())
	var qrImgUrlConfig QRCodeImgLoaderConfig
	if err := yaml.Unmarshal(qrConfigBuffer.Bytes(), &qrImgUrlConfig); err != nil {
		return "", err
	}

	// generate iframe url.
	if qrUrl, err := qrImgUrlConfig.genIframeUrl(); err != nil {
		return "", err
	} else {
		return qrUrl, nil
	}
}

func (i *QrImg) GenQrImg() error {
	imgContent := fmt.Sprintf(SisAuthPath+"/auth?sid=%s", i.Sid)
	qrPng, err := qrcode.Encode(imgContent, qrcode.Medium, 256)
	if err != nil {
		return err
	}
	i.QrImg = qrPng
	return nil
}

func LoadQrAuthImage() (io.Reader, error) {
	var qr QrImg
	if _, err := qr.ParseQRCodeImgUrl(); err != nil {
		return nil, err
	}

	if err := qr.GenQrImg(); err != nil {
		return nil, err
	}
	buf := bytes.NewReader(qr.QrImg)
	return buf, nil
}

// WaitQrState waits qr state and get auth code
func WaitQrState(imgUrl *url.URL) {
	// get https://sis.ustb.edu.cn/connect/state?sid=bf1a027b75d6e21b351f81cdc1b739a2
}