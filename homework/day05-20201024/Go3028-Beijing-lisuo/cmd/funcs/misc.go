package funcs

import (
	"fmt"

	"time"
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

// DateCheck make sure the input date is formatted
func DateCheck(d string) error {
	_, err := time.Parse("2006.01.02", d)
	return err
}

// Message print debug info
func Message(v string) {
	fmt.Println(v)
}
