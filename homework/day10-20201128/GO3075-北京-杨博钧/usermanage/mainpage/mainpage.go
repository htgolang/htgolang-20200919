package mainpage

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"usermanage/model"
	"usermanage/utils"
)

type MainPage struct {
	Template string
	Id string
	Name string
	Sex string
	Addr string
	Tel string
	Birthday string
	Userinfos utils.UserList
}

func NewMainPage(template string) *MainPage {
	return &MainPage{Template: template,}
}

func (this *MainPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	utils.LoadData()
	if r.Method != "POST" {
		this.Id = ""
		this.Name = ""
		this.Sex = ""
		this.Addr = ""
		this.Tel = ""
		this.Birthday = ""
		this.Userinfos = utils.UsersList
		tpl, err := template.ParseFiles(this.Template)
		if err != nil {
			fmt.Println("加载模板出错:", err)
			return
		}
		tpl.ExecuteTemplate(w, "mainpage.html", this)
	} else {
		this.ParseParams(r)
		tpl, err := template.ParseFiles(this.Template)
		if err != nil {
			fmt.Println("加载模板出错:", err)
			return
		}
		tpl.ExecuteTemplate(w, "mainpage.html", this)
	}
}

func (this *MainPage) ParseParams(r *http.Request) {
	this.Userinfos = []*model.User{}
	this.Id = r.FormValue("id")
	this.Name = r.FormValue("name")
	this.Sex = r.FormValue("sex")
	this.Addr = r.FormValue("addr")
	this.Tel = r.FormValue("phone")
	this.Birthday = r.FormValue("birthday")
	for _, v := range utils.UsersList {
		if (strconv.Itoa(v.Id) == this.Id || this.Id == "") &&
			(v.Name == this.Name || this.Name == "") &&
			(v.Sex == this.Sex || this.Sex == "") &&
			(v.Addr == this.Addr || this.Addr == "") &&
			(v.Tel == this.Tel || this.Tel == "") &&
			(v.Birthday == this.Birthday || this.Birthday == "") {
			this.Userinfos = append(this.Userinfos, v)
		}
	}
}