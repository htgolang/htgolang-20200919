package utils

import (
	"fmt"
)

func Menu() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("输入命令有误，请输入help查看所有命令")
			Menu()
		}
	} ()
	persist := GetPersist()
	Load()
	if len(UsersList) == 0 {
		fmt.Println("检测到当前没有用户信息，正在初始化admin账户")
		InitAdmin()
		persist.Save()
	}
	for {
		order := ""
		fmt.Printf("请输入要执行命令:")
		fmt.Scan(&order)
		FuncMap[order]()
		if order == "add" || order == "upd" || order == "del" {
			persist.Save()
		}
	}
}