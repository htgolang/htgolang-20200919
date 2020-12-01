package main

import (
	"fmt"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day09-20201121/Go3028-Beijing-lisuo/user_manager/cmd/funcs"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day09-20201121/Go3028-Beijing-lisuo/user_manager/cmd/socket"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day09-20201121/Go3028-Beijing-lisuo/user_manager/define"
)

func main() {
	serv()
}

func serv() {
	//	var opt string

	// add some users and map cmd to funcs
	funcs.Init(&define.UserList)
	socket.Server()
	fmt.Println(define.UserList)

	// login
	//if !funcs.Login() {
	//	return
	//}

	//// login prompt
	//funcs.ShowHelp()
	//fmt.Print("> ")
	//// main loop for manager users
	//for {
	//	opt = utils.Read()
	//	// exec the corresponding func of the cmd
	//	err := funcs.ExecFunc(opt)
	//	fmt.Print("> ")
	//	if err != nil {
	//		fmt.Print(err)
	//	}
	//}
}
