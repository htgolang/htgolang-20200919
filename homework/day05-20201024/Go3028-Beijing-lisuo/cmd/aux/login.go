package aux

import (
	"crypto/md5"
	"fmt"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/define"
)

var p string = "   "
var password [16]byte = md5.Sum([]byte(p))

func Login() bool {
	var logCount int8 = 0
	var logged bool = true
	for {
		fmt.Print("Input the password: \n> ")
		input := utils.Read()
		inputPasswd := md5.Sum([]byte(input))
		if define.User.passwd == inputPasswd {
			fmt.Println("Logged in.")
			return logged
		} else {
			fmt.Println("Wrong password...")
			logCount++
			if logCount == 5 {
				logged = false
				break
			}
			continue
		}
	}
	return logged
}
