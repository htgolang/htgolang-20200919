package controllers

import (
	"fmt"
	"net/http"
	"html/template"
	"usermanage/model"
	"usermanage/services"
	"usermanage/utils"
)

func AddPage(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("template/addpage.html")
	if err != nil {
		fmt.Println("加载模板错误:", err)
	}
	addpage := model.NewAddUpdatePage()
	if r.Method != "POST" {
		tpl.ExecuteTemplate(w, "addpage.html", addpage)
	} else {
		err := services.ParseAddUpdateParams(r, addpage)
		if utils.IsAddUserExists(r.FormValue("name")){
			addpage.NameError = "当前用户名已存在,不能进行添加!"
			addpage.PasswdError = ""
			addpage.BirthdayError = ""
			tpl.ExecuteTemplate(w, "addpage.html", addpage)
		} else if !utils.IsPasswdSame(r){
			addpage.NameError = ""
			addpage.PasswdError = "两次输入的密码不一致请重新输入"
			addpage.BirthdayError = ""
			tpl.ExecuteTemplate(w, "addpage.html", addpage)
		} else if err != nil {
			addpage.NameError = ""
			addpage.PasswdError = ""
			addpage.BirthdayError = fmt.Sprintf("%v", err)
			tpl.ExecuteTemplate(w, "addpage.html", addpage)
		} else {
			addpage.NameError = ""
			addpage.PasswdError = ""
			addpage.BirthdayError = ""
			services.AddUser(addpage)
			http.Redirect(w, r, "/", 302)
		}
	}
}
