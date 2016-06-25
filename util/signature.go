package util

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
	"strings"
	"time"
	"crypto/md5"
	"fmt"
	"encoding/hex"
)

//CreateSignature creates signature for string following Aliyun rules
func CreateSignature(stringToSignature, accessKeySecret string) string {
	// Crypto by HMAC-SHA1
	hmacSha1 := hmac.New(sha1.New, []byte(accessKeySecret))
	hmacSha1.Write([]byte(stringToSignature))
	sign := hmacSha1.Sum(nil)

	// Encode to Base64
	base64Sign := base64.StdEncoding.EncodeToString(sign)

	return base64Sign
}

func percentReplace(str string) string {
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)

	return str
}

// CreateSignatureForRequest creates signature for query string values
func CreateSignatureForRequest(method string, values *url.Values, accessKeySecret string) string {

	canonicalizedQueryString := percentReplace(values.Encode())

	stringToSign := method + "&%2F&" + url.QueryEscape(canonicalizedQueryString)

	return CreateSignature(stringToSign, accessKeySecret)
}

// CreateSignatureForStreamUrlWithA creates signature for Url string whit method A
func CreateSignatureForStreamUrlWithA(uri, rand, uid, privateKey string, timeout time.Duration) (authKey string, timestamp int64) {
	//timestamp for timeout
	timestamp = time.Now().Add(timeout).Unix()
	//Signature string
	sstring := fmt.Sprintf("%s-%d-%s-%s-%s", uri, timestamp, rand, uid, privateKey)
	//Crypto by HMAC-MD5
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(sstring))
	// Encode to Hex
	hashValue := hex.EncodeToString(md5Ctx.Sum(nil))
	authKey = fmt.Sprintf("%d-%s-%s-%s", timestamp, rand, uid, hashValue)
	return
}