package utils

import "fmt"

// 让用户输入信息并返回信息值
func Input(prompt string) string {
	var text string
	fmt.Print(prompt)
	_, _ = fmt.Scan(&text)
	return text
}
