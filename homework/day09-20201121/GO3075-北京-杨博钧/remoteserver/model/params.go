package model

import (
	"flag"
	"fmt"
)
//客户端请求参数结构体
type Params struct{
	Cmd string `json:"cmd"`
	Path string `json:"path"`
	Src string `json:"src"`
	Dest string `json:"dest"`
}

// 定义设置path函数
func (this *Params) SetPath() {
	fmt.Print("请输入要操作的路径:")
	var path string
	fmt.Scan(&path)
	this.Path = path
}

// 定义设置src和dest函数
func (this *Params) SetSrcDest() {
	fmt.Print("请输入源路径")
	var src string
	fmt.Scan(&src)
	this.Src = src
	fmt.Println()
	fmt.Print("请输入目标路径")
	var dest string
	fmt.Scan(&dest)
	this.Dest = dest
}

// client端参数工厂函数
func NewParams() *Params {
	var params Params
	flag.StringVar(&params.Cmd, "cmd", "", "需要远程执行的命令,ls、put、get、rm")
	flag.Parse()
	switch params.Cmd {
	case "ls", "rm":
		params.SetPath()
	case "put", "get":
		params.SetSrcDest()
	default:
		fmt.Println("输入的操作参数不合规!请用--help查看可执行命令")
		return nil
	}
	return &params
}
