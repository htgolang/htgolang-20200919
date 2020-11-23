package funcs

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day09-20201121/Go3028-Beijing-lisuo/user_manager/define"
)

// ShowUser show a user based on ID
func ShowUser(id int64) {
	t := tablewriter.NewWriter(os.Stdout)
	for _, user := range define.UserList {
		if user.ID == id {
			s := strconv.FormatInt(id, 10)
			t.Append([]string{s, user.Name, user.Cell, user.Address, user.Born.Format("2006.01.02"), user.Passwd})
			//md5.Sum([]byte(user.Passwd))})
		}
	}
	t.Render()
}

// ShowUserList display the UserList's contents
func ShowUserList(ul *[]define.User) {
	//ul := &define.UserList
	t := tablewriter.NewWriter(os.Stdout)
	t.SetAutoFormatHeaders(false)
	t.SetAutoWrapText(false)
	t.SetReflowDuringAutoWrap(false)
	t.SetHeader([]string{"ID", "Name", "Cell", "Address", "Born", "Passwd"})
	for _, user := range *ul {
		id := strconv.FormatUint(uint64(user.ID), 10)
		t.Append([]string{id, user.Name, user.Cell, user.Address,
			user.Born.Format("2006.01.02"), user.Passwd})
	}
	t.Render()
}

// ShowCurrentUserList make takes no arg
func ShowCurrentUserList() {
	ul := &define.UserList
	ShowUserList(ul)
}
