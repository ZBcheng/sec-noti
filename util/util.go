package util

import (
	"crypto/md5"
)

// MD5 : md5算法加密
func MD5(rawData string) string {
	enc := md5.New()
	enc.Write([]byte(rawData))
	result := enc.Sum([]byte(" "))
	return string(result)
}
