package main

import (
	"fmt"
	"usermanage/utils"
)

func main() {
	fmt.Println("欢迎使用用户管理系统")
	persist := utils.InitPersist()
	persist.Load()
	if len(utils.UsersList) > 0 {
		utils.Login()
	}
	utils.Menu()
}