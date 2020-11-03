package controller

import (
	"fmt"
	"strconv"
	"usermanager/users"
	"usermanager/userutils"
)

func Run(udb *users.Userdb){
	var quit bool
	p := 0
	//验证3次密码
	for p <3{
		username := userutils.Input("用户名:")
		passwd := userutils.Input("密码:")
		if  udb.Auth(username,passwd) {
			quit = true
			fmt.Println("欢迎进入系统.")
			break
		}else {
			fmt.Println("密码错误....")
		}
		p++
	}
	if p == 3 && quit == false {
		fmt.Println("3次输入错误...")
	}
	for quit {
		s := userutils.Input("1.添加用户.\n2.删除用户.\n3.修改用户.\n4.查找用户.\n5.退出\n请输入序号(1~5):")
		se,_ := strconv.Atoi(s)
		switch se {
		// 增加用户
		case 1:
			add(udb)
		// 删除用户
		case 2:
			delUser(udb)
		//// 修改用户
		case 3:
			modifyUser(udb)
		// 查找用户
		case 4:
			queryuser(udb)
		// 退出
		case 5:
			fmt.Println("Bey...")
			quit = false

		default:
			fmt.Println("输入错误...")
		}
	}
}
