// .go go源文件后缀

// package 声明包
// package 包名(main) 声明main包
// 每一个源文件都需要有属于的一个包
// main => 声明这个包编译程可执行二进制程序
package main

// import 导入包
// 导入fmt包
import (
	"fmt"
)

// 定义main函数
// func 声明函数
// main函数 => 二进制程序的执行入口
func main() {

	// 调用函数 包名.函数名
	// 函数名 GO内置的函数, 当前包自定的函数
	// fmt go内置的包
	// Println 在控制台输出
	fmt.Println("hello world!!")
}
