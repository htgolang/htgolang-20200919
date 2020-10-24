package main

import (
	"fmt"

	"github.com/imsilence/user/manager"
	_ "github.com/imsilence/user/routers"
	"github.com/imsilence/user/utils"
)

var password = "88cdb337f8c62dc69c1aee4066f80bf5"

func auth() bool {
	for i := 0; i < 3; i++ {
		if utils.Md5text(utils.Input("请输入密码: ")) == password {
			return true
		} else {
			fmt.Println("输入密码错误")
		}
	}
	return false
}

func main() {
	if !auth() {
		fmt.Println("密码输入错误, 程序退出")
		return
	}
	manager.Run()
}
