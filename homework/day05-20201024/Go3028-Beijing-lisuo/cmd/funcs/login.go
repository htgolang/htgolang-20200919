package funcs

import (
	"crypto/md5"
	"fmt"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/define"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/utils"
)

// Login implement login func
func Login() bool {
	var logCount int8 = 0
	var logged bool = true
	for {
		fmt.Print("Input the UserName[admin]: \n> ")
		input := utils.Read()
		user, err := NameFindUser(&define.UserList, input)
		if err != nil {
			fmt.Println("No such user.")
			logCount++
			if logCount == 3 {
				logged = false
				break
			}
		} else {
			fmt.Print("Input the PassWord[qwert]: \n> ")
			input := utils.Read()
			inputPasswd := md5.Sum([]byte(input))
			if user.Passwd == inputPasswd {
				fmt.Println("Logged in.")
				return logged
			}
			fmt.Println("Wrong password...")
			logCount++
			if logCount == 3 {
				logged = false
				break
			}
			continue
		}
	}
	return logged
}
