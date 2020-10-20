package funcs

import (
	"GO3020-beijing-jiangchen/models"
	"os"

	"github.com/olekukonko/tablewriter"
)

//PrintUsers ...
// Print specified Users slices in an elegant way
func PrintUsers(Users *[]map[string]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetHeader([]string{"ID", "Name", "Contact", "Address"})
	for _, value := range *Users {
		table.Append(models.ConvertElementToSlice(value))
	}
	table.SetRowSeparator("-")
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.Render()
}

//PrintAllUsers ...
// Print All Users from Users slices
func PrintAllUsers() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetHeader([]string{"ID", "Name", "Contact", "Address"})
	for _, value := range models.Users {
		table.Append(models.ConvertElementToSlice(value))
	}
	table.SetRowSeparator("-")
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.Render()
}
