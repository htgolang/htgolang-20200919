package funcs

import (
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day10-20201128/Go3028-Beijing-lisuo/user_manager/define"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day10-20201128/Go3028-Beijing-lisuo/user_manager/utils"
)

// input times
const cnt int = 3

// AddUser ask input and add user
func AddUser() {
	ul := &define.UserList
	var uc define.User
	var Name string
	ID := GetMaxID(ul) + 1
	for i := 0; i < cnt; i++ {
		Name = utils.GetField("Name")
		user, err := NameFindUser(ul, Name)
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
	Cell := utils.GetField("Cell")
	for i := 0; i < cnt; i++ {
		if !utils.JustDigits(Cell) {
			utils.Message("Please input a real cell number: ")
			Cell = utils.GetField("Cell")
		} else {
			break
		}
		if i == cnt-1 {
			utils.Message("Don't be silly, You'v got every opportunity!")
			return
		}
	}
	Address := utils.GetField("Address")
	Born := utils.GetField("Born")
	for i := 0; i < cnt; i++ {
		if err := utils.DateCheck(Born); err != nil {
			utils.Message("Please input a legal born time[YYYY.MM.DD]: ")
			Born = utils.GetField("Born")
		} else {
			break
		}
		if i == cnt-1 {
			utils.Message("Don't be silly, You'v got every opportunity!")
			return
		}
	}
	Passwd := utils.GetField("Passwd")
	uc = NewUser(ID, Name, Cell, Address, Born, Passwd)
	define.UserList = append(*ul, uc)
}
