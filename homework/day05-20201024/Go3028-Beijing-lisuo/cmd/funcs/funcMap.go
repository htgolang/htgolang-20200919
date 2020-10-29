package funcs

import (
	"errors"

	utils "github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/utils"
)

// inputCmd contains the cmd input by someone
// funcList contains the corresponding funcs of the cmd
// CmdToFuncs mapping the cmd to corresponding funcs
var inputCmd = []string{"add", "del", "mod", "query",
	"show", "help", "cls"}
var funcList = []func(){AddCurrentUser, DelUser, ModifyUser,
	QueryCurrentUser, ShowUserList, ShowHelp, utils.CallClear}
var CmdToFuncs = make(map[string]func())

// mapping cmd to corresponding funcs
func FuncMap() {
	// fmt.Println("inputCmd len: ", len(inputCmd))
	// fmt.Println("funcList len: ", len(funcList))
	if len(inputCmd) == len(funcList) {
		for i, v := range inputCmd {
			CmdToFuncs[v] = funcList[i]
			//CmdToFuncs[v]()
		}
	}
}

// execute the func of a cmd
func ExecFunc(input string) error {
	for cmd, _ := range CmdToFuncs {
		if cmd == input {
			CmdToFuncs[cmd]()
		} else {
			return errors.New("Please input correct cmd.")
		}
	}
	return nil
}
