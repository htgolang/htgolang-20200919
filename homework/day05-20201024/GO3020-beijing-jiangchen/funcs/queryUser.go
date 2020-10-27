package funcs

import (
	"GO3020-beijing-jiangchen/models"
	"fmt"
	"strings"
)

//QueryUser ...
// Foreach User from User slice, if any of its fields contains given substring,
// save it to the new slice and finally returns it using models.QueryElement.
// Print this new slice using funcs.printUsers
func QueryUser() {
	line := models.NewLiner()
	defer line.Close()

	QueryString, err := line.State.Prompt("Please input Query String > ")
	if err != nil {
		return
	}
	QueryString = strings.ToLower(strings.TrimSpace(QueryString))
	ResultUsers := models.QueryElement(&models.Users, QueryString)
	if len(ResultUsers) == 0 {
		fmt.Printf("********\nNo Users contain the Query String.********\n")
	} else {
		fmt.Printf("********\nUsers that contain the Query String:\n")
		PrintUsers(&ResultUsers)
	}
}
