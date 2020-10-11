package funcs

import (
	"GO3020-beijing-jiangchen/models"
	"fmt"
	"strconv"
	"strings"
)

//RemoveUser ...
// Remove User whose ID equals the given value from User slice.
// Step1: retrieve User from User slice using the given ID.
// If returns -1 which means there's no User can be found, then print sth and return.
// If the return value does not equal -1, which means we find the User,
// Then delete the User using models.RemoveElement
func RemoveUser() {
	line := models.NewLiner()
	defer line.Close()

	var ID string
	for {
		id, err := line.State.Prompt("please input the ID of User you want to delete > ")
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
			Choice, err := line.State.Prompt("Do you want to delete this User? (y/n) > ")
			if err != nil {
				return
			}
			if strings.ToLower(strings.TrimSpace(Choice)) == "y" || strings.ToLower(strings.TrimSpace(Choice)) == "yes" {
				models.RemoveElement(&models.Users, ID)
				fmt.Println("Remove User success.")
				break
			} else if strings.ToLower(strings.TrimSpace(Choice)) == "n" || strings.ToLower(strings.TrimSpace(Choice)) == "no" {
				fmt.Println("Give up removing User, nothing happened.")
				break
			} else {
				fmt.Println("Invalid input...please try again.")
			}
		}
	}
}
