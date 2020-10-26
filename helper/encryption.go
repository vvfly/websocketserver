package helper

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"strconv"

	"github.com/forgoer/openssl"
	. "github.com/luckyweiwei/base/logger"
	"github.com/luckyweiwei/base/utils"
)

const (
	des3Iv  = "01234567"
	des3Key = "c21d31be-4300-4881-a553-156ebb5df087"

	userXorKey         = 0x9034
	timeStampKey int64 = 0x0938431251641621
)

// base64加解密

// base64解密
func DecodeBase64(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// base64加密
func EncodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// 3des加解密

// @title    函数名称
// @description   函数的详细描述
// @auth      作者             时间（2020/6/4   10:57 ）

// @param     输入参数名        参数类型         "解释"
// @return    返回参数名        参数类型         "解释"
func Des3CBCDecrypt(s []byte, key string) ([]byte, error) {
	return openssl.Des3CBCDecrypt(s, []byte(key)[:24], []byte(des3Iv), openssl.PKCS7_PADDING)
}

func DesCBCDecrypt(s []byte) ([]byte, error) {
	return openssl.DesCBCDecrypt(s, []byte(des3Key)[:8], []byte(des3Iv), openssl.PKCS5_PADDING)
}

func DesCBCEncrypt(s []byte) ([]byte, error) {
	return openssl.DesCBCEncrypt(s, []byte(des3Key)[:8], []byte(des3Iv), openssl.PKCS5_PADDING)
}

// 异或加解密
func XorStrInteger(ciphertext string) (txt string) {
	for _, c := range ciphertext {
		txt += string(userXorKey ^ c)
	}
	return txt
}

func Xor4Timestamp(timestamp string) int64 {
	tm, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		Log.Error(err)
		return 0
	}

	return tm ^ timeStampKey
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
