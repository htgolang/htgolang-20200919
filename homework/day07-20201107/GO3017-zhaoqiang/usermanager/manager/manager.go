package manager

import (
	"fmt"
	"os"
	"zhao/controller"
)

var Routers = map[string]func(){}

func Run() {
	for {
		order := controller.CommandInput()
		if action, ok := Routers[order]; ok {
			action()
		} else if order == "exit" {
			os.Exit(0)
		} else {
			fmt.Printf("command not exit\n\n")
		}
	}
}
