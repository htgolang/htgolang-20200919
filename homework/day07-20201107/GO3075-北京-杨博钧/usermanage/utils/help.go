package utils

import "fmt"

func Help() {
	fmt.Println(`
用户管理系统使用提示:
    运行 -init 进行初始化
    list 命令打印当前所有用户信息
    add 命令添加用户
    del 命令删除用户
    upd 命令修改单个用户的信息
    query 命令查询特定数据的用户信息
    exit 命令退出
    help 命令查看所有命令帮助
--------------------------------
`)
}
