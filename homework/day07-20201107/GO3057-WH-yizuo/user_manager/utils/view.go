package utils

import "fmt"

// 初始帮助信息
func View() {
	//使用反引号来定义多行字符串
	str := `--------------------------------------------------
	***********用户管理系统***********
	    a or add     ）添加用户
  	    m or modify  ）修改用户
	    d or delete  ）删除用户
	    l or list    ）用户列表
	    query        ）搜索用户
	    h or help    ）帮助信息
	    q or exit or quit ）退出系统
	********************************
	`
	fmt.Println(str)
}