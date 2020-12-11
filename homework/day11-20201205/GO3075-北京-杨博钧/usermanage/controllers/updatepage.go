package controllers

import (
	"fmt"
	"net/http"
	"html/template"
	"strconv"
	"usermanage/model"
	"usermanage/services"
	"usermanage/utils"
)

func UpdatePage(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("template/updatepage.html")
	if err != nil {
		fmt.Println("加载模板错误:", err)
	}
	updatepage := model.NewAddUpdatePage()
	if r.Method != "POST" {
		updatepage.NameError = ""
		updatepage.PasswdError = ""
		updatepage.BirthdayError = ""
		id, _ := strconv.Atoi(r.FormValue("Id"))
		services.GetPageById(id, updatepage)
		updatepage.Id = id
		tpl.ExecuteTemplate(w, "updatepage.html", updatepage)
	} else {
		err := services.ParseAddUpdateParams(r, updatepage)
		id, _ := strconv.Atoi(r.FormValue("Id"))
		if utils.IsUpdateUserExists(id, r.FormValue("name")){
			updatepage.NameError = "当前用户名已存在,不能进行添加!"
			updatepage.PasswdError = ""
			updatepage.BirthdayError = ""
			tpl.ExecuteTemplate(w, "updatepage.html", updatepage)
		} else if !utils.IsPasswdSame(r){
			updatepage.NameError = ""
			updatepage.PasswdError = "两次输入的密码不一致请重新输入"
			updatepage.BirthdayError = ""
			tpl.ExecuteTemplate(w, "updatepage.html", updatepage)
		} else if err != nil {
			updatepage.NameError = ""
			updatepage.PasswdError = ""
			updatepage.BirthdayError = fmt.Sprintf("%v", err)
			tpl.ExecuteTemplate(w, "updatepage.html", updatepage)
		} else {
			updatepage.NameError = ""
			updatepage.PasswdError = ""
			updatepage.BirthdayError = ""
			services.UpdateUser(id, updatepage)
			http.Redirect(w, r, "/", 302)
		}
	}
}