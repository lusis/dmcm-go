package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Endpoint string
	Version string
	Accesskey string
	SecretKey string
	UserAgent string
	Detail string
	VerifySsl bool
}

const ES_UA = "gostratus"

func GetTimeString() string {
	s := time.Now()
	now := s.Unix()*1000 + int64(s.Nanosecond()/1e6)
	return strconv.FormatInt(now, 10)
}

func hmacB64(message string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func SignRequest(accessKey string, secretKey string,
	userAgent string, httpMethod string,
	httpPath string, apiVersion string) string {
	signpath := strings.Join([]string{"/api/enstratus", apiVersion, httpPath}, "/")
	timestamp := GetTimeString()
	s := strings.Join([]string{accessKey, httpMethod, signpath, timestamp, userAgent}, ":")
	hash := hmacB64(s, secretKey)
	return hash
}
