package manager

import (
	"fmt"
	"os"
	"zhao/utils"
)

//Routers 程序路由
var Routers = map[string]func(){}

// Run 查询routes 执行相应命令
func Run(useraccount string) {
	order := utils.GetCommandLineInput(useraccount)
	if action, ok := Routers[order]; ok {
		action()
	} else if order == "exit" {
		os.Exit(0)
	} else {
		fmt.Printf("command not exist\n\n")
	}

}
