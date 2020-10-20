package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 定义增加用户函数
func addUser(users []map[string]string) {

	// 用户信息变量
	var (
		id    string
		name  string
		phone string
		email string
	)

	// 从标准输入获取用户填入信息
	fmt.Println("请输入姓名：")
	fmt.Scan(&name)
	fmt.Println("请输入电话号码：")
	fmt.Scan(&phone)
	fmt.Println("请输入Email: ")
	fmt.Scan(&email)

	// 初始化一个用户映射，并组装用户信息
	u := map[string]string{"name": name, "phone": phone, "email": email}
	// 生成id，users中没有数据时id就为1，若有数据，找出id的最大值，那新用户信息的id为 id + 1
	if len(users) == 0 { // 用户列表为空时，新增加用户id为“1”
		id = strconv.Itoa(1)
		u["id"] = id
	} else {
		ids := []int{} // 存放已有用户的id值
		for _, v := range users {
			if res, err := strconv.Atoi(v["id"]); err == nil {
				ids = append(ids, res)
			} else {
				fmt.Printf("不能转换为数字id %q\n", v["id"])
			}
		}
		// 找出已有用户的id最大值
		// 假设nums[0]为最大值
		maxNum := ids[0]
		for i := 1; i < len(ids); i++ {
			if ids[i] > maxNum {
				maxNum = ids[i]
			}
		}
		// 获取增加用户的id
		u["id"] = strconv.Itoa(maxNum + 1)
	}
	// 向users切片映射中增加用户信息
	users = append(users, u)
	fmt.Println(users)

}

func dellUserByID(users []map[string]string) {
	// 接收用户输入id
	var userID string
	fmt.Println("请输入要删除用户的id：")
	fmt.Scan(&userID)
	// 较验证用户输入是否为id是否能被转换成int类型
	if _, err := strconv.Atoi(userID); err == nil { // 能成功转换成int，那用户输入id符合要求
		// 较验id是否在users映射中，如果存在再次确认后可删除，如果不存在则打印提示信息
		for i, v := range users {
			if v["id"] == userID {
				var action string
				fmt.Println("确定要删除该用户吗 [y|n] ")
				fmt.Scan(&action)
				if strings.ToLower(action) == "y" { // 确认后进行删除操作
					dst := users[i:len(users)]
					src := users[i+1 : len(users)]
					copy(dst, src)
					users = users[0 : len(users)-1]
				} else if strings.ToLower(action) == "n" {
					fmt.Println("你放弃了操作！")
				} else {
					fmt.Println("请输入正确的指令。")
				}
				break
			}
			// 循环到最后一个元素都没有break时，说明用户输入的id在users中没有
			if i == len(users)-1 {
				fmt.Println("未找到id")
			}
		}

	}
	fmt.Println(users)
}

func modifyUserByID(users []map[string]string) {
	// 修改用户信息
	var userID string
	fmt.Println("请输入要修改用户信息的id: ")
	fmt.Scan(&userID)
	// 判断id是否存在
	for i, v := range users {
		if userID == v["id"] {
			// 打印该用户信息
			fmt.Println(users[i])
			// 确放是否修改该用户
			var action string
			fmt.Println("确认修改该用户信息 [y|n]")
			fmt.Scan(&action)
			if strings.ToLower(action) == "y" {
				var (
					name  string
					phone string
					email string
				)
				fmt.Println("请输入新用户名，为空将保留原用户名：")
				fmt.Scan(&name)
				fmt.Println("请输入新电话号码，为空将保留原电话号码：")
				fmt.Scan(&phone)
				fmt.Println("请输入新的email地址，为空将保留原email：")
				fmt.Scan(&email)
				if strings.TrimSpace(name) != "" {
					v["name"] = name
				}
				if strings.TrimSpace(phone) != "" {
					v["phone"] = phone
				}
				if strings.TrimSpace(email) != "" {
					v["email"] = email
				}
			}
			break
		}
		if i == len(users)-1 {
			fmt.Println("未找到id")
		}
	}
	fmt.Println(users)
}

func queryByStr(users []map[string]string) {
	// 关键字符查找用户信息，有匹配就输出用户信息
	var keyWord string
	fmt.Println("请输入查找字符：")
	fmt.Scan(&keyWord)
	var filterUsers []map[string]string // 存放查找到的用户信息
	for _, m := range users {
		for _, v := range m {
			if strings.Contains(v, keyWord) {
				filterUsers = append(filterUsers, m)
			}
		}
	}
	// 输出查询到的用户信息
	if len(filterUsers) > 0 {
		for _, m := range filterUsers {
			for k, v := range m {
				fmt.Printf("%s : %s\n", k, v)
			}
		}
	} else {
		fmt.Println("未检索到用户信息。")
	}

}

func main() {
	// 定义一个切片映射数据结构
	var users = []map[string]string{}
	// 增加一个初始的用户信息，为了测试id + 1的效果
	users = append(users, map[string]string{"id": "10", "name": "张三", "phone": "123456", "email": "zs@qq.com"})
	users = append(users, map[string]string{"id": "20", "name": "李五", "phone": "654321", "email": "lw@qq.com"})

	var uChoice int
	fmt.Println(`
	1：增加用户
	2：删除用户
	3：修改用户信息
	4：查询用户`)
	fmt.Scan(&uChoice)
	switch {
	case uChoice == 1:
		addUser(users)
	case uChoice == 2:
		dellUserByID(users)
	case uChoice == 3:
		modifyUserByID(users)
	case uChoice == 4:
		queryByStr(users)
	}

}
