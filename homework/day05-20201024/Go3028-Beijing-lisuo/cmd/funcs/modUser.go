package funcs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/define"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/utils"
)

// ModifyUser modify user based on ID or Name,
// if input a ID, find the user use this ID and prompt to input new value
// if input a Name, find the user use this Name and prompt to input new value
// if no this user, prompt error
func ModifyUser() {
	var input string
	var name string
	fmt.Print("Who you want to modify(ID/Name)?\n> ")
	input = utils.Read()
	if s, err := strconv.Atoi(strings.TrimSpace(input)); err == nil {
		id := int64(s)
		//fmt.Printf("idType: %T  idValue: %v\n", id, id)
		user, err := IDFindUser(&define.UserList, id)
		if err != nil {
			fmt.Println("No such user.", user)
		} else {
			ShowUser(id)
			fmt.Printf("Find user: %v\nAre you sure to modify %v?(y/n)\n> ", user.Name, user.Name)
			input = utils.Read()
			if strings.ToLower(input) == "y" {
				fmt.Println("modifying...........")
				user, err := IDModUser(&define.UserList, id)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("Modified user: %v\n", user.Name)
					ShowUser(id)
				}
			} else if strings.ToLower(input) == "n" {
				fmt.Println("Nothing changed.")
			}
		}
	} else {
		name = strings.ToLower(strings.TrimSpace(input))
		//fmt.Printf("nameType: %T  nameValue: %v\n", name, name)
		user, err := NameFindUser(&define.UserList, name)
		if err != nil {
			fmt.Println("No such user.", user)
		} else {
			ShowUser(user.ID)
			fmt.Printf("Find user: %v\nAre you sure to modify %v?(y/n)\n> ", user.Name, user.Name)
			input = utils.Read()
			if strings.ToLower(input) == "y" {
				fmt.Println("modifying...........")
				NameModUser(&define.UserList, name)
			} else if strings.ToLower(input) == "n" {
				fmt.Println("Nothing changed.")
			}
		}
	}
}
