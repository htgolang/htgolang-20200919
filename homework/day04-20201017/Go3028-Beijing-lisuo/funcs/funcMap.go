package funcs

import (
	"fmt"

	utils "github.com/htgolang/htgolang-20200919/tree/master/homework/day04-20201017/Go3028-Beijing-lisuo/utils"
)

var inputCmd = []string{"add", "del", "mod", "query",
	"show", "h", "cls", ""}
var funcList = []func(){AddCurrentUser, DelUser, ModifyUser,
	QueryCurrentUser, ShowUserList, ShowHelp, utils.CallClear, Default}
var directAndFuncs = make(map[string]func())

func FuncMap() {
	fmt.Println("inputCmd len: ", len(inputCmd))
	for i, v := range inputCmd {
		directAndFuncs[v] = funcList[i]
		//directAndFuncs[v]()
	}
}
