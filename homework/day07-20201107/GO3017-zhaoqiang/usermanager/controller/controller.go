package controller

import (
	"fmt"
	"os"
	"strings"
	"zhao/models"
	"zhao/utils"
)

//PasswdAuth 用户输入 count 次失败则退出
func PasswdAuth(count int) bool {

	models.File.LoadUsers()

	for i := 0; i < count; i++ {
		if models.PasswdAuth() {
			return true
		}
		fmt.Printf("Permission denied,  trey again.\n\n")
	}
	return false
}

//Add 添加指令
func Add() {
	if err := models.File.LoadUsers(); err != nil {
		fmt.Println("[Add_LoadUsers false]", err)
		os.Exit(-1)
	}

	user, err := models.GetUserInfo("add")
	if err != nil {
		fmt.Println("[get user information]", err)
		return
	}
	if err := models.AddAuth(user); err != nil {
		fmt.Println("[user Auth]: ", err)
		return
	}

	if err := models.AddUser(user); err == nil {
		fmt.Printf("[+]user add sucessful\n\n")
	} else {
		fmt.Println("[Add_Storage false]", err)
		return
	}

}

//Query 查询指令
func Query() {
	if err := models.File.LoadUsers(); err != nil {
		fmt.Println("[Add_LoadUsers false]", err)
		os.Exit(-1)
	}

	querystr := utils.GetInput("input generic query message: ")
	u := models.GenericQuery(querystr)
	if len(u) == 0 {
		fmt.Printf("no matching fields\n\n")
		return
	}
	models.PrintUsers(u)
}

//Del 删除指令
func Del() {
	if err := models.File.LoadUsers(); err != nil {
		fmt.Println("[Add_LoadUsers false]", err)
		os.Exit(-1)
	}

	id := utils.GetInput("input want delete ID: ")
	user, _, err := models.IDQuery(id)
	if err != nil {
		fmt.Println("[Del IDQuery must inter]", err)
		return
	}
	//打印用户， 确认是否删除
	models.PrintUsers([]models.User{user})

	confirm := utils.GetInput("are sure delete this user(y/n): ")
	if strings.ToLower(confirm) == "y" {
		if err := models.DelUser(user); err != nil {
			fmt.Println("[Del_Storage false]", err)
			return
		}
		fmt.Printf("[+] delete user sucessful\n\n")
		return
	}
}

//ModefyUser 修改指令
func ModefyUser() {
	if err := models.File.LoadUsers(); err != nil {
		fmt.Println("[Add_LoadUsers false]", err)
		os.Exit(-1)
	}

	id := utils.GetInput("input want modify ID: ")
	user, userindex, err := models.IDQuery(id)
	if err != nil {
		fmt.Println("[modify IDQuery]", err)
		return
	}
	//登陆用户密码验证
	err = models.ModifypasswdAuther(user)
	if err != nil {
		fmt.Println("[ModifypasswdAuther]", err)
	}

	//打印用户， 确认是否修改
	models.PrintUsers([]models.User{user})

	confirm := utils.GetInput("are sure modify this user(y/n): ")
	if strings.ToLower(confirm) == "y" {
		userinputUserinfo, err := models.GetUserInfo("modify")
		if err != nil {
			fmt.Println("[ModifyUser]", err)
			return
		}
		// fmt.Println(userindex)
		if er := models.ModifyUser(userindex, userinputUserinfo); er == nil {
			fmt.Printf("[+] modify user sucessful\n\n")
			return
		} else {
			fmt.Println("[-] modify user faled!!", er)
			return
		}
	}
}

//PrintA 打印全部用户
func PrintA() {
	models.PrintAllusers()
}
