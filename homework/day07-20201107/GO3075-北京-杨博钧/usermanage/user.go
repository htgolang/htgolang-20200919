package main

import (
	"fmt"
	"usermanage/utils"
)

func main() {
	utils.Flag()
	fmt.Println("欢迎使用用户管理系统")
	utils.Load()
	if len(utils.UsersList) > 0 {
		utils.Login()
	}
	utils.Menu()
}