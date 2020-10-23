package authentication

import (
	"crypto/md5"
	"fmt"
	"log"

	"github.com/howeyc/gopass"
)

const (
	passWD string = "4e7d489b49ec93dbf53ce37aee778593" //123@qwe
)

//UserInput return a string from user input
func UserInput() (string, error) {
	fmt.Printf("输入密码：")
	UserInputPasswd, err := gopass.GetPasswd()
	if err != nil {
		return "", err
	}
	UserInPasswd := string(UserInputPasswd)
	return UserInPasswd, nil
}

// AuthUserPW  authentication the user enter the passwd correct
// count is set passwd miss counters
func AuthUserPW(count int) bool {
	for i := 0; i < count; i++ {
		passwd, err := UserInput()
		if err != nil {
			fmt.Println(err)
		}

		passwd = fmt.Sprintf("%x", md5.Sum([]byte(passwd)))
		if passwd == passWD {
			return true
		}
		log.Println("Permission denied, please try again.")
	}
	log.Fatalf("passwd input count %d, processes exit\n", count)
	return false
}
