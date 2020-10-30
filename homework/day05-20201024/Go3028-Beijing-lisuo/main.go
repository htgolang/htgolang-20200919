package main

import (
	"fmt"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/aux"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/define"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/funcs"
)

func main() {
	Serv()
	fmt.Println("Exit.")
}

func Serv() {
	var opt string

	// add some users and map cmd to funcs
	userop.Init(&define.UserList)

	if !aux.Login() {
		return
	}

	for {
		fmt.Scanln(&opt)
		// exec the corresponding func of the cmd
		funcs.ExecFunc(opt)
		switch opt {
		case "":
			fmt.Print("[\"help\" for help]> ")
		case "q", "Q":
			return
		default:
			defaultTip()
			opt = ""
			continue
		}
	}
}

func defaultTip() {
	fmt.Print("\n|Illegal input|\ntype \"h\" show help list.\n> ")
}
