package funcs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/define"
)

// DelUser del a user
// one can use Id or Name to find a user and del it
// if input a Id, find the user use this Id
// if input a Name, find the user use this Name
// if no this user, prompt error
func DelUser() {
	var name string
	var input string
	utils.Message("Who you want to del(Id/Name)?\n> ")
	input = GetField("Name")
	if s, err := strconv.Atoi(strings.TrimSpace(input)); err == nil {
		id := int64(s)
		//fmt.Printf("idType: %T  idValue: %v\n", id, id)
		user, _ := utils.IdFindUser(&define.UserList, id)
		if user.Name == "" {
			fmt.Println("No such user.", user)
		} else {
			fmt.Printf("Find user %v --> %v\nAre you sure to delete %v?(y/n) ", user.Name, user, user.Name)
			input = utils.Read()
			if strings.ToLower(input) == "y" {
				fmt.Println("deleted...........")
				utils.IdDelUser(&define.UserList, id)
			} else if strings.ToLower(input) == "n" {
				fmt.Println("Nothing changes.")
			}
		}
	} else {
		name = strings.ToLower(strings.TrimSpace(input))
		//fmt.Printf("nameType: %T  nameValue: %v\n", name, name)
		user, _ := utils.NameFindUser(&define.UserList, name)
		if (user == define.User{}) {
			fmt.Println("No such user.", user)
		} else {
			fmt.Printf("Find user %v --> %v\nAre you sure to delete %v?(y/n) ", name, user, name)
			input = utils.Read()
			if strings.ToLower(input) == "y" {
				fmt.Println("deleted...........")
				utils.NameDelUser(&define.UserList, name)
			} else if strings.ToLower(input) == "n" {
				fmt.Println("Nothing changes.")
			}
		}
	}
}
