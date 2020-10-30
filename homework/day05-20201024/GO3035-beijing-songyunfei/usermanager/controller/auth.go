package controller

import (
	"fmt"
	"usermanager/users"
	"usermanager/userutils"
)

func auth(udb *users.Userdb) bool {
	//限制次数
	retry := 3
	for p := 0; p <retry; p++{
		username := userutils.Input("用户名:")
		passwd := userutils.Input("密码:")
		if  udb.Auth(username,passwd) {
			fmt.Println("欢迎进入系统.")
			return true

		}else {
			fmt.Println("用户名或密码错误....")
		}
	}
	fmt.Printf("%d次输入错误...",retry)
	return  false
}
