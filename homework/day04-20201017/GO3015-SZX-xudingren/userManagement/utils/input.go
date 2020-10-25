package utils

import "fmt"

func Input(prompt, arg string) string {
	fmt.Print(prompt)
	fmt.Scanln(&arg)
	return arg
}
