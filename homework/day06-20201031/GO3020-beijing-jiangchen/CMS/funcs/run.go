package funcs

import (
	"CMS/models"
	"fmt"
	"reflect"
	"strings"
)

//FuncMap ...
var FuncMap = map[string]interface{}{
	"add":     AddUsers,
	"show":    PrintAllUsers,
	"delete":  RemoveUser,
	"modify":  ModifyUser,
	"query":   QueryUser,
	"syncdb":  SyncToDB,
	"syncmem": SyncToMem,
	"help":    Help,
	"exit":    "exit",
	"":        0,
}

//Run ...
//Just a wrapper of some functions
func Run() {

	err := Login()
	if err != nil {
		return
	}

	Help()

	line := models.NewLiner()
	defer line.Close()

	for {
		input, err := line.State.Prompt("CMS > ")
		if err != nil {
			return
		}
		input = strings.TrimSpace(input)
		f, ok := FuncMap[input]
		if ok {
			if reflect.TypeOf(f).Name() == "string" {
				fmt.Println("Bye :)")
				break
			} else if reflect.TypeOf(f).Name() == "int" {
				continue
			} else {
				f.(func())()
			}
		} else {
			fmt.Println("input error...please try again.")
		}
	}
}
