package funcs

import (
	"GO3020-beijing-jiangchen/models"
	"fmt"
	"strconv"
	"strings"
)

//ModifyUser ...
// Modify attributes of User which belongs to User slice.
// Step 1: retrieve User from User slice using the given ID.
// If returns -1 which means there's no User can be found, then print sth and return.
// If the return value does not equal -1, which means we find the User,
// then modify this User using models.ModifyElement
func ModifyUser() {
	line := models.NewLiner()
	defer line.Close()

	var ID string
	for {
		id, err := line.State.Prompt("please input the ID of User you want to modify > ")
		if err != nil {
			return
		}
		if strings.ToLower(strings.TrimSpace(id)) == "" {
			continue
		} else if _, err := strconv.Atoi(strings.TrimSpace(id)); err != nil {
			fmt.Println("Invalid input format...ID must be integer, please try again.")
			continue
		} else {
			ID = strings.ToLower(strings.TrimSpace(id))
			break
		}
	}
	if index := models.IndexOfElement(&models.Users, ID); index == -1 {
		fmt.Println("no User found...")
	} else {
		fmt.Printf("********\nFind User:\n********\n")
		models.PrintElement(models.Users[index])
		for {
			Choice, err := line.State.Prompt("Do you want to modify this User? (y/n) > ")
			if err != nil {
				return
			}
			if strings.ToLower(strings.TrimSpace(Choice)) == "y" || strings.ToLower(strings.TrimSpace(Choice)) == "yes" {
				Name, err := line.State.Prompt("Please input new Name > ")
				if err != nil {
					return
				}
				if strings.TrimSpace(Name) == "" {
					fmt.Println("input Name is blank, so keep origin Name as default.")
					Name = models.Users[index]["Name"]
				}
				Contact, err := line.State.Prompt("Please input new Contact Number > ")
				if err != nil {
					return
				}
				if strings.TrimSpace(Contact) == "" {
					fmt.Println("input Contact is blank, so keep origin Contact as default.")
					Contact = models.Users[index]["Contact"]
				}
				Address, err := line.State.Prompt("Please input new Address > ")
				if err != nil {
					return
				}
				if strings.TrimSpace(Address) == "" {
					fmt.Println("input Address is blank, so keep origin Address as default.")
					Address = models.Users[index]["Address"]
				}
				models.ModifyElement(&models.Users, ID, strings.TrimSpace(Name), strings.TrimSpace(Contact), strings.TrimSpace(Address))
				break
			} else if strings.ToLower(strings.TrimSpace(Choice)) == "n" || strings.ToLower(strings.TrimSpace(Choice)) == "no" {
				fmt.Println("Give up modifying User, nothing happened.")
				break
			} else {
				fmt.Println("Invalid input...please try again.")
			}
		}
	}
}
