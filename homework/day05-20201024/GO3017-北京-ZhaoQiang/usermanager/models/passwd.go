package models

import (
	"fmt"
	"zhao/utils"
)

//AuthenPassWD authentication account and password it true
func AuthenPassWD() (string, bool) {
	for {
		usercount := utils.GetUserInputString("user login: ")
		if usercount == "" {
			fmt.Printf("\n")
			continue
		}
		userpasswd := utils.GetUserInputPassWD("PassWorld: ")

		passwd := utils.Md5Text(userpasswd)

		for _, user := range users {
			if usercount == user.name && passwd == user.passwd {
				return usercount, true
			}
		}
		return "", false
	}
}
