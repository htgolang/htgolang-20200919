package funcs

import (
	"fmt"
	"homework/pkg/models"
)

func Run() {
	Bool := md5Check()
	if Bool == false{
		return
	}
	Show()
	var userChoice,cmd string
	for cmd != "exit" {
		fmt.Println("请选择你要执行的操作：")
		fmt.Scan(&userChoice)
		switch userChoice {
		case "a", "add", "append":
			userAdd()
		case "d", "del", "delete":
			userDel()
		case "m", "mod", "modify":
			userModify()
		case "l", "list":
			usersList(&models.Users)
		case "q", "query":
			userQuery()
		case "h", "help":
			Show()
		case "exit", "quit":
			fmt.Println("退出")
			cmd = "exit"
			break
		default:
			fmt.Println("请检查您输入的内容，可输入'h'查看帮助。")	
		}
	}
}