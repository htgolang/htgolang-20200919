package controllers

import (
	"CMS/models"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/spf13/cast"
)

type Cookie struct {
	Name string
	Sex  int
	Addr string
}

var DbType string

var user models.Users

func GetUsers(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("template/user.html"))
	tpl.ExecuteTemplate(w, "user.html", models.GetUsers())
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl := template.Must(template.ParseFiles("template/create.html"))
		tpl.ExecuteTemplate(w, "create.html", nil)
	} else {
		fmt.Println(r.FormValue("sex"), "add")
		if models.AddUser(
			r.FormValue("name"),
			r.FormValue("sex") == "1",
			r.FormValue("addr"),
		) {
			SaveDb()
			fmt.Fprintf(w, `{"result":"sucess"}`)

		} else {
			fmt.Fprintf(w, `{"result":"failed"}`)
		}
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	models.DeleteUser(cast.ToInt(r.FormValue("id")))
	SaveDb()
	http.Redirect(w, r, "/", 302)
}

func ModifyUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tpl := template.Must(template.ParseFiles("template/modify.html"))
		tpl.ExecuteTemplate(w, "modify.html", models.ModifyUserGet(r.FormValue("id")))
	} else {
		fmt.Println(r.FormValue("sex"), "modi")
		if models.ModifyUserPost(
			r.FormValue("id"),
			r.FormValue("name"),
			r.FormValue("sex") == "1",
			r.FormValue("addr"),
		) {
			SaveDb()
			fmt.Fprintf(w, `{"result":"sucess"}`)

		} else {
			fmt.Fprintf(w, `{"result":"failed"}`)
		}
	}
}
func SaveDb() {
	// confirm := utils.Input("请选择是否保存数据?(Y/n)：")
	confirm := "y"
	if strings.ToLower(confirm) == "y" || strings.ToLower(confirm) == "yes" {
		if DbType == "csv" {
			models.DbToCsv()
			// fmt.Println("1")
		} else if DbType == "gob" {
			models.DbToGob()
		} else if DbType == "json" {
			models.DbToJson()
		}

	} else {
		fmt.Println("请继续使用")
		return
	}
}
