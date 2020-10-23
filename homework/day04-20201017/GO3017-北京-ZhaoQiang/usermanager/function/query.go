package function

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	queryPrompts string = "Enter the information you are looking for: "
)

//Query sss
func Query() {
	tmpUsers, _ := getUserMessage(queryPrompts)

	printTable(tmpUsers)
}

func getQueryUserInput(prompts string) string {
	scanner := bufio.NewScanner(os.Stdin)

	var searchString string
	for {
		fmt.Printf(prompts)
		if scanner.Scan() {
			if strings.TrimSpace(scanner.Text()) != "" {
				searchString = scanner.Text()
				break
			} else {
				fmt.Println("输入的是空字符串")
				continue
			}
		}
	}
	
	return searchString
}

func getUserMessage(prompts string) ([]map[string]string, []int) {
	searchString := getQueryUserInput(prompts)
	tempUsers := []map[string]string{}
	tempusersIndex := make([]int, 0, 5)

	for indexIndex, user := range Users {
		for _, userString := range user {
			if userString == searchString {
				tempUsers = append(tempUsers, user)
				tempusersIndex = append(tempusersIndex, indexIndex)
			}
		}
	}

	return tempUsers, tempusersIndex
}
