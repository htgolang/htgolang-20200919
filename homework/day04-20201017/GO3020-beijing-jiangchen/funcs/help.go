package funcs

import (
	"fmt"
)

//Help ...
//Just print Motd
func Help() {
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("show:     show current users.")
	fmt.Println("add:      add new user.")
	fmt.Println("delete:   delete user.")
	fmt.Println("modify:   modify user.")
	fmt.Println("query:    query user.")
	fmt.Println("help:     print help messages.")
	fmt.Println("exit:     exit.")
	fmt.Println("")
}
