package models

import (
	"GO3004-zhaoweiping/utils"
	"fmt"
)

func Auth() bool {
	fmt.Println("程序运行成功，请使用账号密码登录本系统：")
	for i := 0; i < 3; i++ {
		NameTemp := utils.Input("请输入账号：")
		passwordMd5Temp := utils.Md5text(utils.Input("请输入密码："))
		// fmt.Println(NameTemp, passwordMd5Temp)
		for k, user := range UsersDb {
			if NameTemp == k && passwordMd5Temp == user.passwd {
				return true
			} else {
				fmt.Printf("账号密码不对，请重新输入！！！\n")
				break
			}
		}
	}
	return false
}
