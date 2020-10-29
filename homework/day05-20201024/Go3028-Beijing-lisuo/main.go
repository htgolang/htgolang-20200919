package main

import (
	"fmt"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/define"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/userOp"
)

func main() {
	//funcs.Serv()
	//funcs.ExecFunc()
	//funcs.Serv()
	define.Init()
	userOp.NameFindUser(&define.UserList, "jack")
	fmt.Println("Exit.")
}
