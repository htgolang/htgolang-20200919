package function

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// AddRun run addfuncation some func
func AddRun() {
	//生成Users[] 格式的 map用户信息
	id := getID()
	usermess := getUserInputMessage() //用户输入的信息
	usermess["id"] = id               // 添加id字段
	Users = append(Users, usermess)

	tmpUsers := []map[string]string{}
	tmpUsers = append(tmpUsers, usermess)
	printTable(tmpUsers)
}

func getID() string {
	var n int64 = 0
	for _, v := range Users {
		idst := v["id"]

		idint, _ := strconv.ParseInt(idst, 10, 64)

		if n < idint {
			n = idint
		}
	}
	n++
	idStr := strconv.FormatInt(n, 10)
	return idStr
}

func getUserInputMessage() map[string]string {
	ut := [3]string{"name", "tel", "addr"}
	usermessage := map[string]string{"name": "", "tel": "", "addr": ""}

	scanner := bufio.NewScanner(os.Stdin)

	for _, k := range ut {
		for {
			fmt.Printf("user %s: ", k)
			if scanner.Scan() {
				//用户输入信息合法性判断
				s := strings.TrimSpace(scanner.Text())
				if s != "" {
					usermessage[k] = s
					break
				} else {
					fmt.Println("不能是空字符串")
					continue
				}
			}
		}
	}
	return usermessage
}
