package main

import (
	"GO3004-zhaoweiping/manager"
	"GO3004-zhaoweiping/models"
	_ "GO3004-zhaoweiping/routers"
	"GO3004-zhaoweiping/utils"
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
