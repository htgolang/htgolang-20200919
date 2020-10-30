package controller

import (
	"fmt"
	"strconv"
	"usermanager/users"
	"usermanager/userutils"
)

func Run(udb *users.Userdb){
	quit := auth(udb)
	for quit {
		header := []string{"序号","功能",}
		data := [][]string{{"1","添加用户."},{"2","删除用户."},{"3","修改用户."},{"4","查找用户."},{"5","退出系统."}}
		userutils.Showintable(header,data)
		s := userutils.Input("请输入序号(1~5):")
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
