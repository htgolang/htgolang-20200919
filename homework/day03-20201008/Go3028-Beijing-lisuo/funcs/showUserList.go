package funcs

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	define "github.com/htgolang/htgolang-20200919/tree/master/homework/day03-20201008/Go3028-Beijing-lisuo/define"
)

func ShowUserList() {
	fmt.Println("|...Users list...|")
	fmt.Println("|...Id...|...Name...|...Phone...|...Address...|")
	for _, user := range define.UserList {
		for k, v := range user {
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0|tabwriter.Debug)
			s := strconv.FormatInt(k, 10)
			fmt.Fprintln(w, "|"+s+"\t"+v.Name+"\t"+v.Phone+"\t"+v.Address+" |")
			w.Flush()
		}
	}
	fmt.Println("")
}
