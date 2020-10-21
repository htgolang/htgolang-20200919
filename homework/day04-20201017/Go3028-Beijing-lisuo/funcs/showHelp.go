package funcs

import (
	"fmt"
)

func ShowHelp() {
	fmt.Print("  add:    addUser\n  del:    delUser\n  mod:    modifyUser\n" +
		"query:    queryUser\n show:    showUserList\n  cls:    clearConsole\n" +
		" help:    showHelp\n  Q|q:    Quit\n\n> ")
}

func Default() {
	fmt.Print("\n|Illegal input|\ntype \"h\" show help list.\n> ")
}
