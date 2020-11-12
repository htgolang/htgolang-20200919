package funcs

import (
	"CMS/models"
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

	models.MUE.PersistentStorage.SyncFromDBToMemory()

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
	if index := models.MUE.IndexOfElement(func() (ret int) { ret, _ = strconv.Atoi(ID); return }()); index == -1 {
		fmt.Println("no User found...")
	} else {
		fmt.Printf("********\nFind User:\n********\n")
		models.Users[index].PrintElement()
		for {
			Choice, err := line.State.Prompt("Do you want to delete this User? (y/n) > ")
			if err != nil {
				return
			}
			if strings.ToLower(strings.TrimSpace(Choice)) == "y" || strings.ToLower(strings.TrimSpace(Choice)) == "yes" {
				models.MUE.RemoveElement(func() (ret int) { ret, _ = strconv.Atoi(ID); return }())
				err = models.MUE.PersistentStorage.DeleteFromDB(func() (ret int) { ret, _ = strconv.Atoi(ID); return }())
				if err != nil {
					fmt.Println("Remove User in DB Failed...")
					fmt.Printf("Error Message: %v\n", err)
					return
				}
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
