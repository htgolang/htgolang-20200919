package main

import (
	"fmt"
	"strconv"
	"strings"
)

var users []map[string]string

//添加用户函数
func addUser() {
	for {
		var name, tel, addr, yesOrno string
		fmt.Printf("添加用户:\n请输入用户名:")
		_, _ = fmt.Scanln(&name)
		fmt.Printf("请输入联系电话:")
		_, _ = fmt.Scanln(&tel)
		fmt.Printf("请输入地址:")
		_, _ = fmt.Scanln(&addr)
		fmt.Printf("您录入的信息:\n用户名: %s. 联系电话: %s. 地址: %s.\n确认添加?[y/n]", name, tel, addr)
		_, _ = fmt.Scanln(&yesOrno)
		if yesOrno == "y" {
			// 调用添加用户的函数
			//users = addUser(name, tel, addr, users)
			aduser := make(map[string]string)
			id := func() int {
				if len(users) == 0 {
					return 1
				}
				id := 1
				for i := 0; i < len(users); i++{
					tid,_ := strconv.Atoi(users[i]["id"])
					if id < tid{
						id = tid
					}
				}
				return id +1
			}()
			aduser["id"] = strconv.Itoa(id)
			aduser["name"] = name
			aduser["tel"] = tel
			aduser["addr"] = addr
			users = append(users,aduser)
			var q string
			fmt.Printf("q.退出, 任意键继续....")
			_, _ = fmt.Scanln(&q)
			if q == "q"{
				break
			}else {
				continue
			}

		} else if yesOrno == "n" {
			var q string
			fmt.Printf("q.退出, 任意键继续....")
			_, _ = fmt.Scanln(&q)
			if q == "q"{
				break
			}else {
				continue
			}
		}
	}
}

//通过ID查找用户
func findByid(id int) (index int,err error) {
	for k,v := range users{
		uid,err := strconv.Atoi(v["id"])
		if err != nil {
			_ = fmt.Errorf("id转换错误:%s",err)
			return -1, err
		}
		if uid == id {
			return k,nil
		}
	}
	return -1, fmt.Errorf("未找到")

}

//关键字查找用户函数
func queryUser()  {
	var str string
	var f int
	fmt.Printf("请出入关键字:")
	_,_ = fmt.Scanln(&str)
	for _,v := range users{
		if strings.Contains(v["name"],str) || strings.Contains(v["tel"],str)|| strings.Contains(v["addr"],str) {
			fmt.Printf("找到了!!!\nId: %s\n用户名: %s\n联系电话: %s\n地址: %s\n",v["id"],v["name"],v["tel"],v["addr"])
			f = 1
			break
		}
	}
	if f == 0 {
		fmt.Println("未找到~~~~~~~~~~~~~~~~~~~")
	}

}


// 删除用户函数
func delUser() {
	for {
		var uid string
		fmt.Printf("提示:按q 退出\n请输入用户ID:")
		_, _ = fmt.Scanln(&uid)
		if uid == "q"{
			break
		}
		fid,_ := strconv.Atoi(uid)
		index, err := findByid(fid)
		fmt.Println(index)
		if err != nil {
			fmt.Println("未找到用户id:", uid)
		} else {
			fmt.Printf("Id: %s, 用户名: %s, 联系方式: %s, 地址:%s.\n确认删除?(q 退出.)[y/n/q]", users[index]["id"], users[index]["name"], users[index]["tel"], users[index]["addr"])
			var yOrn string
			_, _ = fmt.Scanln(&yOrn)
			if yOrn == "y" {
				if index == 0 && len(users) == 1{
					users = []map[string]string{}
					break
				}
				if index == len(users) -1 {
					users = users[:index]
					break
				}
				users = append(users[:index],users[index+1:]...)
				break
			}else if yOrn == "n"{
				continue
			}else if yOrn == "q"{
				break
			}else {
				break
			}
		}
	}
}

// 修改用户函数
func modifyUser()  {
	for {
		var input string
		fmt.Printf("提示:按q 退出\n请输入用户ID:")
		_, _ = fmt.Scanln(&input)
		if input == "q" {
			break
		}
		uid, _ := strconv.Atoi(input)
		index, err := findByid(uid)
		if err != nil {
			fmt.Println("未找到ID:",uid)
			continue
		}
		fmt.Printf("Id: %s, 用户名: %s, 联系方式: %s, 地址:%s.\n", users[index]["id"], users[index]["name"], users[index]["tel"], users[index]["addr"])
		var name,tel,addr,yOrn string
		fmt.Printf("添加用户:\n请输入新用户名:")
		_, _ = fmt.Scanln(&name)
		fmt.Printf("请输入新联系电话:")
		_, _ = fmt.Scanln(&tel)
		fmt.Printf("请输入新地址:")
		_, _ = fmt.Scanln(&addr)
		fmt.Printf("用户名: %s ---> %s\n联系电话: %s ---> %s\n地址: %s ---> %s\n确认以上修改?[y/n]:",
			users[index]["name"], name, users[index]["tel"], tel, users[index]["addr"], addr)
		_,_ = fmt.Scanln(&yOrn)
		if yOrn == "y"{
			users[index]["name"] = name
			users[index]["tel"] = tel
			users[index]["addr"] = addr
			fmt.Printf("修改成功. 继续请按 c. 任意键退出...\n")
			var t string
			_,_ = fmt.Scanln(&t)
			if t == "c"{
				continue
			}
			break

		}else if yOrn == "n"{
			continue
		}
	}
}

func main()  {

	quit := true
	for quit {
		var s int
		fmt.Printf("1.添加用户.\n2.删除用户.\n3.修改用户.\n4.查找用户.\n5.退出\n请输入序号(1~5):")
		_,_ = fmt.Scanln(&s)
		switch s {
		// 增加用户
		case 1:
			addUser()
			fmt.Println(users)
		// 删除用户
		case 2:
			delUser()
			fmt.Println(users)
		// 修改用户
		case 3:
			modifyUser()
			fmt.Println(users)
		// 查找用户
		case 4:
			queryUser()
		// 退出
		case 5:
			fmt.Println("Bey...")
			quit = false

		default:
			fmt.Println("输入错误...")
		}
	}
}
