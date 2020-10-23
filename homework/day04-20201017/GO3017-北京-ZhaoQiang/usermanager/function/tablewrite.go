package function

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func printTable(u []map[string]string) {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetAutoFormatHeaders(false)
	t.SetAutoWrapText(false)
	t.SetReflowDuringAutoWrap(false)
	t.SetHeader([]string{"id", "name", "tel", "address"})

	for _, userMap := range u {
		t.Append([]string{userMap["id"], userMap["name"], userMap["tel"], userMap["addr"]})
	}

	t.Render()
}
