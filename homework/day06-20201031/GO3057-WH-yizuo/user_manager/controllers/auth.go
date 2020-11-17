package controllers

import (
	"fmt"
	"yizuo/models"
	"yizuo/utils"
)

// 用户登录检测
func UserLoginAuth() bool {
	/*
		循环三次，对比用户输入的用户名密码是否正确
			密码正确返回 true
			密码错误返回 false
	*/
	for i := 0; i < 3; i++ {
		if i != 0 {
			fmt.Println(`密码输入错误，请重新输入密码：`)
		}
		user:=utils.Input(`请输入登录用户：`)
		passwd:=utils.Input(`请输入密码：`)


		// 对比输入值的md5是否与默认值的md5一致
		if models.UserLoginAuth(user,passwd) {
			return true
		}
	}
	return false
}
