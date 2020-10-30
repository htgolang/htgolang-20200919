package funcs

import (
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/define"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/userop"
)

// AddUser ask input and add user
func AddUser(ul *[]define.User) {
	var uc define.User
	ID := GetField("Id")
	Name := GetField("Name")
	Address := GetField("Address")
	Cell := GetField("Cell")
	Born := GetField("Born")
	Passwd := GetField("Passwd")
	userop.NewUser(ID, Name, Address, Cell, Born, Passed)
	define.UserList = append(*ul, uc)
}
