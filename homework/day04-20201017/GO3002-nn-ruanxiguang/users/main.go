package main

import (
	"fmt"
	"users/manager"
	_ "users/routers"
	"users/utils"
)

var password = "202cb962ac59075b964b07152d234b70"

func auth() bool {
	for i := 0; i < 3; i++ {
		if utils.Md5text(utils.Input("请输入密码：")) == password {
			return true
		} else {
			fmt.Println("密码输入错误")
		}
	}
	return false
}

func main() {
	if !auth() {
		fmt.Println("密码输入错误，程序退出")
		return
	}
	manager.Run()
}
