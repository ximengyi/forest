package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func GetMd5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Encode(data string) (string, error) {
	h := md5.New()
	_, err := h.Write([]byte(data))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func ParseServerAddr(serverAddr string) (host, port string) {
	serverInfo := strings.Split(serverAddr, ":")
	if len(serverInfo) == 2 {
		host = serverInfo[0]
		port = serverInfo[1]
	} else {
		host = serverAddr
		port = ""
	}
	return host, port
}

func InArrayString(s string, arr []string) bool {
	for _, i := range arr {
		if i == s {
			return true
		}
	}
	return false
}

//Substr 字符串的截取
func Substr(str string, start int64, end int64) string {
	length := int64(len(str))
	if start < 0 || start > length {
		return ""
	}
	if end < 0 {
		return ""
	}
	if end > length {
		end = length
	}
	return string(str[start:end])
}