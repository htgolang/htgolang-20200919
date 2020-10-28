package utils

import (
	"crypto/md5"
	"fmt"
)

// Md5Text return passwd string
func Md5Text(text string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(text)))
}

//PasswdStr md5转换后的密码
func PasswdStr() string {
	p := GetUserInputString("enter pass world")
	return Md5Text(p)
}
