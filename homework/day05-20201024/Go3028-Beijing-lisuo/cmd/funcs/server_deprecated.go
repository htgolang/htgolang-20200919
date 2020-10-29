package funcs

import (
	"fmt"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/utils"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/define"
)

func Serv() {
	var opt string
	help := func() {
		fmt.Print("  add:    addUser\n  del:    delUser\n  mod:    modifyUser\n" +
			"query:    queryUser\n show:    showUserList\n  cls:    clearConsole\n" +
			"    h:    showHelp\n  Q|q:    Quit\n\n> ")
	}
	if !Login() {
		return
	}
	Init()
	help()
	for {
		fmt.Scanln(&opt)
		switch opt {
		case "add", "a":
			fmt.Print("\n|addUser|\n")
			AddCurrentUser()
			opt = ""
			continue
		case "del", "d":
			fmt.Printf("\n|delUser|\n")
			DelUser()
			opt = ""
			continue
		case "mod", "m":
			fmt.Printf("\n|modUser|\n")
			ModifyUser()
			opt = ""
			continue
		case "query":
			fmt.Printf("\n|queryUser|\n")
			SearchUser(&define.UserList)
			opt = ""
			continue
		case "show", "s":
			fmt.Printf("\n|showUserList|\n")
			ShowUserList()
			opt = ""
			continue
		case "h":
			fmt.Printf("\n|showHelp|\n")
			help()
			opt = ""
			continue
		case "cls", "c":
			fmt.Printf("\n|clearConsole|\n")
			utils.CallClear()
			fmt.Print("[\"h\" for help]> ")
			opt = ""
			continue
		case "":
			fmt.Print("[\"h\" for help]> ")
		case "q", "Q":
			return
		default:
			fmt.Print("\n|Illegal input|\ntype \"h\" show help list.\n> ")
			opt = ""
			continue
		}
	}
}
