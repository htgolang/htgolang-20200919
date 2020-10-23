package user

import (
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

var users = []map[string]string{}

/*
	common code
*/
func commonInput() (string, string, string) {
	var name, phone, address string
	fmt.Printf("Please enter user name: ")
	fmt.Scan(&name)
	fmt.Printf("Please enter user contact information: ")
	fmt.Scan(&phone)
	fmt.Printf("Please enter user address: ")
	fmt.Scan(&address)

	return name, phone, address
}

func AllUsersTable(users []map[string]string) {
	t := tablewriter.NewWriter(os.Stdout)
	headers := []string{"ID", "Name", "Phone", "Address"}
	var data = [][]string{}

	for _, user := range users {
		var element = []string{}
		for _, header := range headers {
			element = append(element, user[strings.ToLower(header)])

		}
		data = append(data, element)
	}

	t.SetHeader(headers)
	t.SetAutoFormatHeaders(false)
	t.SetAutoWrapText(false)
	t.SetReflowDuringAutoWrap(false)
	t.AppendBulk(data)
	t.Render()
}

func GenPasswd(pwd string) string {
	md5Value := fmt.Sprintf("%X", md5.Sum([]byte(pwd)))
	return md5Value
}

func add() {
	name, phone, address := commonInput()

	// id++
	maxId := func(users []map[string]string) string {
		if len(users) == 0 {
			return "1"
		} else {
			for i := 0; i < len(users)-1; i++ {
				if users[i]["id"] < users[i+1]["id"] {
					users[i], users[i+1] = users[i+1], users[i]
				}
			}
		}
		max, _ := strconv.Atoi(users[0]["id"])
		return strconv.Itoa(max + 1)
	}(users)
	users = append(users, map[string]string{
		"id":      maxId,
		"name":    name,
		"phone":   phone,
		"address": address,
	})

	AllUsersTable(users)
}

/*
	del user
*/
func del() {
	var id string
	fmt.Printf("Please enter the user id to be deleted: ")
	fmt.Scan(&id)

	for userIndex, user := range users {
		value, ok := user["id"]
		if ok && value == id {
			var _users []map[string]string
			_users = append(_users, user)
			AllUsersTable(_users)
			var isDel string
			fmt.Printf("Confirm delete y/n: ")
			fmt.Scan(&isDel)
			if isDel == "y" {
				users = append(users[:userIndex], users[userIndex+1:]...)
				AllUsersTable(users)
			}
		}
	}
}

/*
	filter user by user id
*/
func filterById(id string) (bool, map[string]string) {
	var flag bool
	var u = make(map[string]string)
	for _, user := range users {
		value, ok := user["id"]
		if ok && value == id {
			flag = true
			u = user
			break
		} else {
			flag = false
		}
	}
	return flag, u
}

/*
	modify
*/
func modify() {
	var id string
	fmt.Printf("Please enter user id: ")
	fmt.Scan(&id)

	flag, user := filterById(id)

	if flag {
		var _users []map[string]string
		_users = append(_users, user)
		AllUsersTable(_users)
		var isModify string
		fmt.Printf("Confirm the changes y/n: ")
		fmt.Scan(&isModify)
		if isModify == "y" {
			name, phone, address := commonInput()
			user["name"] = name
			user["phone"] = phone
			user["address"] = address
			AllUsersTable(users)
		} else {
			AllUsersTable(users)
		}
	} else {
		fmt.Println("No such user...")
	}
}

/*
	filter user by string
*/
func filterByString(str string) (bool, []map[string]string) {
	var flag bool
	var u []map[string]string
	for _, user := range users {
		for _, userValue := range user {
			if userValue == str {
				flag = true
				u = append(u, user)
				break
			}
		}
	}
	return flag, u
}

/*
	query
*/
func query() {
	var queryStr string
	fmt.Printf("Please enter the string you want to query: ")
	fmt.Scan(&queryStr)

	flag, user := filterByString(queryStr)
	if flag {
		AllUsersTable(user)
	} else {
		fmt.Println("Did not match")
	}
}

/*
	help
*/
func Help() {
	t := tablewriter.NewWriter(os.Stdout)
	data := [][]string{
		{"add", "add user"},
		{"delete", "del user"},
		{"modify", "modify user"},
		{"query", "query user"},
		{"exit", "quit"},
		{"help", "help"},
	}
	t.SetHeader([]string{"Command", "Features"})
	t.SetAutoFormatHeaders(false)
	t.SetAutoWrapText(false)
	t.SetReflowDuringAutoWrap(false)
	t.AppendBulk(data)
	t.Render()
}

func Run() {
	for {
		var choiceOne string
		fmt.Printf("Please enter the command: ")
		fmt.Scan(&choiceOne)
		switch choiceOne {
		case "add":
			add()
		case "query":
			query()
		case "del":
			del()
		case "modify":
			modify()
		case "exit", "quit":
			os.Exit(0)
		case "help":
			Help()
		}
	}
}
