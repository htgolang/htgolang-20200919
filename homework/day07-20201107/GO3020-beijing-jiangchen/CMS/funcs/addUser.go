package funcs

import (
	"CMS/miscs"
	"CMS/models"
	"fmt"
	"strings"
)

//AddUsers ...
// Add a User from std, just a simple wrap of models.AddElement
func AddUsers() {
	line := models.NewLiner()
	defer line.Close()

	models.MUE.PersistentStorage.SyncFromDBToMemory()

	Name, err := line.State.Prompt("Please input Name > ")
	if err != nil {
		return
	}
	if strings.TrimSpace(Name) == "" {
		fmt.Println("input Name is blank, so take Name \"Donald Trump\" as default.")
		Name = "Donald Trump"
	}
	Tel, err := line.State.Prompt("Please input Tel Number > ")
	if err != nil {
		return
	}
	if strings.TrimSpace(Tel) == "" {
		fmt.Println("input Tel is blank, so take Tel \"+1 2024561111\" as default.")
		Tel = "+1 2024561111"
	}
	Address, err := line.State.Prompt("Please input Address > ")
	if err != nil {
		return
	}
	if strings.TrimSpace(Address) == "" {
		fmt.Println("input Address is blank, so take Address \"1600 Pennsylvania Avenue NW, Washington, DC 20500, United States\" as default.")
		Address = "1600 Pennsylvania Avenue NW, Washington, DC 20500, United States"
	}
	var Birthday string
	for i := 0; i < 6; i++ {
		Birthday, err = line.State.Prompt("Please input Birthday(format: YYYY-MM-DD) > ")
		if err != nil {
			return
		}
		if err := miscs.DateCheck(strings.TrimSpace(Birthday)); err != nil {
			if i == 5 {
				fmt.Println("invalid input 5 times. so take Birthday \"1990-07-18\" as default.")
				Birthday = "1990-07-18"
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
		PasswordPlain, err := line.State.PasswordPrompt("Please input password > ")
		if err != nil {
			return
		}
		if strings.TrimSpace(PasswordPlain) == "" {
			if i == 5 {
				fmt.Println("password empty for 5 times, so take password \"admin\" as default.")
				PasswordPlain = "admin"
				PasswordPlainExt = PasswordPlain
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
	if models.MUE.CheckUserName(strings.TrimSpace(Name)) {
		fmt.Println("name duplicate in CMS...Abort...")
		return
	}
	element := models.GenerateElement(models.MUE.GetMaxIDPlusOne(), strings.TrimSpace(Name), strings.TrimSpace(Tel), strings.TrimSpace(Address), strings.TrimSpace(Birthday), strings.TrimSpace(PasswordPlainExt))
	models.MUE.AddElement(*element)
	err = models.MUE.PersistentStorage.InsertToDB(element)
	if err != nil {
		fmt.Println("Add User to DB Failed...")
		fmt.Printf("Error Message: %v\n", err)
		return
	}
	fmt.Println("Add User Finish.")
}
