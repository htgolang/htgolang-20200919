package userutils

import (
	"crypto/md5"
	"fmt"
)

func Summd5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return fmt.Sprintf("%X",ctx.Sum(nil))
}
