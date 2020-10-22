package funcs

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

//func ShowHelp() {
//	fmt.Print("  add:    addUser\n  del:    delUser\n  mod:    modifyUser\n" +
//		"query:    queryUser\n show:    showUserList\n  cls:    clearConsole\n" +
//		" help:    showHelp\n  Q|q:    Quit\n\n> ")
//}

func Default() {
	fmt.Print("\n|Illegal input|\ntype \"h\" show help list.\n> ")
}

func ShowHelp() {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetAutoFormatHeaders(false)
	t.SetAutoWrapText(false)
	t.SetReflowDuringAutoWrap(false)

	//fmt.Println("|CMD help list|")
	t.SetHeader([]string{"CMD", "Function"})
	t.Append([]string{"add", "Add a User"})
	t.Append([]string{"del", "Delete a User"})
	t.Append([]string{"query", "Search User"})
	t.Append([]string{"show", "Show User List"})
	t.Append([]string{"cls", "Clean the terminal"})
	t.Append([]string{"help", "Show this help list"})
	t.Append([]string{"Q|q", "Exit"})
	t.Render()
}
