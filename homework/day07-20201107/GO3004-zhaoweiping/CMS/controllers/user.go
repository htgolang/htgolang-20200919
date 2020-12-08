package controllers

import (
	"CMS/models"
	"CMS/utils"
	"fmt"
	"os"
	"strings"
	"time"
)

type users struct {
	id       int
	name     string
	addr     string
	tel      string
	birthday time.Time
	passwd   string
}

var DbType string

var user models.Users

func AddUserInit() bool {
	user1 := users{
		id:       0,
		name:     utils.Input("请输入用户名："),
		addr:     utils.Input("请输入联系地址："),
		tel:      utils.Input("请输入联系方式："),
		birthday: utils.InputTime("请输入生日日期："),
		passwd:   utils.Md5text(utils.Input("请输入密码：")),
	}
	user2 := models.NewUser(user1.id, user1.name, user1.addr, user1.tel, user1.birthday, user1.passwd)
	if user2.AddUser() {
		SaveDb()
	}
	return true
}

func AddUser() {
	user1 := users{
		id:       0,
		name:     utils.Input("请输入用户名："),
		addr:     utils.Input("请输入联系地址："),
		tel:      utils.Input("请输入联系方式："),
		birthday: utils.InputTime("请输入生日日期："),
		passwd:   utils.Md5text(utils.Input("请输入密码：")),
	}
	user2 := models.NewUser(user1.id, user1.name, user1.addr, user1.tel, user1.birthday, user1.passwd)
	if user2.AddUser() {
		SaveDb()
	}
}

func DeleteUser() {
	id := utils.InputInt("请输入需要删除的用户 ID：")
	user := user.FindUserById(id)
	// fmt.Println(user)
	if user == nil {
		fmt.Println("用户信息不存在！")
	} else {
		fmt.Println("你将要删除的用户信息如下：")
		user.PrintUser()
		confirm := utils.Input("确定删除吗?(Y/n)：")
		if strings.ToLower(confirm) == "y" || strings.ToLower(confirm) == "yes" {
			user.DeleteUserById()
		}
	}
	SaveDb()
}

func ModifyUser() {
	id := utils.InputInt("请输入需要修改的用户 ID：")
	user := user.FindUserById(id)
	if user == nil {
		fmt.Println("用户信息不存在！")
	} else {
		fmt.Println("你将要修改的用户信息如下：")
		user.PrintUser()
		confirm := utils.Input("你确定修改吗？(Y/n)")
		if strings.ToLower(confirm) == "y" || strings.ToLower(confirm) == "yes" {
			user1 := users{
				id:       id,
				name:     utils.Input("请输入用户名："),
				addr:     utils.Input("请输入联系地址："),
				tel:      utils.Input("请输入联系方式："),
				birthday: utils.InputTime("请输入生日日期："),
				passwd:   utils.Md5text(utils.Input("请输入密码：")),
			}
			user2 := models.NewUser(user1.id, user1.name, user1.addr, user1.tel, user1.birthday, user1.passwd)
			user2.ModifyUserById()
			SaveDb()
		}
	}

}

func QueryUser() {
	q := utils.Input("请输入查询信息：")
	fmt.Println("查询结果")
	users := models.QueryUser(q)
	for _, user := range users {
		user.PrintUser()
	}
}

func PringUsersAll() {
	models.PrintUsersDb()
}

func PringHelpMsg() {
	utils.PrintMsg()
}

func QuitServer() {
	if DbType == "csv" {
		models.DbToCsv()
		// fmt.Println("1")
	} else if DbType == "gob" {
		models.DbToGob()
	} else if DbType == "json" {
		models.DbToJson()
	}
	fmt.Println("结束本次服务，再见客官！！！")
	os.Exit(1)
}

func SaveDb() {
	confirm := utils.Input("请选择是否保存数据?(Y/n)：")
	if strings.ToLower(confirm) == "y" || strings.ToLower(confirm) == "yes" {
		if DbType == "csv" {
			models.DbToCsv()
			// fmt.Println("1")
		} else if DbType == "gob" {
			models.DbToGob()
		} else if DbType == "json" {
			models.DbToJson()
		}

	} else {
		fmt.Println("请继续使用")
		return
	}
}
