package util

import (
	"crypto/md5"
)

// MD5 : md5算法加密
func MD5(rawData []byte) string {
	enc := md5.New()
	enc.Write(rawData)
	result := enc.Sum([]byte(""))
	return string(result)
}
