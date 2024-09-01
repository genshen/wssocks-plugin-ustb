package vpn

import (
	"fmt"
	"testing"
)

func TestQRCodeHtmlUrl(t *testing.T) {
	u, err := ParseQRCodeHtmlUrl()
	if err != nil {
		t.Error("error in loading qr code html url:", err)
	}
	fmt.Println(u)
}

func TestQRCodeImgUrl(t *testing.T) {
	u, err := ParseQRCodeImgUrl()
	if err != nil {
		t.Error("error in loading qr code img url:", err)
	}
	fmt.Println(u)
}
