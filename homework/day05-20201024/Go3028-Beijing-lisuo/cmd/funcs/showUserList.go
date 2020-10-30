package funcs

import (
	_ "fmt"
	"os"
	"strconv"
	_ "strconv"
	_ "text/tabwriter"

	"github.com/olekukonko/tablewriter"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/define"
)

func ShowUserList() {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetAutoFormatHeaders(false)
	t.SetAutoWrapText(false)
	t.SetReflowDuringAutoWrap(false)

	t.SetHeader([]string{"ID", "Name", "Address", "Cell", "Born", "Passwd"})
	for _, user := range define.UserList {
		id := strconv.FormatUint(uint64(user.Id), 10)
		t.Append([]string{id, user.Name, user.Cell, user.Address})
	}
	t.Render()
}
