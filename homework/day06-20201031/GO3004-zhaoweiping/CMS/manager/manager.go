package manager

import (
	"CMS/utils"
	"fmt"
)

//路由
var routers = map[string]func(){}

func Register(op string, callback func()) {
	if _, ok := routers[op]; ok {
		panic(fmt.Sprintf("指令%s 已经存在", op))
	}
	routers[op] = callback
}

func Registers(ops []string, callback func()) {
	for _, op := range ops {
		if _, ok := routers[op]; ok {
			panic(fmt.Sprintf("指令%s 已经存在", op))
		}
		// fmt.Println(op)
		routers[op] = callback
	}

}
func Run() {
	for {
		text := utils.Input("请输入指令：")
		// if text == "quit" {
		// 	fmt.Println("结束本次服务，再见客官！！！")
		// 	break
		// }
		if action, ok := routers[text]; ok {
			action()
		} else {
			fmt.Println("指令错误！！！")
		}
	}
}
