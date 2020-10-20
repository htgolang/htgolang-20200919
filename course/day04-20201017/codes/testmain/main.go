package main

import (
	"fmt"
	_ "testmain/pkg"
)

/*
	程序入口
	开发者写代码的执行入口
*/
func main() {

}

// 初始化函数
// 导入包时执行
func init() {
	fmt.Println("main.init")
}
