package funcs

import (
	"fmt"

	define "github.com/htgolang/htgolang-20200919/tree/master/homework/day03-20201008/Go3028-Beijing-lisuo/define"
	utils "github.com/htgolang/htgolang-20200919/tree/master/homework/day03-20201008/Go3028-Beijing-lisuo/utils"
)

// AddUser purely add a define.User type user to define.UserList
func AddUser(usersList *[]map[int64]define.User, user define.User) {
	// generate the id of current user
	idc := utils.GenId()
	userc := map[int64]define.User{idc: user}
	define.UserList = append(*usersList, userc)
	fmt.Printf("User %v added.\n", userc[idc].Name)
}

// ask input and add user
func AddCurrentUser() {
	var name, phone, address string
	fmt.Print("Please input the user's name: \n> ")
	fmt.Scan(&name)
	fmt.Print("Please input the user's phone: \n> ")
	fmt.Scan(&phone)
	// make sure the phone number contains only pure digits
	for utils.JustDigits(phone) == false {
		fmt.Print("Please input a legal phone number: \n> ")
		fmt.Scan(&phone)
		if utils.JustDigits(phone) == true {
			break
		}
	}
	fmt.Print("Please input the user's address: \n> ")
	fmt.Scan(&address)
	AddUser(&define.UserList, NewUser(name, phone, address))
}
