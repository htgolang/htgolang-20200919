package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var nums int = rand.Intn(101)
	fmt.Println(nums)

	const (
		userTotle int = 5
	)

	for i := 0; i < userTotle; i++ {
		var userInput int
		fmt.Printf("输入guess number: ")
		fmt.Scan(&userInput)

		switch {
		case userInput == nums:
			fmt.Println("正确")
			os.Exit(0)
		case userInput < nums:
			fmt.Println("小")
		default:
			fmt.Println("大")
		}
	}
	fmt.Println("ben")

}