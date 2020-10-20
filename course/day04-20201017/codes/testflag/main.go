package main

import (
	"flag"
	"fmt"
)

func main() {
	// ssh -P port[22] root@host
	var port int
	var help bool
	var password string
	// var xxx string

	// 指定变量与命令行参数(通过参数名称)中的关系
	flag.IntVar(&port, "P", 22, "port")
	flag.BoolVar(&help, "h", false, "help")
	flag.StringVar(&password, "p", "", "password")

	flag.Usage = func() {
		fmt.Println("usage: ssh -P 22 root@locahost")
		flag.PrintDefaults()
	}

	// 解析命令行参数
	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	fmt.Println(port, password)

	// 未指定参数名称的参数列表
	fmt.Println(flag.Args())
}
