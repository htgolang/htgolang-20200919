package main

import (
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day10-20201128/Go3028-Beijing-lisuo/user_manager/cmd/funcs"
	_ "github.com/htgolang/htgolang-20200919/tree/master/homework/day10-20201128/Go3028-Beijing-lisuo/user_manager/cmd/funcs"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day10-20201128/Go3028-Beijing-lisuo/user_manager/define"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day10-20201128/Go3028-Beijing-lisuo/user_manager/web"
)

func main() {
	serv()
	web.Serv()

}

func serv() {
	//	var opt string
	//
	//	// add some users and map cmd to funcs
	funcs.Init(&define.UserList)
	//
	//	// login
	//	if !funcs.Login() {
	//		return
	//	}
	//
	//	// login prompt
	//	funcs.ShowHelp()
	//	fmt.Print("> ")
	//	// main loop for manager users
	//	for {
	//		opt = utils.Read()
	//		// exec the corresponding func of the cmd
	//		err := funcs.ExecFunc(opt)
	//		fmt.Print("> ")
	//		if err != nil {
	//			fmt.Print(err)
	//		}
	//	}
}
