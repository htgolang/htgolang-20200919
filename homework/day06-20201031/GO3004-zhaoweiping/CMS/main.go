package main

import (
	"CMS/manager"
	"CMS/models"
	_ "CMS/routers"
	"CMS/utils"
	"fmt"
)

func main() {
	if !models.Auth() {
		fmt.Println("账号密码输入错误，程序退出！！！")
		return
	}
	utils.PrintMsg()
	manager.Run()
}
