package funcs

import (
	"GO3020-beijing-jiangchen/models"
	"crypto/md5"
	"errors"
	"fmt"
	"strings"
)

//Login ...
// Login function. the default password is admin.
// Get input from stdin and check whether the input equals the default password,
// force exit if type the wrong password in 3 consecutive times, or start the program normally if not so
func Login() (e error) {

	CMSBanner()
	line := models.NewLiner()
	defer line.Close()

	fmt.Printf("\nWelcome to CMS System.\nPlease input login password.\n")
	for {
		input, err := line.State.PasswordPrompt("password : ")
		if err != nil {
			e = err
			return
		}
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		} else {
			if md5.Sum([]byte(input)) == models.Password {
				fmt.Printf("\nlogin success.\n")
				break
			} else {
				fmt.Println("password incorrect...")
				models.LoginCount++
			}
			if models.LoginCount == 3 {
				fmt.Println("Bye :(")
				e = errors.New("login failure time reach 3")
				break
			}
		}
	}
	return
}
