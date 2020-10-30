package funcs

import (
	"fmt"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/define"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/utils"
)

// GetField prompt input the User's field
func GetField(f string) string {
	for _, field := range define.UserField {
		if field == f {
			fmt.Printf("Please input %v: ", f)
			input := utils.Read()
			return input
		}
	}
	return f
}
