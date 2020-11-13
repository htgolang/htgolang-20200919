package models

import (
	"fmt"
	"os"
	"strings"
	"zhao/utils"
)

//LoginUser recode login status
var LoginUser User

//PasswdAuth 登陆用户密码认证
func PasswdAuth() bool {
	username, passwd := cmdGetNameAndPasswd()
	for _, user := range users {
		if user.Name == username && user.Passwd == passwd {
			LoginUser = user
			return true
		}
	}
	return false
}

func cmdGetNameAndPasswd() (string, string) {
	var (
		username string
		passwd   string
	)
	for {
		//用户名
		for {
			username = utils.GetInput("Login: ")
			if strings.TrimSpace(username) == "" {
				continue
			}
			break
		}
		//密码
		p, err := utils.GetPasswd("Passwd: ")
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		passwd = utils.Md5Convert([]byte(p))
		return username, passwd
	}
}
