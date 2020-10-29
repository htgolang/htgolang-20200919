package main

import (
	"fmt"
	"usermanagement/controllers"
	"usermanagement/manager"
	"usermanagement/modules"
	_ "usermanagement/routes"
	"usermanagement/utils"
)

func help() {
	data := [][]string{
		{"add", "add user"},
		{"del", "del user"},
		{"modify", "modify user"},
		{"query", "query user"},
		{"exit", "quit"},
	}
	header := []string{"Command", "Features"}
	utils.OutputTable(data, header)
}

func auth(passwd string) bool {
	for i := 0; i < 3; i++ {
		if utils.Md5Text(utils.Input("请输入密码: ")) == passwd {
			return true
		} else {
			fmt.Println("密码输入错误")
		}
	}
	return false
}

func main() {
	lib := modules.NewUserManager()
	// init
	lib.SetDefaultUser()

	if controllers.Login(lib) {
		help()
		manager.Run(lib)
	}

}
