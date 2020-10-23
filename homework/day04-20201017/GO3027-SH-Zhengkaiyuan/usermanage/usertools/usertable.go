package usertools

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func userTable(userinfo []string) {

	t := tablewriter.NewWriter(os.Stdout)
	t.SetAutoFormatHeaders(false)
	t.SetAutoWrapText(false)
	t.SetReflowDuringAutoWrap(false)

	t.SetHeader([]string{"ID", "名字", "联系方式", "联系地址"})
	t.Append(userinfo)

	t.Render()
}
