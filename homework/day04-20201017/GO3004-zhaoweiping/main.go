package main

import (
	"GO3004-zhaoweiping/manager"
	_ "GO3004-zhaoweiping/routers"
	"GO3004-zhaoweiping/utils"
	"fmt"
)

var passwordMd5 = "c4ca4238a0b923820dcc509a6f75849b" //密码 1

func auth() bool {
	for i := 0; i < 3; i++ {
		if utils.Md5text(utils.Input("请输入密码：")) == passwordMd5 {
			return true
		} else {
			fmt.Println("输入密码错误：")
		}
	}
	return false
}

func main() {
	if !auth() {
		fmt.Println("密码输入错误，程序退出")
		return
	}
	utils.PrintMsg()
	manager.Run()
}
