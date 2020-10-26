package helper

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"

	"github.com/forgoer/openssl"
	"github.com/luckyweiwei/base/utils"
)

const (
	des3Iv = "01234567"
)

// base64解密
func DecodeBase64(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// base64加密
func EncodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Des3CBCDecrypt4WebsocketMsg(key, s []byte) ([]byte, error) {
	defer utils.CatchException()

	dst, err := DecodeBase64(string(s))
	if err != nil {
		return nil, err
	}

	return openssl.Des3CBCDecrypt(dst, []byte(key)[:24], []byte(des3Iv), openssl.PKCS5_PADDING)
}

func Des3CBCEncrypt4WebsocketMsg(key, s []byte) ([]byte, error) {
	defer utils.CatchException()

	d, err := openssl.Des3CBCEncrypt([]byte(s), []byte(key)[:24], []byte(des3Iv), openssl.PKCS5_PADDING)
	if err != nil {
		return nil, err
	}

	dst := EncodeBase64(d)

	return []byte(dst), nil
}

func MD5(in string) string {

	alg := md5.New()
	alg.Write([]byte(in))

	return hex.EncodeToString(alg.Sum(nil))
}
