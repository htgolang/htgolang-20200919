package function

import (
	"fmt"
	"strings"
)

const (
	deletePrompts   string = "Are you sure delete (y/n): "
	idDeletePrompts string = "enter you want delete user is id: "
)

//Del xxx
func Del() {
	usermap, usersIndex := getUserMessage(idDeletePrompts)
	if len(usersIndex) == 0 {
		fmt.Println("user not exist!!")
	} else if len(usersIndex) == 1 {
		printTable(usermap)
		yn := getQueryUserInput(deletePrompts)
		if strings.ToLower(yn) == "y" {
			delUsers(usersIndex[0])
		}
	} else {
		fmt.Println("enter user is primary id : ")
	}
}

func delUsers(index int) {
	if index == len(Users) {
		Users = Users[:len(Users)-1]
	} else if index == 0 {
		Users = Users[1:]
	} else {
		Users = append(Users[:index], Users[index+1:]...)
	}
}
