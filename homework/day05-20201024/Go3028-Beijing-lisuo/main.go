package main

import (
	"fmt"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/aux"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/define"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/funcs"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/userOp"
)

func main() {
	Serv()
	fmt.Println("Exit.")
}

func Serv() {
	// add some users and map cmd to funcs
	userOp.Init(&define.UserList)
	if !aux.Login() {
		return
	}
	for {
		funcs.AddUser(&define.UserList)
	}

}
