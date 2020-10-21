package main

import (
	"fmt"
	"os"
	"usermanager/user"
)

func main() {
	var userInputPwd string
	pwd := user.GenPasswd("123456")

	var flag bool = true
	count := 0
	for count < 3 {
		if flag {
			fmt.Printf("Please login: ")
		} else {
			fmt.Printf("The password is incorrect, please re-enter: ")
		}
		fmt.Scan(&userInputPwd)
		if user.GenPasswd(userInputPwd) == pwd {
			flag = true
			break
		}
		flag = false
		count++
	}
	if flag {
		fmt.Println("Login Successful!!")
		user.Help()
		user.Run()
	} else {
		os.Exit(0)
	}

}
