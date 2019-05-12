package vpn_plugin

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestHostEncrypt(t *testing.T) {
	const key = "wrdvpnisthebest!"
	var aes_e = newAesEncrypt(key)
	s, _ := aes_e.Encrypt("console.hpc.gensh.me")
	if hex.EncodeToString(s) != "f3f84f8f283c6d1e76188ae29f502d2667c3c311" {
		t.Error("aes result is  not as expected, ", hex.EncodeToString(s))
	}

	fmt.Println(hex.EncodeToString([]byte(key)))

	var aes_e_2 = newAesEncrypt(key)
	s_2, _ := aes_e_2.Encrypt("proxy.gensh.me")
	if hex.EncodeToString(s_2) != "e0e54e843e7e6f55701b81e29550" {
		t.Error("aes result is  not as expected, ", hex.EncodeToString(s_2))
	}
}
