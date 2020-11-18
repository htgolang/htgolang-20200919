package utils

import (
	"fmt"
	"strings"
)

// 让用户输入信息并返回信息值
func Input(prompt string) string {
	var text string
	fmt.Println(prompt)
	_, _ = fmt.Scan(&text)
	return strings.TrimSpace(text)
}

// 让用户输入信息并返回信息值
func IntInput(prompt string) int {
	var text int
	fmt.Println(prompt)
	_, _ = fmt.Scan(&text)
	return text
}
