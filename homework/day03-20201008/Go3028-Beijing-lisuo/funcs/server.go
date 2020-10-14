package funcs

import (
	"fmt"
)

func Serv() {
	var opt string
	help := func() {
		fmt.Print("1. addUser\n2. delUser\n3. modifyUser\n" +
			"4. queryUser\n5. showUserList\nh. showHelp\nq. Quit\n\n")
	}
	help()
	for {
		fmt.Scanln(&opt)
		switch opt {
		case "1":
			fmt.Println("addUser")
			AddCurrentUser()
			opt = ""
			continue
		case "2":
			fmt.Println("delUser")
			DelUser()
			opt = ""
			continue
		case "3":
			fmt.Println("modUser")
			opt = ""
			continue
		case "4":
			fmt.Println("queryUser")
			opt = ""
			continue
		case "5":
			fmt.Println("showUserList")
			ShowUserList()
			opt = ""
			continue
		case "h":
			fmt.Println("showHelp")
			help()
			opt = ""
			continue
		case "":
		case "q", "Q":
			return
		default:
			fmt.Println("Illegal input")
			opt = ""
			continue
		}
	}
}
