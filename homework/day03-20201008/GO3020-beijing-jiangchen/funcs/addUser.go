package funcs

import (
	"GO3020-beijing-jiangchen/models"
	"fmt"
	"strings"
)

//AddUsers ...
// Add a User from std, just a simple wrap of models.AddElement
func AddUsers() {
	line := models.NewLiner()
	defer line.Close()

	Name, err := line.State.Prompt("Please input Name > ")
	if err != nil {
		return
	}
	if strings.TrimSpace(Name) == "" {
		fmt.Println("input Name is blank, so take Name \"Donald Trump\" as default.")
		Name = "Donald Trump"
	}
	Contact, err := line.State.Prompt("Please input Contact Number > ")
	if err != nil {
		return
	}
	if strings.TrimSpace(Contact) == "" {
		fmt.Println("input Contact is blank, so take Contact \"+1 2024561111\" as default.")
		Contact = "+1 2024561111"
	}
	Address, err := line.State.Prompt("Please input Address > ")
	if err != nil {
		return
	}
	if strings.TrimSpace(Address) == "" {
		fmt.Println("input Address is blank, so take Address \"1600 Pennsylvania Avenue NW, Washington, DC 20500, United States\" as default.")
		Address = "1600 Pennsylvania Avenue NW, Washington, DC 20500, United States"
	}
	models.AddElement(&models.Users, models.GenerateElement(models.GetMaxIDPlusOne(&models.Users), strings.TrimSpace(Name), strings.TrimSpace(Contact), strings.TrimSpace(Address)))
	fmt.Println("Add Users Finish.")
}
