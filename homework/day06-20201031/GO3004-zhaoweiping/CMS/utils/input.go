package utils

import (
	"fmt"
	"time"

	"github.com/spf13/cast"
)

func Input(prompt string) string {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	return text
}
func InputTime(prompt string) time.Time {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	ret, _ := time.Parse("2006-01-02 15:04:05", text)
	return ret
}
func InputInt(prompt string) int {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	text1 := cast.ToInt(text)
	return text1
}

func PrintMsg() {
	fmt.Println("============================")
	fmt.Println("=========== 用户管理系统  ===========")
	fmt.Println("==      请根据以下提示信息操作       ==")

	fmt.Println("==== 查看用户：输入 all 或者 1 后按回车 ====")

	fmt.Println("==== 添加用户：输入 add 或者 2 后按回车 ====")

	fmt.Println("==== 删除用户：输入 modify 或者 3  后按回车 ====")

	fmt.Println("==== 修改用户：输入 del 或者 4 后按回车 ====")

	fmt.Println("==== 查找用户：输入 query 或者 5  后按回车 ====")

	fmt.Println("==== 查看帮助：输入 help 或者 6 后按回车  ==")

	fmt.Println("==== 退出系统：输入 quit 后按回车  ==")
}
