package models

import (
	"CMS/utils"
	"fmt"
)

func Auth() bool {
	fmt.Println("程序运行成功，请使用账号密码登录本系统：")
	for i := 0; i < 3; i++ {
		NameTemp := utils.Input("请输入账号：")
		passwordMd5Temp := utils.Md5text(utils.Input("请输入密码："))
		for k, user := range UsersDb {
			user1 := user.Passwd
			// fmt.Println(user1)
			if NameTemp == k && passwordMd5Temp == user.Passwd {
				fmt.Println(NameTemp, k)
				fmt.Println(passwordMd5Temp, user1)
				return true
			}
			// else {
			// 	fmt.Printf("账号密码不对，请重新输入！！！\n")
			// 	break
			// }
		}
	}
	return false
}
