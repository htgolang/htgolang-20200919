package controllers

import (
	"fmt"
	"net/http"
	"html/template"
	"usermanage/model"
	"usermanage/services"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("template/mainpage.html")
	if err != nil {
		fmt.Println("加载模板出错:", err)
		return
	}
	mainpage := model.NewMainPage()
	if r.Method != "POST" {
		mainpage.Userinfos = services.GetAllUser()
		tpl.ExecuteTemplate(w, "mainpage.html", mainpage)
	} else {
		err := services.ParseQueryParams(r, mainpage)
		if err != nil {
			mainpage.Error = fmt.Sprintf("%v", err)
			mainpage.Userinfos = services.GetAllUser()
			tpl.ExecuteTemplate(w, "mainpage.html", mainpage)
		} else {
			mainpage.Error = ""
			mainpage.Userinfos = services.GetQueryUser(mainpage)
			tpl.ExecuteTemplate(w, "mainpage.html", mainpage)
		}
	}
}