package userutils

import (
"github.com/olekukonko/tablewriter"
"os"
)

func Showintable(header []string, data [][]string){
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
