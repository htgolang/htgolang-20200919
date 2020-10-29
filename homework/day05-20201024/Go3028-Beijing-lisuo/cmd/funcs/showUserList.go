package funcs

import (
	_ "fmt"
	"os"
	"strconv"
	_ "strconv"
	_ "text/tabwriter"

	"github.com/olekukonko/tablewriter"

	define "github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/define"
)

// show user list when init (text/tabwriter)
//func ShowUserList() {
//	fmt.Println("|...Users list...|")
//	fmt.Println("|...Id...|...Name...|...Phone...|...Address...|")
//	for _, user := range define.UserList {
//		for k, v := range user {
//			w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0|tabwriter.Debug)
//			s := strconv.FormatInt(k, 10)
//			fmt.Fprintln(w, "|"+s+"\t"+v.Name+"\t"+v.Phone+"\t"+v.Address+" |")
//			w.Flush()
//		}
//	}
//	fmt.Println("")
//}

func ShowUserList() {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetAutoFormatHeaders(false)
	t.SetAutoWrapText(false)
	t.SetReflowDuringAutoWrap(false)

	t.SetHeader([]string{"ID", "Name", "Phone", "Location"})
	for _, user := range define.UserList {
		for k, v := range user {
			s := strconv.FormatInt(k, 10)
			t.Append([]string{s, v.Name, v.Phone, v.Address})
		}
	}
	t.Render()
}
