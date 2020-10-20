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

func wFile(filename string, userList map[string]string) bool {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("file Error")
		return false
	}
	defer file.Close()
	fmt.Println(userList)
	jsonCode := json.NewEncoder(file)
	jsonCode.Encode(userList)
	fmt.Println(jsonCode)

	return true

}

// ok
func rFile(filename, sinfo string, userList map[string]string) (map[string]string, bool) {
	jsontxt, _ := ioutil.ReadFile(filename)
	if err := json.Unmarshal(jsontxt, &userList); err != nil {
		fmt.Println(err)
	}
	return userList, true

}

func addUser(filename string, userList map[string]string) {
	name, age, phone := scanInput()
	userList = make(map[string]string)
	userList["name"] = name
	userList["age"] = age
	userList["phone"] = phone
	status := wFile(filename, userList)
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

func searUser(filename string, userList map[string]string) {
	var sinfo string
	var flag bool
	userList = make(map[string]string)
	fmt.Print("Please input search username:\t")
	fmt.Scan(&sinfo)
	_, status := rFile(filename, sinfo, userList)
	// fmt.Println(status)
	if status == true {
		flag = true
	} else {
		flag = false
	}
	// if strings.Contains(userList["name"], sinfo) || strings.Contains(userList["age"], sinfo) || strings.Contains(userList["phone"], sinfo) {
	// if strings.Contains(userList["name"], sinfo) {
	if userList["name"] == sinfo {
		fmt.Println(strings.Repeat("-", 60))
		fmt.Printf("status:\t[%v]\t search user info\n", flag)
		fmt.Println(strings.Repeat("-", 60))
		title := fmt.Sprintf("%-10s|%5s|%20s|", "NAME", "AGE", "TEL")
		fmt.Println(title)
		fmt.Printf("%-10s|%5s|%20s|\n", userList["name"], userList["age"], userList["phone"])

	} else {
		fmt.Println("the user does not exist")
	}

}

func main() {

	userList := make(map[string]string)
	var input_nu int

	fmt.Printf(`
	

			用户管理系统

[%s]

		1. 增加用户
		2. 查询用户
		3. 删除用户
		4. 修改用户
	`, strings.Repeat("*", 100))
	// addUser("userInfo.json", userList)
	// searUser("userInfo.json", userList)
	fmt.Println("")
	fmt.Print("选择功能按键 :")
	fmt.Scan(&input_nu)

	switch input_nu {
	case 1:
		addUser("userInfo.json", userList)
	case 2:
		searUser("userInfo.json", userList)
	case 3:
		return
	case 4:
		return
	default:
		fmt.Println("parameter error")

	}

}
