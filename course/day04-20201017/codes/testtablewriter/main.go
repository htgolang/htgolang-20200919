package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetAutoFormatHeaders(false)
	t.SetAutoWrapText(false)
	t.SetReflowDuringAutoWrap(false)

	t.SetHeader([]string{"ID", "名字", "联系方式"})
	t.Append([]string{"1", "kk", "1xxxxxxxx"})
	t.Append([]string{"2", "kk2", "1xxxxxxxx"})
	t.Append([]string{"3", "kk3", "1xxxxxxxx"})

	t.Render()
}
