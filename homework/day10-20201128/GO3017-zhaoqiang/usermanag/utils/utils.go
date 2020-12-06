package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5Convert(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}
