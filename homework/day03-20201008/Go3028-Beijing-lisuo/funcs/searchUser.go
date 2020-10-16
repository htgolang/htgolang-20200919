package funcs

import (
	"fmt"
	"strings"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day03-20201008/Go3028-Beijing-lisuo/utils"

	define "github.com/htgolang/htgolang-20200919/tree/master/homework/day03-20201008/Go3028-Beijing-lisuo/define"
)

// if define.User.Name or define.User.Address or define.User.Phone
// contains the input string, then Show those matched users
func SearchUser(user *[]map[int64]define.User) {
	var input string
	fmt.Print("Please input query string: \n> ")
	input = utils.Read()
	for _, userMap := range *user {
		for k, v := range userMap {
			if strings.Contains(strings.ToLower(v.Name), input) ||
				strings.Contains(strings.ToLower(v.Address), input) ||
				strings.Contains(v.Phone, input) {
				utils.ShowUser(k)
			}
		}

	}
}
