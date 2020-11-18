package controllers

import (
	"fmt"
	"strings"
	"yizuo/models"
	"yizuo/utils"
)
/*
用户相关的所有操作在此文件中进行
*/

// 添加用户
func AddUser() {
	/*
		根据用户数据的内容来添加用户
	*/
	name := utils.Input("请输入用户名: ")
	// 判断用户是否存在，存在即退出
	if models.FindElementUser(name) {
		fmt.Printf("用户%v已存在，已退出。\n",name)
		return
	}
	passwd := utils.Input("请输入用户密码:")
	phone := utils.Input("请输入联系方式（例如：17612345678）:")
	address := utils.Input("请输入联系地址（例如：WuHan）:")
	birthday := utils.Input("请输入你的生日（例如：1994-04-06 18:08:08）:")

	models.AddUser(name,passwd,phone,address,birthday)
	models.WritesUsersDataToCsv()
	fmt.Println("用户添加成功，现所有用户如下：")
	ListUser()
}

// 删除用户（软删除）
func DeleteUser(){
	/*
	根据用户输入的用户ID，将Status调整为0，软删除数据，打印时Status为0的数据将不再显示。
	 */
	// 获取用户输入的用户ID
	UserID := utils.IntInput("请输入要删除的用户ID：")

	// 在数据中检索查询ID是否存在
	if models.FindElementID(UserID) {
		if models.DeleteUser(UserID) {
			models.WritesUsersDataToCsv()
			fmt.Println("用户删除成功！")
		} else {
			fmt.Println("用户删除失败！")
		}
	}  else {
		fmt.Println("未找到该用户ID，已退出")
	}
}

// 修改用户数据
func ModifyUser()  {
	/*
		用户输入用户ID，根据用户ID变更对应的条目数据
	*/
	// 获取用户输入的用户ID
	UserID := utils.IntInput("请输入要变更的用户ID：")

	// 在数据中检索查询ID是否存在
	if models.FindElementID(UserID) {
		if !models.ModifyUser(UserID) {
			fmt.Println("用户信息变更失败！")
		}
		models.WritesUsersDataToCsv()
	}   else  {
		fmt.Println("未找到此ID对应的用户数据，已退出。")
	}
}

// 根据用户输入的信息查询ID/用户名称是否存在
func QueryUser()  {
	// 获取用户输入的用户名称
	UserName := utils.Input("请输入要查询的用户名称:")

	u,ok := models.QueryUser(UserName)
	if ok != nil {
		fmt.Println("查询用户失败。")
	}
	models.FormatTableOut(u)
}

// 打印所有状态为1的用户列表
func ListUser()  {
	u := models.ListUser()
	models.FormatListTableOut(u)
}

// 初始化系统
func InitAllUser()  {
	opt := strings.ToLower(utils.Input("你确定要初始化吗，初始化会删除现有所有数据。(yes/no):"))
	if opt == "y" || opt == "yes"{
		models.InitAllUser()
		fmt.Println("用户初始化，成功。")
	} else {
		fmt.Println("用户初始化，取消。")
	}
}