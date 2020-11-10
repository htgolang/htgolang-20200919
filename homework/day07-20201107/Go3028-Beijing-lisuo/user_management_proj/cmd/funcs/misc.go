package funcs

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

// ShowHelp display help info
func ShowHelp() {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetAutoFormatHeaders(false)
	t.SetAutoWrapText(false)
	t.SetReflowDuringAutoWrap(false)

	//fmt.Println("|CMD help list|")
	t.SetHeader([]string{"CMD", "Function"})
	t.Append([]string{"add", "Add a User"})
	t.Append([]string{"del", "Delete a User"})
	t.Append([]string{"mod", "Modify a User"})
	t.Append([]string{"get", "Search User"})
	t.Append([]string{"show", "Show User List"})
	//t.Append([]string{"save", "Save User List to file"})
	t.Append([]string{"cls", "Clean the terminal"})
	t.Append([]string{"help", "Show this help list"})
	t.Append([]string{"mycmd", "Binding a new cmd to a exist func"})
	t.Append([]string{"Q|q|quit|exit", "Exit"})
	t.Render()
}

// DefaultTip show default prompt in main
func DefaultTip() {
	fmt.Print("\n|type \"h\" show help list|\n> ")
}
