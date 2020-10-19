package funcs

import (
	"fmt"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day04-20201017/Go3028-Beijing-lisuo/utils"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day04-20201017/Go3028-Beijing-lisuo/define"
)

func Serv() {
	var opt string
	help := func() {
		fmt.Print("1. addUser\n2. delUser\n3. modifyUser\n" +
			"4. queryUser\n5. showUserList\nc. clearConsole\nh. showHelp\nq. Quit\n\n> ")
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
			SearchUser(&define.UserList)
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
		case "c":
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
