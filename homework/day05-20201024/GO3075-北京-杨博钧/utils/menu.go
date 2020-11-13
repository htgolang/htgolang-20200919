package utils

import "fmt"

func Menu() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("输入命令有误，请输入help查看所有命令")
			Menu()
		}
	} ()
	for {
		order := ""
		fmt.Printf("请输入要执行命令:")
		fmt.Scan(&order)
		FuncMap[order]()
	}
}