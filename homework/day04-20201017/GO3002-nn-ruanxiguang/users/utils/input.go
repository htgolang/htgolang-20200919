package utils

import "fmt"

func Input(enter string) string{
	var text string
	fmt.Print(enter)
	fmt.Scan(&text)
	return text
}
