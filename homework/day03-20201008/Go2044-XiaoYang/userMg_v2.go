package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func scanInput() (string, string, string) {
	var (
		name  string
		age   string
		phone string
	)

	fmt.Printf("Input name:\t")
	fmt.Scan(&name)
	fmt.Printf("Input age:\t")
	fmt.Scan(&age)
	fmt.Printf("Input phone:\t")
	fmt.Scan(&phone)
	return name, age, phone
}

func wFile(filename string, userList []map[string]string) bool {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatal("file Error")
		return false
	}
	defer file.Close()
	jsonCode := json.NewEncoder(file)
	jsonCode.Encode(userList)
	return true

}

func rFile(filename string, users []map[string]string) ([]map[string]string, bool) {
	jsontxt, _ := ioutil.ReadFile(filename)
	if err := json.Unmarshal(jsontxt, &users); err != nil {
		fmt.Println(err)
	}
	// fmt.Println(users)
	return users, true

}

func addUser(filename string, users []map[string]string) {
	userList := map[string]string{}
	name, age, phone := scanInput()
	userList = map[string]string{
		"name":  name,
		"age":   age,
		"phone": phone,
	}

	users = append(users, userList)
	status := wFile(filename, users)
	if status == true {
		fmt.Println(strings.Repeat("-", 60))
		fmt.Println("['+']\tAdd user sucessful ")
		fmt.Println(strings.Repeat("-", 60))
		title := fmt.Sprintf("%-10s|%5s|%20s|", "NAME", "AGE", "TEL")
		fmt.Println(title)
		fmt.Printf("%-10s|%5s|%20s|\n", userList["name"], userList["age"], userList["phone"])
	} else {
		fmt.Println("Add user fileds")
	}
}

func searUser(filename string, users []map[string]string) {
	var sinfo string
	var flag bool
	fmt.Print("Please input search username:\t")
	fmt.Scan(&sinfo)
	users, status := rFile(filename, users)
	if status == true {
		flag = true
	} else {
		flag = false
	}
	title := fmt.Sprintf("%-10s|%5s|%20s|", "NAME", "AGE", "TEL")
	fmt.Println(title)
	fmt.Printf("status:\t[%v]\t search user info\n", flag)
	for _, userList := range users {
		if userList["name"] == sinfo {
			fmt.Println(strings.Repeat("-", 60))
			fmt.Printf("%-10s|%5s|%20s|\n", userList["name"], userList["age"], userList["phone"])
			fmt.Println(strings.Repeat("-", 60))
		}
	}

}

func main() {
	users := []map[string]string{}
	filename := "test_v1.json"
	var input_nu int

	fmt.Printf(`
	

			用户管理系统

[%s]

		1. 增加用户
		2. 查询用户
		3. 删除用户
		4. 修改用户
	`, strings.Repeat("*", 100))
	fmt.Println("")
	fmt.Print("选择功能按键 :")
	fmt.Scan(&input_nu)

	switch input_nu {
	case 1:
		addUser(filename, users)
	case 2:
		searUser(filename, users)
	case 3:
		return
	case 4:
		return
	default:
		fmt.Println("parameter error")

	}

}
