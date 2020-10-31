package funcs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/define"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/utils"
)

// DelUser del a user
// one can use ID or Name to find a user and del it
// if input a ID, find the user use this ID
// if input a Name, find the user use this Name
// if no this user, prompt error
func DelUser() {
	var name string
	var input string
	utils.Message("Who you want to del(ID/Name)?\n> ")
	input = utils.GetField("Name")
	if s, err := strconv.Atoi(strings.TrimSpace(input)); err == nil {
		id := int64(s)
		//fmt.Printf("idType: %T  idValue: %v\n", id, id)
		user, _ := IDFindUser(&define.UserList, id)
		if user.Name == "" {
			fmt.Println("No such user.", user)
		} else {
			ShowUser(id)
			fmt.Printf("Find user: %v\nAre you sure to delete %v?(y/n)\n> ", user.Name, user.Name)
			input = utils.Read()
			if strings.ToLower(input) == "y" {
				fmt.Println("deleted...........")
				IDDelUser(&define.UserList, id)
			} else if strings.ToLower(input) == "n" {
				fmt.Println("Nothing changes.")
			}
		}
	} else {
		name = strings.ToLower(strings.TrimSpace(input))
		//fmt.Printf("nameType: %T  nameValue: %v\n", name, name)
		user, _ := NameFindUser(&define.UserList, name)
		if (user == define.User{}) {
			fmt.Println("No such user.", user)
		} else {
			ShowUser(user.ID)
			fmt.Printf("Find user: %v\nAre you sure to delete %v?(y/n)\n> ", user.Name, user.Name)
			input = utils.Read()
			if strings.ToLower(input) == "y" {
				fmt.Println("deleted...........")
				NameDelUser(&define.UserList, name)
			} else if strings.ToLower(input) == "n" {
				fmt.Println("Nothing changes.")
			}
		}
	}
}
