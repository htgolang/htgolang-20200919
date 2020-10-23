package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var users []map[string]string

// var name, tel, addr string = "", "", ""

func addUser() {
	m := map[string]string{}
	var n, t, a string
START:
	fmt.Println("请输入：姓名  电话  住址")
	fmt.Scanln(&n, &t, &a)
	// fmt.Println("len(users):", len(users))
	if n == "" {
		fmt.Println("先给洒家报上名来！")
		goto START
	}
	if len(users) == 0 {
		rand.Seed(time.Now().Unix())
		m["Id"] = strconv.Itoa(rand.Intn(100))
	} else {
		// fmt.Println(users[len(users)-1]["Id"])
		tmp, _ := strconv.Atoi(users[len(users)-1]["Id"])
		// fmt.Printf("%T %v\n", tmp, tmp)

		//找出最大ID
		for _, m := range users {
			mid, _ := strconv.Atoi(m["Id"])
			if mid > tmp {
				tmp = mid
			}
		}
		// fmt.Println(tmp)
		m["Id"] = strconv.Itoa(tmp + 1)
		// fmt.Println(m["Id"])
	}

	m["Name"] = n
	m["Tel"] = t
	m["Addr"] = a
	users = append(users, m)
	fmt.Printf("成功增加一个用户：")
	fmt.Printf("ID:%s\tName:%s\tTel:%s\tAddr:%s\n", m["Id"], m["Name"], m["Tel"], m["Addr"])
}
func delUser() {
	fmt.Println("请输入要删除的用户ID：")
	tmpid := -1
	fmt.Scanln(&tmpid)

	for idx, m := range users {
		mid, _ := strconv.Atoi(m["Id"])
		// fmt.Println(tmpid, mid)
		if mid == tmpid {
			users = append(users[:idx], users[idx+1:]...)
			fmt.Println("删除成功！")
			return
		}
	}
	fmt.Println("未找到此ID的用户。")
}
func updateUser() {
	fmt.Println("请输入要修改的用户ID：")
	tmpid := -1
	fmt.Scanln(&tmpid)

	for _, m := range users {
		mid, _ := strconv.Atoi(m["Id"])
		if mid == tmpid {
			fmt.Println("请输入要修改的用户信息：\nName    Tel     Addr")
			n, t, a := "", "", ""
			fmt.Scanln(&n, &t, &a)
			m["Name"], m["Tel"], m["Addr"] = n, t, a
			fmt.Println("修改结果：")
			fmt.Printf("ID:%s\tName:%s\tTel:%s\tAddr:%s\n", m["Id"], m["Name"], m["Tel"], m["Addr"])
			return
		}
	}
	fmt.Println("未找到此ID的用户。")
}
func findUser() {
	fmt.Println("请输入查询关键字：")
	keyword := ""
	fmt.Scanln(&keyword)

	fmt.Println("查找结果：")
	found := 0
	for _, m := range users {
		switch {
		case strings.Contains(m["Name"], keyword):
			fmt.Printf("%s\t%s\t%s\t%s\n", m["Id"], m["Name"], m["Tel"], m["Addr"])
			found = 1
		case strings.Contains(m["Tel"], keyword):
			fmt.Printf("%s\t%s\t%s\t%s\n", m["Id"], m["Name"], m["Tel"], m["Addr"])
			found = 1
		case strings.Contains(m["Addr"], keyword):
			fmt.Printf("%s\t%s\t%s\t%s\n", m["Id"], m["Name"], m["Tel"], m["Addr"])
			found = 1
		}
	}
	//循环结束后未找到
	if found == 0 {
		fmt.Println("未找到相关的用户信息。")
	}
}

func listUsers() {
	fmt.Println("--------- 当前用户列表 --------\nId\tName\tTel\tAddr")
	for _, v := range users {
		fmt.Printf("%s\t%s\t%s\t%s\n", v["Id"], v["Name"], v["Tel"], v["Addr"])
	}
}

func main() {
	fmt.Println(`
=========================
-----简易用户管理程序-----
  ------可选操作：-----
增 / 删 / 改 / 查 / 退出
`)
	for {
		fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=")
		fmt.Println("*****请输入操作选项*****")
		ops := ""
		fmt.Scanln(&ops)
		switch ops {
		case "增":
			addUser()
		case "删":
			delUser()
		case "改":
			updateUser()
		case "查":
			findUser()
		case "退出":
			return
		default:
			fmt.Println("输入有误，请重新输入")
			listUsers()
		}
	}

}
