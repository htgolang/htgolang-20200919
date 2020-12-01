package addpage

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"net/http"
	"usermanage/model"
	"usermanage/utils"
)

// 定义结构体存储模板信息、报错信息、输入属性
type AddPage struct {
	Template string
	NameError string
	PasswdError string
	Name string
	IsMale string
	IsFeMale string
	Addr string
	Tel string
	Birthday string
	Passwd string
}

func NewAddPage(template string) *AddPage {
	return &AddPage{Template: template,}
}

func (this *AddPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		this.NameError = ""
		this.PasswdError = ""
		this.Name = ""
		this.IsMale = "checked"
		this.Addr = ""
		this.Tel = ""
		this.Birthday = ""
		this.Passwd = ""
		tpl, err := template.ParseFiles(this.Template)
		if err != nil {
			fmt.Println("加载模板出错:", err)
			return
		}
		tpl.ExecuteTemplate(w, "addpage.html", this)
	} else {
		this.ParseParams(r)
		if _,exist := utils.IsUserExists(r.FormValue("name")); exist {
			this.NameError = "当前用户名已存在,不能进行添加!"
			this.PasswdError = ""
			tpl, err := template.ParseFiles(this.Template)
			if err != nil {
				fmt.Println("加载模板出错:", err)
				return
			}
			tpl.ExecuteTemplate(w, "addpage.html", this)
		} else if !utils.IsPasswdSame(r){
			this.NameError = ""
			this.PasswdError = "两次输入的密码不一致请重新输入"
			tpl, err := template.ParseFiles(this.Template)
			if err != nil {
				fmt.Println("加载模板出错:", err)
				return
			}
			tpl.ExecuteTemplate(w, "addpage.html", this)
		} else {
			this.NameError = ""
			this.PasswdError = ""
			this.AddUser()
			http.Redirect(w, r, "/", 302)
		}
	}
}

// 添加数据
func (this *AddPage) AddUser() {
	var sex string
	if this.IsMale == ""{
		sex = "女"
	} else {
		sex = "男"
	}
	userInfo := model.User{
		Id: utils.GetMaxId() + 1,
		Name: this.Name,
		Sex: sex,
		Addr: this.Addr,
		Tel: this.Tel,
		Birthday: this.Birthday,
		Password: fmt.Sprintf("%x", md5.Sum([]byte(this.Passwd))),
	}
	utils.UsersList = append(utils.UsersList, &userInfo)
	utils.SaveData()
}

// 将请求解析到结构体中
func (this *AddPage) ParseParams(r *http.Request) {
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