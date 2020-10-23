package usertools

import (
	"crypto/md5"
	"fmt"
	"os"
)

func userLogin(authkey string, num int) bool {
	var pwd string
	if num == 0 {
		fmt.Println("invalid user!!")
		os.Exit(1)
	}
	fmt.Printf("please enter your password (accept), you have %d chance: ", num)
	fmt.Scan(&pwd)
	pwdmd5 := fmt.Sprintf("%x", md5.Sum([]byte(pwd)))

	if authkey == pwdmd5 {
		return true
	}
	return false

}
