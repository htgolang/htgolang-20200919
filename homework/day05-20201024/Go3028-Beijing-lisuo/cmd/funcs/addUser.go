package funcs

import (
	"fmt"

	define "github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/define"
	utils "github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/utils"
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
	name = utils.Read()
	if name == "" {
		fmt.Println("The man/women must have a name.")
		return
	}
	fmt.Print("Please input the user's phone: \n> ")
	phone = utils.Read()
	// make sure the phone number contains only pure digits
	for utils.JustDigits(phone) == false {
		fmt.Print("Please input a legal phone number: \n> ")
		fmt.Scan(&phone)
		if utils.JustDigits(phone) == true {
			break
		}
		if phone == "" {
			phone = "999"
		}
	}
	fmt.Print("Please input the user's address: \n> ")
	address = utils.Read()
	if address == "" {
		address = "Beijing"
		fmt.Printf("The default address will be %v.\n", phone)
	}
	AddUser(&define.UserList, NewUser(name, phone, address))
}
