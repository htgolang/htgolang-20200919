package main

import (
	"fmt"
	"usermanage/utils"
)

func main() {
	fmt.Println("欢迎使用用户管理系统")
	if len(utils.UsersList) == 0 {
		fmt.Println("检测到当前没有用户信息，正在初始化admin账户")
		utils.InitAdmin()
	}
	utils.Menu()
}