package usertools

import "fmt"

const auth string = "4abe77c201ff11663ccdf52fd6ecea86"

// Start is ..
func Start() {
	for i := 3; i >= 0; i-- {
		result := userLogin(auth, i)
		if result {
			i := true

			for i {
				var choose string
				fmt.Print("1: 添加用户\n2: 删除用户\n3: 修改用户\n4: 查询用户\n5：退出\n请输入你要做的事情：")
				fmt.Scan(&choose)

				switch {
				case choose == "1":
					UserAdd()
				case choose == "2":
					UserDel()
				case choose == "3":
					UserModify()
				case choose == "4":
					UserQuery()
				case choose == "5":
					i = false
				}
			}
			break
		}
	}

}
