package funcs

import(
	"fmt"
)

func Show() {
	menu := `--------------------------------------------------
	***********用户管理系统***********
	    a or add     ）添加用户
  	    m or modify  ）修改用户
	    d or delete  ）删除用户
	    l or list    ）用户列表
	    q or query   ）搜索用户
	    h or help    ）帮助信息
	    exit or quit ）退出系统
	********************************
	`
	fmt.Println(menu)
}