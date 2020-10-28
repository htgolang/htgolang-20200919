package controller

import (
	"fmt"
	"strings"
	"zhao/models"
	"zhao/utils"
)

//PrintAll 打印用户信息
func PrintAll() {
	models.PrintAll()
}

//AddRun adduser run functions
func AddRun() {
	user, err := models.GetUserMessage()
	if err == nil {
		if models.Authentication(user) {
			models.AddUser(user)
			fmt.Printf("[+]user add sucessful\n\n")
		} else {
			fmt.Printf("[-] user add fall ,already username\n\n")
		}
	} else {
		fmt.Println(err)
	}
}

//DelRun adduser run functions
func DelRun() {
	id := utils.GetUserInputInt("enter want delete ID: ")

	user, err := models.FindUserByID(id)
	if err != nil {
		fmt.Printf("user not exist\n\n")
	} else {
		models.Printfunc([]models.User{user})

		confirm := utils.GetUserInputString("are sure delete user (y/n):")
		if strings.ToLower(strings.TrimSpace(confirm)) == "y" {
			models.DelUser(user)
			fmt.Println("[+] delete user sucess")
		}
	}

}

//QueryRun 查询用户信息
func QueryRun() {
	order := utils.GetUserInputString("enter search message: ")
	u := models.Query(order)
	if len(u) == 0 {
		fmt.Printf("user not exist \n\n")
		return
	}
	models.Printfunc(u)
}

// ModifyRun 修改用户信息
func ModifyRun() {
	id := utils.GetUserInputInt("enter want modify ID: ")
	user, err := models.FindUserByID(id)
	if err != nil {
		fmt.Printf("%v\n\n", err)
		return
	}

	models.Printfunc([]models.User{user})

	confirm := utils.GetUserInputString("are sure modify user (y/n):")
	if strings.ToLower(strings.TrimSpace(confirm)) == "y" {

		modifyUser, err := models.GetModifyUserMessage()
		if err != nil {
			return
		}
		if err := models.ModifyUserMessageByID(user, modifyUser); err == nil {
			fmt.Printf("[+] user modify sucess\n\n")
		} else {
			fmt.Printf("%v\n\n", err)
		}
	}

}
