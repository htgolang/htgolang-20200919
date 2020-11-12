package utils

import (
	"fmt"
	"os"
)

func Exit() {
	choise := ""
	for {
		fmt.Print("是否退出程序(y/n)?:")
		fmt.Scan(&choise)
		switch choise {
		case "y", "Y":
			os.Exit(0)
		case "n", "N":
			return
		default:
			fmt.Println("输入有误，请重新输入")
			continue
		}
	}
}