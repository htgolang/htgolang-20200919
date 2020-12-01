package main

import (
	"usermanage/deletepage"
	"usermanage/mainpage"
	"net/http"
	"usermanage/addpage"
	"usermanage/updatepage"
)

func main() {
	// 定义端口绑定5849
	addr := ":5849"

	// 主页面，展示所有用户列表
	mainPage := mainpage.NewMainPage("template/mainpage.html")
	http.Handle("/", mainPage)

	// 新增页面
	addPage := addpage.NewAddPage("template/addpage.html")
	http.Handle("/add/", addPage)

	// 删除页面
	deletePage := deletepage.NewDeletePage()
	http.Handle("/delete/", deletePage)

	// 更新页面
	updatePage := updatepage.NewUpdatePage("template/updatepage.html")
	http.Handle("/update/", updatePage)

	// 启动监听服务
	http.ListenAndServe(addr, nil)
}

