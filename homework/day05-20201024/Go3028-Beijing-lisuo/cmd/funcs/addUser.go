package funcs

import (
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/define"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/funcs"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/utils"
)

// input times
const cnt int = 3

// AddUser ask input and add user
func AddUser() {
	ul := &define.UserList
	var uc define.User
	var Name string
	ID := funcs.GetMaxID(ul) + 1
	for i := 0; i < cnt; i++ {
		Name = GetField("Name")
		user, err := funcs.NameFindUser(ul, Name)
		// find the user name, so prompt reinput
		if err == nil {
			utils.Message("The person already exists: ")
			utils.Message(user.Name)
		} else if Name == "" {
			utils.Message("Must specify a name!")
		} else {
			break
		}
		if i == cnt-1 {
			utils.Message("Don't be silly, You'v got every opportunity!")
			return
		}
	}
	Cell := GetField("Cell")
	for i := 0; i < cnt; i++ {
		if !utils.JustDigits(Cell) {
			utils.Message("Please input a real cell number: ")
			Cell = GetField("Cell")
		} else {
			break
		}
		if i == cnt-1 {
			utils.Message("Don't be silly, You'v got every opportunity!")
			return
		}
	}
	Address := GetField("Address")
	Born := GetField("Born")
	for i := 0; i < cnt; i++ {
		if err := DateCheck(Born); err != nil {
			utils.Message("Please input a legal born time[YYYY.MM.DD]: ")
			Born = GetField("Born")
		} else {
			break
		}
		if i == cnt-1 {
			utils.Message("Don't be silly, You'v got every opportunity!")
			return
		}
	}
	Passwd := GetField("Passwd")
	uc = funcs.NewUser(ID, Name, Cell, Address, Born, Passwd)
	define.UserList = append(*ul, uc)
}
