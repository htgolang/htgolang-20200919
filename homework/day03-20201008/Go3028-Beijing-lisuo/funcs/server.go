package funcs

import (
	"fmt"
)

func Serv() {
	var opt string
	help := func() {
		fmt.Print("1. addUser\n2. delUser\n3. modifyUser\n" +
			"4. queryUser\n5. showUserList\nh. showHelp\nq. Quit\n\n> ")
	}
	help()
	for {
		fmt.Scanln(&opt)
		switch opt {
		case "1":
			fmt.Print("\n|addUser|\n")
			AddCurrentUser()
			opt = ""
			continue
		case "2":
			fmt.Printf("\n|delUser|\n")
			DelUser()
			opt = ""
			continue
		case "3":
			fmt.Printf("\n|modUser|\n")
			ModifyUser()
			opt = ""
			continue
		case "4":
			fmt.Printf("\n|queryUser|\n")
			opt = ""
			continue
		case "5":
			fmt.Printf("\n|showUserList|\n")
			ShowUserList()
			opt = ""
			continue
		case "h":
			fmt.Printf("\n|showHelp|\n")
			help()
			opt = ""
			continue
		case "":
			fmt.Print("> ")
		case "q", "Q":
			return
		default:
			fmt.Print("\n|Illegal input|\ntype \"h\" show help list.\n> ")
			opt = ""
			continue
		}
	}
}
