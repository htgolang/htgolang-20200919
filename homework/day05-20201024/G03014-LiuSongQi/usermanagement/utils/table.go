package utils

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func OutputTable(data [][]string, header []string) {
	t := tablewriter.NewWriter(os.Stdout)
	// data := [][]string{
	// 	{"add", "add user"},
	// 	{"del", "del user"},
	// 	{"modify", "modify user"},
	// 	{"query", "query user"},
	// 	{"exit", "quit"},
	// 	{"help", "help"},
	// }
	// t.SetHeader([]string{"Command", "Features"})
	t.SetHeader(header)
	t.SetAutoFormatHeaders(false)
	t.SetAutoWrapText(false)
	t.SetReflowDuringAutoWrap(false)
	t.AppendBulk(data)
	t.Render()
}
