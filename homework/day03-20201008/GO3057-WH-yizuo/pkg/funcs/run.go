package funcs

import (
	"GO3057-WH-yizuo/pkg/models"
	"fmt"
)

//Run  ...
//运行逻辑
func Run() {
	View()
	var Input, decide string
	for decide != "exit" {
		fmt.Println("请输入您要执行的操作：")
		fmt.Scan(&Input)
		switch Input {
		case "a", "add", "append":
			UserAdd()
		case "d", "del", "delete":
			UserDelete()
		case "m", "mod", "modify":
			UserModify()
		case "l", "list":
			UsersList(&models.Users)
		case "q", "query":
			UserQuery()
		case "h", "help":
			View()
		case "exit", "quit":
			fmt.Println("Bay~")
			decide = "exit"
			break
		default:
			fmt.Println("请检查您输入的内容，可输入'h'查看帮助。")
		}
	}
}
