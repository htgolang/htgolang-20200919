package main

import (
	"yizuo/manager"
	_ "yizuo/routers"
)

// 主程序
func main() {
	/*
	   主执行函数
	   初始用户密码如下：
	      yizuo yizuo
	*/
	manager.Run()
}
