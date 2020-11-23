package utils

import (
	"fmt"
)

func Input(prompt string) string {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	return text
}

func PrintUser(user map[string]string) {
	// fmt.Println(user)
	fmt.Printf("ID：%s, 名称：%s, 联系方式：%s, 通信地址：%s\n", user["id"], user["name"], user["addr"], user["tel"])
}

func PrintMsg() {
	fmt.Println("============================")
	fmt.Println("=========== 用户管理系统  ===========")
	fmt.Println("==      请根据以下提示信息操作       ==")

	fmt.Println("==== 查看用户：输入 all 后按回车 ====")

	fmt.Println("==== 添加用户：输入 add 后按回车 ====")

	fmt.Println("==== 删除用户：输入 del 后按回车 ====")

	fmt.Println("==== 修改用户：输入 modify 后按回车 ====")

	fmt.Println("==== 查找用户：输入 query 后按回车 ====")

	fmt.Println("==== 查看帮助：输入 help 后按回车  ==")

	fmt.Println("==== 退出系统：输入 quit 后按回车  ==")
}

func PrintUsers(users []map[string]string) {
	//打印用户信息
	fmt.Println("\n以下是目前已存在的用户：")
	for _, v := range users {
		fmt.Printf("ID：%s, 名称：%s, 联系方式：%s, 通信地址：%s\n", v["id"], v["name"], v["addr"], v["tel"])
	}
}
