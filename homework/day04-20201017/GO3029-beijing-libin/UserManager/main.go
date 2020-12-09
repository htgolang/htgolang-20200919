package main

import (
	"UserManager/handle"
	"UserManager/model"
	"fmt"
	"os"
)

var (
	billboard string = `
	---------- USER MANAGEMENT ----------
1. LOGIN
2. REGISTER
3. EXIT
`
	loginPage    string = "---------- LOGIN PAGE    ----------"
	registerPage string = "---------- REGISTER PAGE ---------- "
	userPage     string = `
	---------- USER PAGE ----------
	1. SHOW ACTIVE USERS
	2. EDIT USER
	3. BLOCK USER
	4. EXIT
	`
)

func main() {
	// var inputUser string
	var youChoice int

	//这里会根据文件去取是否存在userSlice
	userslice := model.NewUserSlice()
	// var userinfo *model.UserInfos

	for {
		fmt.Println(billboard)
		fmt.Println("make your choice: ")
		fmt.Scanln(&youChoice)
		switch youChoice {
		case 1:
			//走登陆逻辑
			showLogin(userslice)
		case 2:
			//表示用户不存在，走注册逻辑
			showRegister(userslice)
		case 3:
			fmt.Println("byebye")
			os.Exit(0)
		default:
			fmt.Println("input number not valid")
			break
		}

	}

	//在main函数结束前，将userSlice的结果保存到本地的文件中

}

func showRegister(userSlice *model.UserSlice) {
	var username string
	var passwd string
	var action int
	var userinfo *model.UserInfo
	for {
		fmt.Println(registerPage)
		fmt.Println("USERNAME: ")
		fmt.Scanln(&username)
		fmt.Println("PASSWORD: ")
		fmt.Scanln(&passwd)
		fmt.Println("Press 1 to accept and exit, Press not 1 to loop until you accept")
		fmt.Scanln(&action)
		if action == 1 {
			handle.Register(username, passwd, userinfo, userSlice)
			fmt.Println("thanks!")
			break
		}
		fmt.Println("--------------------------------")
	}
}

func showLogin(userslice *model.UserSlice) {
	var loginuser string
	var loginpasswd string
	// var decide int
	for {
		fmt.Println(loginPage)
		fmt.Println("USERNAME: ")
		fmt.Scanln(&loginuser)
		userExist := userslice.FindByNameExist(loginuser)
		if userExist {
			//这里传递
			userinfo := userslice.FindUserByName(loginuser)
			fmt.Println("PASSWORD: ")
			fmt.Scanln(&loginpasswd)
			hashpwd := model.CheckMd5(loginpasswd)
			if hashpwd == userinfo.HashPwd {
				//表示MD5结果相同
				fmt.Printf("WELCOME BACK，%s\n", userinfo.Name)

				//这里进入
				showUser(userslice)
			} else {
				fmt.Println("PASSWORD ERROR..")
			}
		} else {
			//说明用户不存在
			fmt.Printf("%s not exist, plz input valid user\n", loginuser)
			break
		}

	}
}

func showUser(userslice *model.UserSlice) {
	var choice int
	for {
		fmt.Println(userPage)
		fmt.Println("make your choice ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			//显示可用的用户列表
			users := userslice.ShowActiveUser()
			for _, v := range users {
				fmt.Println(v)
			}
		case 2:
			editname := ""
			editpass := ""
			fmt.Println("plz input username to edit")
			fmt.Scanln(&editname)
			//这里判断要编辑的用户是否存在
			ok := userslice.FindByNameExist(editname)
			if ok {
				user := userslice.ReturnUser(editname)
				fmt.Printf("current username[%s] plz input a new one\n", user.Name)
				fmt.Scanln(&editname)
				fmt.Scanln(&editpass)

				user.EditUser(editname, editpass)
				fmt.Println("edit successfully")
			}
		case 3:
			//block user
			userPoint := ""
			fmt.Println("要屏蔽的用户是 ")
			fmt.Scanln(&userPoint)
			userslice.RemoveUser(userPoint)
		case 4:
			fmt.Println("byebye")
			os.Exit(0)
		default:
			fmt.Println("Invalid input, try again")

		}
	}
}
