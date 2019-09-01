package vpn

import (
	"regexp"
	"testing"
)

func TestRep(t *testing.T) {
	var text1 = `var logoutOtherToken = 'e97e5e358c2713c2'`
	matched_1, err := regexp.Match(`logoutOtherToken[\s]+=[\s]+'[\w]+`, []byte(text1))
	if !matched_1 {
		t.Error(err)
	}

	var text2 = `var logoutOtherToken = 'e97e5e358c2713c2'  \n`
	matched_2, err := regexp.Match(`logoutOtherToken[\s]+=[\s]+'[\w]+`, []byte(text2))
	if !matched_2 {
		t.Error(err)
	}

	var text3 = `var logoutOtherToken = ''  \n`
	matched_3, err := regexp.Match(`logoutOtherToken[\s]+=[\s]+'[\w]+`, []byte(text3))
	if matched_3 {
		t.Error(err)
	}
}
