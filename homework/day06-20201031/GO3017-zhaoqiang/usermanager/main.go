package main

import (
	"fmt"
	"os"
	"zhao/controller"
	"zhao/manager"
	_ "zhao/router"
)

func main() {
	var passwdCount int = 3
	if !controller.PasswdAuth(passwdCount) {
		fmt.Printf("passwd input count %d, processes exit\n", passwdCount)
		os.Exit(-1)
	}
	controller.View()

	manager.Run()

}
