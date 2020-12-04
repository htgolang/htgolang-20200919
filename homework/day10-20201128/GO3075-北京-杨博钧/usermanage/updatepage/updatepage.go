package updatepage

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"strconv"
	"usermanage/utils"
	"html/template"
)

type UpdatePage struct {
	Template string
	NameError string
	PasswdError string
	Id int
	Name string
	IsMale string
	IsFeMale string
	Addr string
	Tel string
	Birthday string
	Passwd string
}

func NewUpdatePage(template string) *UpdatePage {
	return &UpdatePage{Template: template,}
}

func (this *UpdatePage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		this.NameError = ""
		this.PasswdError = ""
		id, _ := strconv.Atoi(r.FormValue("Id"))
		this.ParseUserInfo(id)
		tpl, err := template.ParseFiles(this.Template)
		if err != nil {
			fmt.Println("加载模板出错:", err)
			return
		}
		tpl.ExecuteTemplate(w, "updatepage.html", this)
	} else {
		this.ParseParams(r)
		if i,exist := utils.IsUserExists(r.FormValue("name")); exist && utils.UsersList[i].Id != this.Id {
			this.NameError = "当前用户名已存在,不能进行更新!"
			this.PasswdError = ""
			tpl, err := template.ParseFiles(this.Template)
			if err != nil {
				fmt.Println("加载模板出错:", err)
				return
			}
			tpl.ExecuteTemplate(w, "updatepage.html", this)
		} else if !utils.IsPasswdSame(r){
			this.NameError = ""
			this.PasswdError = "两次输入的密码不一致请重新输入"
			tpl, err := template.ParseFiles(this.Template)
			if err != nil {
				fmt.Println("加载模板出错:", err)
				return
			}
			tpl.ExecuteTemplate(w, "updatepage.html", this)
		} else {
			this.NameError = ""
			this.PasswdError = ""
			this.UpdateUser()
			http.Redirect(w, r, "/", 302)
		}
	}
}

// 定义更新数据函数
func (this *UpdatePage) UpdateUser() {
	User := utils.GetUserInfoById(this.Id)
	User.Id = this.Id
	User.Name = this.Name
	if this.IsMale == "" {
		User.Sex = "女"
	} else {
		User.Sex = "男"
	}
	User.Addr = this.Addr
	User.Birthday = this.Birthday
	User.Password = fmt.Sprintf("%x",md5.Sum([]byte(this.Passwd)))
	utils.SaveData()
}

// 将请求解析到结构体中
func (this *UpdatePage) ParseUserInfo(Id int) {
	User := utils.GetUserInfoById(Id)
	this.Id = Id
	this.Name = User.Name
	if User.Sex == "男" {
		this.IsMale = "checked"
		this.IsFeMale = ""
	} else {
		this.IsMale = ""
		this.IsFeMale = "checked"
	}
	this.Addr = User.Addr
	this.Tel = User.Tel
	this.Birthday = User.Birthday
	//this.Passwd = User.Password
}

// 将请求解析到结构体中
func (this *UpdatePage) ParseParams(r *http.Request) {
	this.Name = r.FormValue("name")
	if r.FormValue("sex") == "1" {
		this.IsMale = "checked"
		this.IsFeMale = ""
	} else {
		this.IsMale = ""
		this.IsFeMale = "checked"
	}
	this.Addr = r.FormValue("addr")
	this.Tel = r.FormValue("phone")
	this.Birthday = r.FormValue("birthday")
	this.Passwd = r.FormValue("password")
}