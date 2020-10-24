package utils

import "fmt"

// 输入信息
func Input(prompt string) string {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	return text
}
