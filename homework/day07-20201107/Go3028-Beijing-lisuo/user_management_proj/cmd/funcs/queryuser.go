package funcs

import (
	"fmt"
	"strings"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day07-20201107/Go3028-Beijing-lisuo/user_management_proj/define"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day07-20201107/Go3028-Beijing-lisuo/user_management_proj/utils"
)

// QueryUser for search one or more user use define.User.Name
// or define.User.Address or define.User.Cell or define.User.Born
func QueryUser() {
	ul := &define.UserList
	var input string
	var gotUsers []define.User
	fmt.Print("Please input query string: \n> ")
	input = utils.Read()
	for _, user := range *ul {
		if ContainsIput(user, input) {
			gotUsers = append(gotUsers, user)
		}
	}
	ShowUserList(&gotUsers)
}

// ContainsIput check if a field of a define.User type
// user contains a string
func ContainsIput(u define.User, input string) bool {
	return strings.Contains(strings.ToLower(u.Name), input) ||
		strings.Contains(strings.ToLower(u.Address), input) ||
		strings.Contains(u.Cell, input) ||
		strings.Contains(u.Born.Format("2006.01.02"), input)
}
