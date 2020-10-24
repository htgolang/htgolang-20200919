package controllers

import (
	"fmt"
	"strings"

	"github.com/imsilence/user/models"
	"github.com/imsilence/user/utils"
)

// 添加用户
func AddUser() {
	// 用户输入
	user := map[string]string{
		"name": utils.Input("请输入用户名: "),
		"addr": utils.Input("请输入联系地址: "),
		"tel":  utils.Input("请输入联系方式: "),
	}

	// 对数据添加
	models.AddUser(user)

	// 用户响应
	fmt.Println("[+]添加用户成功")
}

func ModifyUser() {
	id := utils.Input("请输入需要修改的用户ID: ")
	user := models.FindUserById(id)
	if user == nil {
		fmt.Println("用户信息不存在")
	} else {
		fmt.Println("你将要修改的用户信息如下: ")
		fmt.Println(user)
		confirm := utils.Input("确定修改吗?(Y/n)")
		if strings.ToLower(confirm) == "y" || strings.ToLower(confirm) == "yes" {
			user := map[string]string{
				"id":   id,
				"name": utils.Input("name: "),
				"addr": utils.Input("addr: "),
				"tel":  utils.Input("tel: "),
			}
			models.ModifyUserById(user, id)
		}
	}
}

// 删除用户信息
func DeleteUser() {
	id := utils.Input("请输入需要删除的用户ID: ")
	user := models.FindUserById(id)
	if user == nil {
		fmt.Println("用户信息不存在")
	} else {
		fmt.Println("你将要删除的用户信息如下: ")
		fmt.Println(user)
		confirm := utils.Input("确定删除吗?(Y/n)")
		if strings.ToLower(confirm) == "y" || strings.ToLower(confirm) == "yes" {
			models.DeleteUserById(id)
		}
	}
}

// 打印用户数据
func printUser(user map[string]string) {
	fmt.Println(user)
}

// 查询用户
func QueryUser() {
	// 输入
	q := utils.Input("请输入查询信息: ")
	fmt.Println("查询结果")

	// 调用models对数据过滤
	users := models.QueryUser(q)

	for _, user := range users {
		printUser(user)
	}
}
