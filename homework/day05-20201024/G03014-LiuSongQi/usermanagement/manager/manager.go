package manager

import (
	"fmt"
	"usermanagement/modules"
	"usermanagement/utils"
)

// // register route
var routers = map[string]func(lib *modules.UserManager){}

func Register(op string, callback func(lib *modules.UserManager)) {
	if _, ok := routers[op]; ok {
		panic(fmt.Sprintf("command %s already exists", op))
	}
	routers[op] = callback
}

func Run(lib *modules.UserManager) {
	for {
		text := utils.Input("Please input command: ")
		if text == "exit" || text == "quit" {
			fmt.Println("Bye!")
			break
		}

		if action, ok := routers[text]; ok {
			action(lib)
		} else {
			fmt.Println("command error")
		}

	}
}
