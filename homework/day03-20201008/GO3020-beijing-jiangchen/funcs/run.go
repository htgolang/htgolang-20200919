package funcs

import (
	"GO3020-beijing-jiangchen/models"
	"fmt"
	"strings"
)

//Run ...
//Just a wrapper of some functions
func Run() {
	line := models.NewLiner()
	defer line.Close()

	Motd()
	for {
		input, err := line.State.Prompt("command > ")
		if err != nil {
			return
		}
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		} else if input == "1" {
			PrintUsers(&models.Users)
		} else if input == "2" {
			AddUsers()
		} else if input == "3" {
			RemoveUser()
		} else if input == "4" {
			ModifyUser()
		} else if input == "5" {
			QueryUser()
		} else if input == "h" {
			Motd()
		} else if input == "q" {
			break
		} else {
			fmt.Println("input error...please try again.")
		}
	}
}
