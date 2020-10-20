package funcs

import (
	"fmt"
)

//Motd ...
//Just print Motd
func Motd() {
	fmt.Println("============================")
	fmt.Println("=========== CMS  ===========")
	fmt.Println("==                        ==")
	fmt.Println("==== Press 1 and enter: ====")
	fmt.Println("==Show current Users.     ==")
	fmt.Println("==                        ==")
	fmt.Println("==== Press 2 and enter: ====")
	fmt.Println("==Add new User.           ==")
	fmt.Println("==                        ==")
	fmt.Println("==== Press 3 and enter: ====")
	fmt.Println("==Delete User.            ==")
	fmt.Println("==                        ==")
	fmt.Println("==== Press 4 and enter: ====")
	fmt.Println("==Modify User.            ==")
	fmt.Println("==                        ==")
	fmt.Println("==== Press 5 and ender: ====")
	fmt.Println("==Query User.             ==")
	fmt.Println("==                        ==")
	fmt.Println("==== Type h to show motd  ==")
	fmt.Println("==== Type q to quit       ==")
}
