package funcs

import (
	"errors"
	"fmt"
)

// CmdToFuncMap mapping the cmd to corresponding func
var CmdToFuncMap = map[string]func(){}

// CmdToFunc mapping cmd to corresponding funcs
func CmdToFunc(cmd string, f func()) {
	if _, ok := CmdToFuncMap[cmd]; ok {
		panic(fmt.Sprintf("command %s already exists", cmd))
	}
	CmdToFuncMap[cmd] = f
}

// AddFunc register cmd to func
func AddFunc() {
	CmdToFunc("add", AddUser)
	CmdToFunc("del", DelUser)
}

// ExecFunc execute the func of a cmd
func ExecFunc(input string) error {
	if f, ok := CmdToFuncMap[input]; ok {
		f()
	} else {
		return errors.New("please input correct cmd")
	}
	return nil
}
