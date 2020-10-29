package funcs

import (
	"fmt"
)

func Serv() {
	var opt string
	if !Login() {
		return
	}
	// add some users and map cmd to funcs
	Init()
	CmdToFuncs["help"]()
	for {
		fmt.Scanln(&opt)
		// exec the corresponding func of the cmd
		for cmd, exeFunc := range CmdToFuncs {
			if opt == cmd {
				fmt.Print("\n|" + cmd + "User|\n")
				exeFunc()
				opt = ""
			}
		}
		switch opt {
		case "":
			fmt.Print("[\"help\" for help]> ")
		case "q", "Q":
			return
		default:
			Default()
			opt = ""
			continue
		}
	}
}
