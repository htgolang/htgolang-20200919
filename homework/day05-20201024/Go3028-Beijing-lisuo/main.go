package main

import (
	"fmt"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/define"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/funcs"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/utils"
)

func main() {
	serv()
}

func serv() {
	var opt string

	// add some users and map cmd to funcs
	funcs.Init(&define.UserList)

	// login
	if !funcs.Login() {
		return
	}

	// login prompt
	funcs.ShowHelp()
	fmt.Print("> ")
	// main loop for manager users
	for {
		opt = utils.Read()
		// exec the corresponding func of the cmd
		err := funcs.ExecFunc(opt)
		fmt.Print("> ")
		if err != nil {
			fmt.Print(err)
		}
	}
}
