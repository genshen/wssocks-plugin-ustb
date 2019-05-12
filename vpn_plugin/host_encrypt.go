package vpn_plugin

import (
	"crypto/aes"
	"crypto/cipher"
)

type AesEncrypt struct {
	key []byte
}

func newAesEncrypt(key string) (aes *AesEncrypt) {
	var en = AesEncrypt{key: []byte(key)}
	return &en
}

func (encrypt *AesEncrypt) Encrypt(strMesg string) ([]byte, error) {
	var iv = []byte(encrypt.key)[:aes.BlockSize]
	encrypted := make([]byte, len(strMesg))
	aesBlockEncrypter, err := aes.NewCipher(encrypt.key)
	if err != nil {
		return nil, err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(encrypted, []byte(strMesg))
	return encrypted, nil
}
