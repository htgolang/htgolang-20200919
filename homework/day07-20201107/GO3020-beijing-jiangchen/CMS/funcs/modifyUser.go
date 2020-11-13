package funcs

import (
	"CMS/miscs"
	"CMS/models"
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

	models.MUE.PersistentStorage.SyncFromDBToMemory()

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
	if index := models.MUE.IndexOfElement(func() (ret int) { ret, _ = strconv.Atoi(ID); return }()); index == -1 {
		fmt.Println("no User found...")
	} else {
		fmt.Printf("********\nFind User:\n********\n")
		models.Users[index].PrintElement()
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
					Name = models.Users[index].Name
				}
				Tel, err := line.State.Prompt("Please input new Tel Number > ")
				if err != nil {
					return
				}
				if strings.TrimSpace(Tel) == "" {
					fmt.Println("input Tel is blank, so keep origin Tel as default.")
					Tel = models.Users[index].Tel
				}
				Address, err := line.State.Prompt("Please input new Address > ")
				if err != nil {
					return
				}
				if strings.TrimSpace(Address) == "" {
					fmt.Println("input Address is blank, so keep origin Address as default.")
					Address = models.Users[index].Address
				}
				var Birthday string
				for i := 0; i < 6; i++ {
					Birthday, err = line.State.Prompt("Please input new Birthday(format: YYYY-MM-DD) > ")
					if err != nil {
						return
					}
					if err := miscs.DateCheck(strings.TrimSpace(Birthday)); err != nil {
						if i == 5 {
							fmt.Println("invalid input 5 times. so keep origin Birthday as default.")
							Birthday = models.Users[index].Birthday.Format("2006-01-02")
							break
						} else {
							fmt.Println("invalid input format, please input as format YYYY-MM-DD.")
							continue
						}
					} else {
						break
					}
				}
				var PasswordPlainExt string
				for i := 0; i < 6; i++ {
					PasswordPlain, err := line.State.PasswordPrompt("Please input new password > ")
					if err != nil {
						return
					}
					if strings.TrimSpace(PasswordPlain) == "" {
						if i == 5 {
							fmt.Println("password empty for 5 times, so keep origin Password as default.")
							break
						} else {
							fmt.Println("password is empty...please input again.")
							continue
						}
					} else {
						PasswordPlainExt = PasswordPlain
						break
					}
				}
				if models.MUE.CheckUserName(strings.TrimSpace(Name)) && Name != models.Users[index].Name {
					fmt.Println("name duplicate in CMS...Abort...")
					return
				}
				models.MUE.ModifyElement(func() (ret int) { ret, _ = strconv.Atoi(ID); return }(), strings.TrimSpace(Name), strings.TrimSpace(Tel), strings.TrimSpace(Address), strings.TrimSpace(Birthday), strings.TrimSpace(PasswordPlainExt))
				err = models.MUE.PersistentStorage.ModifyFromDB(models.GenerateElement(func() (ret int) { ret, _ = strconv.Atoi(ID); return }(), strings.TrimSpace(Name), strings.TrimSpace(Tel), strings.TrimSpace(Address), strings.TrimSpace(Birthday), strings.TrimSpace(PasswordPlainExt)))
				if err != nil {
					fmt.Println("Modify User to DB Failed...")
					fmt.Printf("Error Message: %v\n", err)
					return
				}
				fmt.Println("Modify User Finish.")
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
