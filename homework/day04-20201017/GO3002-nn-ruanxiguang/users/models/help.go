package models

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func Help() {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetAutoFormatHeaders(false)
	t.SetAutoWrapText(false)
	t.SetReflowDuringAutoWrap(false)
	// 表头
	t.SetHeader([]string{"CMD", "FUNCTION"})

	//内容
	t.Append([]string{"add", "Add a User"})
	t.Append([]string{"delete", "Delete a User"})
	t.Append([]string{"modify", "Modify a User"})
	t.Append([]string{"query", "Search User"})
	t.Append([]string{"help", "Show this help list"})
	t.Append([]string{"exit", "Exit"})
	t.Render()
}
