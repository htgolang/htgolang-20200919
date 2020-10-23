package main

import (
	"GO3032-SH-chenjing/utils"
	"crypto/md5"
	"fmt"
)


func main() {
	var (
		loginNum int
		pwd      string
		isExist  bool
		isLogin  bool
	)

	adminPwd := "alongpassword"
	adminPwdMd5Value := fmt.Sprintf("%X", md5.Sum([]byte(adminPwd)))

	for {
		if loginNum > 2 {
			isExist = true
		}

		if isExist {
			break
		}

		fmt.Println("请输入密码(输错3次将自动退出)：")
		fmt.Scan(&pwd)
		pwdMd5Value := fmt.Sprintf("%X", md5.Sum([]byte(pwd)))

		if adminPwdMd5Value == pwdMd5Value {
			isLogin = true
			isExist = true
		} else {
			loginNum++
			fmt.Printf("密码错误!!!,错误次数%d\n", loginNum)
		}

	}

	if isLogin {
		mainUserManager := utils.MyUserManager()
		// fmt.Printf("%T\n", mainUserManager) //输出*utils.UserManager
		// fmt.Println(mainUserManager)        //&{1    []}
		mainUserManager.MainMenu()
	}

}
