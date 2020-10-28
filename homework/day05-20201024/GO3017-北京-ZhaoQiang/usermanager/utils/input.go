package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/howeyc/gopass"
)

//GetCommandLineInput 获取命令行的输入
func GetCommandLineInput(useraccount string) string {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("[userManager-%s] $ ", useraccount)
		if scanner.Scan() {
			if strings.TrimSpace(scanner.Text()) == "" {
				continue
			}
			return strings.TrimSpace(scanner.Text())
		}
	}
}

//GetUserInputPassWD get user input passwd
func GetUserInputPassWD(prompt string) string {
	fmt.Printf(prompt)

	passwd, err := gopass.GetPasswd()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return string(passwd)
}

//GetUserInputString 获取用户输入信息 返回一个字符串
func GetUserInputString(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf(prompt)
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text())
	}
	return ""
}

//GetUserInputInt 获取用户输入信息 返回一个int
func GetUserInputInt(prompt string) int {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(prompt)
		if scanner.Scan() {
			str := strings.TrimSpace(scanner.Text())
			if num, err := strconv.Atoi(str); err == nil {
				return num
			}
			fmt.Printf("enter ID is not numbers, try again \n\n")
			continue
		}
	}
}

//GetUserInputTime 获取用户输入信息 ,返回一个字符串
func GetUserInputTime(prompt string) time.Time {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(prompt)
		if scanner.Scan() {
			t, err := time.Parse("2006/01/02", strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("输入的格式为： 2006/01/02 , try again.")
				continue
			}
			return t
		}
	}

}
