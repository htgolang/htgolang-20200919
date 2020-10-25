package controllers

import (
	"fmt"
	"strings"
	"users/models"
	"users/utils"
)

// 添加用户
func AddUser() {
	// 用户输入
	user := map[string]string{
		"name": utils.Input("请输入用户名："),
		"addr": utils.Input("请输入联系地址："),
		"tel":  utils.Input("请输入联系方式："),
	}

	// 添加数据
	models.AddUser(user)

	//响应
	models.PrintUser(user)
	fmt.Println("[+]添加用户成功")
}

// 删除用户信息
func DeleteUser() {
	id := utils.Input("请输入需要删除的用户ID：")
	user := models.FindUserById(id)
	if user == nil {
		fmt.Println("不存在此用户")
	} else {
		fmt.Println("即将删除的用户信息如下：")
		models.PrintUser(user)
		confirm := utils.Input("确定删除吗?(Y/N)")
		if strings.ToLower(confirm) == "y" || strings.ToLower(confirm) == "yes" {
			models.DeleteUserById(id)
		}
	}
}

// 修改用户信息
func ModifyUser() {
	id := utils.Input("请输入需要修改的用户ID：")
	user := models.FindUserById(id)
	if user == nil {
		fmt.Println("无此用户信息")
	} else {
		fmt.Println("将要修改的用户信息如下：")
		models.PrintUser(user)
		confirm := utils.Input("确定修改吗(Y/N)")
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

// 查询用户信息
func QueryUser() {
	q := utils.Input("请输入查询信息：")
	fmt.Println("查询结果")

	// 调用models对数据过滤
	users := models.QueryUser(q)

	models.PrintUsers(users)
}
