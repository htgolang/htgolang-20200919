package controllers

import (
	"GO3004-zhaoweiping/models"
	"GO3004-zhaoweiping/utils"
	"fmt"
	"strings"
)

func AddUser() {
	user := map[string]string{
		"name": utils.Input("请输入用户名："),
		"addr": utils.Input("请输入联系地址："),
		"tel":  utils.Input("请输入联系方式："),
	}
	models.AddUser(user)
	fmt.Printf("用户【%v】添+成功！！！\n", user["name"])
}

func ModifyUser() {
	id := utils.Input("请输入需要修改的用户 ID：")
	user := models.FindUserById(id)
	if user == nil {
		fmt.Println("用户信息不存在！")
	} else {
		fmt.Println("你将要修改的用户信息如下：")
		utils.PrintUser(user)
		confirm := utils.Input("你确定修改吗？(Y/n)")
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

func DeleteUser() {
	id := utils.Input("请输入需要删除的用户 ID：")
	user := models.FindUserById(id)
	if user == nil {
		fmt.Println("用户信息不存在！")
	} else {
		fmt.Println("你将要删除的用户信息如下：")
		utils.PrintUser(user)
		confirm := utils.Input("确定删除吗?(Y/n)：")
		if strings.ToLower(confirm) == "y" || strings.ToLower(confirm) == "yes" {
			models.DeleteUserById(id)
		}
	}
}

func QueryUser() {
	q := utils.Input("请输入查询信息：")
	fmt.Println("查询结果")
	users := models.QueryUser(q)
	for _, user := range users {
		utils.PrintUser(user)
	}
}

func PringUsersAll() {
	users := models.ReturnUsers()
	utils.PrintUsers(users)
}

func PringHelpMsg() {
	utils.PrintMsg()
}
