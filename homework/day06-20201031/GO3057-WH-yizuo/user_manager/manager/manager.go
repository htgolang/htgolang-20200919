package manager

import (
	"fmt"
	"strings"
	"yizuo/controllers"
	"yizuo/models"

	"yizuo/utils"
)

// 创建路由，密名函数
var routers = map[string]func(){}

// 添加指令至对应的路由中
func Register(op string, callback func()) {
	if _, ok := routers[op]; ok {
		panic(fmt.Sprintf("指令%s已经存在", op))
	}
	routers[op] = callback
}

// 运行程序
func Run() {
	// 帮助信息
	utils.View()
	// 用户登录系统，如果登录失败则打印错误信息并推出。
	if !controllers.UserLoginAuth() {
		fmt.Println("密码错误次数超过三次，已退出。Bay~，")
		return
	}
	// 主逻辑
	for {
		text := utils.Input("请输入指令: ")
		if text == "exit" || text == "quit" || text == "q" {
			models.WritesUsersDataToCsv()
			fmt.Println("保存用户数据并退出！")
			break
		}
		// fmt.Println(text)
		if action, ok := routers[strings.ToLower(text)]; !ok {
			// fmt.Println(action, ok)
			fmt.Println("指令错误")
		} else {
			action()
		}
	}
}
