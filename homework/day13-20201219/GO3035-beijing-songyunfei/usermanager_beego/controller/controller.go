package controller

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"usermanager_beego/users"
)

var Udb users.Mydb
type Usermanager struct {
	beego.Controller
}

func (c *Usermanager) Login()  {
	if user := c.GetSession("User"); user != nil {
		c.Redirect("/usermanager/list",302)
		return
	}
	if c.Ctx.Input.IsGet(){
		c.TplName = "login.html"
	}else {
		username := c.Input().Get("username")
		passwd := c.Input().Get("password")
		if Udb.Auth(username,passwd) {
			c.SetSession("User",username)
			c.Redirect("/usermanager/list",302)
		}else {
			c.Data["form"] = username
			c.Data["errors"] = "用户名或密码错误"
			c.TplName = "login.html"
		}
	}
}


func (c *Usermanager) List()  {
	c.Data["users"] = Udb.Getall()
	c.TplName = "index.html"
}

func (c *Usermanager) Add()  {
	if user := c.GetSession("User"); user == nil {
		c.Redirect("/usermanager/login",302)
		return
	}
	if c.Ctx.Input.IsPost()  {
		un := c.Input().Get("username")
		ua := c.Input().Get("addr")
		us := c.Input().Get("sex")
		ut := c.Input().Get("tel")
		up := c.Input().Get("password")
		ub := c.Input().Get("brth")

		if err := Udb.Add(un, ua,us, ut, up, ub); err != nil {
			c.Ctx.WriteString(fmt.Sprintf("%s",err))
		} else {
			c.Redirect("/",302)
		}

	}else {
		c.TplName = "add.html"
	}
}

func (c *Usermanager) Del()  {
	if user := c.GetSession("User"); user == nil {
		c.Redirect("/usermanager/login",302)
		return
	}
	id:= c.Input().Get("Id")
	if id == "" {
		c.Ctx.WriteString("Id 不能为空")
	}
	if fid, err := strconv.ParseInt(id,10,64);err != nil {
		fmt.Println(err)
		c.Ctx.WriteString("非法ID")
	}else {
		err := Udb.Del(fid)
		if err != nil {
			c.Ctx.WriteString(fmt.Sprintf("%s",err))
		}
		c.Redirect("/",302)
	}
}

func (c *Usermanager) Modify(){
	if user := c.GetSession("User"); user == nil {
		c.Redirect("/usermanager/login",302)
		return
	}
	id := c.Input().Get("Id")
	if id == "" {
		c.Ctx.WriteString("Id 不能为空")
	}
	uid, err := strconv.ParseInt(id,10,64)
	if err != nil {
		fmt.Println(err)
		c.Ctx.WriteString("非法ID")
	}
	uinfo, err := Udb.FindByid(uid)
	if err != nil {
		c.Ctx.WriteString(fmt.Sprintf("%s",err))
	}
	if c.Ctx.Input.IsGet() {
		c.Data["uinfo"] = uinfo
		c.TplName = "modify.html"
	}else {
		id := c.Input().Get("Id")
		name := c.Input().Get("username")
		addr := c.Input().Get("addr")
		tel := c.Input().Get("tel")
		sex:= c.Input().Get("sex")
		up := c.Input().Get("password")
		ub := c.Input().Get("brth")
		uid, err := strconv.ParseInt(id,10,64)
		if err != nil {
			c.Ctx.WriteString(fmt.Sprintf("非法ID:%d\n",id))
		}
		err = Udb.Modify(uid,name,addr,sex,tel,ub,up)
		if err != nil {
			c.Ctx.WriteString(fmt.Sprintf("%s",err))
		}
		c.Redirect("/", 302)
	}
}

func (c *Usermanager) Query() {
	if user := c.GetSession("User"); user == nil {
		c.Redirect("/usermanager/login",302)
		return
	}
	type mydata struct{
		Ud users.Userinfo
		Ok bool
	}
	if c.Ctx.Input.IsPost() {
		str := c.Input().Get("querystr")
		u,ok := Udb.QueryUser(str)
		rdata := mydata{
			Ud: u,
			Ok: ok,
		}
		c.Data["rdata"] = rdata
		c.TplName = "query.html"
	}else {
		rdata := mydata{
			Ud: users.Userinfo{},
			Ok: false,
		}
		c.Data["rdata"] = rdata
		c.TplName = "query.html"
	}
}

func (c *Usermanager) Entrance()  {
	if user := c.GetSession("User"); user == nil {
		c.Redirect("/usermanager/login",302)
		return
	}
	c.Redirect("/usermanager/list",302)
}

func (c *Usermanager) Logout() {
	c.DestroySession()
	c.Redirect("/usermanager/login/", 302)
}