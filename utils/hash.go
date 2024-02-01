package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(text string) string {
	md5 := md5.New()
	md5.Write([]byte(text))
	digest := md5.Sum(nil)
	return hex.EncodeToString(digest)
}
