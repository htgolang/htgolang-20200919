package funcs

import "fmt"

// 初始视图，定义帮助信息
func View() {

	//使用反引号来定义多行字符串
	str := `--------------------------------------------------
	***********用户管理系统***********
	    a or add     ）添加用户
	    d or delete  ）删除用户
  	    m or modify  ）修改用户
	    l or list    ）用户列表
	    q or query   ）搜索用户
	    h or help    ）帮助信息
	    exit or quit ）退出系统
	********************************
	`
	fmt.Println(str)
}
