package main

import (
	"usermanagement/controllers"
	"usermanagement/manager"
	"usermanagement/modules"
	_ "usermanagement/routes"
	"usermanagement/utils"
)

func help() {
	data := [][]string{
		{"add", "add user"},
		{"del", "del user"},
		{"modify", "modify user"},
		{"query", "query user"},
		{"exit", "quit"},
	}
	header := []string{"Command", "Features"}
	utils.OutputTable(data, header)
}


func main() {
	lib := modules.NewUserManager()
	// init
	lib.SetDefaultUser()

	if controllers.Login(lib) {
		help()
		manager.Run(lib)
	}

}
