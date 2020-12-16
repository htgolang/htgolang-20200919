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

func (c *Usermanager) List()  {
	c.Data["users"] = Udb.Getall()
	c.TplName = "index.html"
}

func (c *Usermanager) Add()  {
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
	id:= c.Input().Get("Id")
	if id == "" {
		c.Ctx.WriteString("Id 不能为空")
	}
	if fid, err := strconv.Atoi(id);err != nil {
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
	id := c.Input().Get("Id")
	if id == "" {
		c.Ctx.WriteString("Id 不能为空")
	}
	uid, err := strconv.Atoi(id)
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
		err := Udb.Modify(id,name,addr,sex,tel,ub,up)
		if err != nil {
			c.Ctx.WriteString(fmt.Sprintf("%s",err))
		}
		c.Redirect("/", 302)
	}
}

func (c *Usermanager) Query() {
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
	c.Redirect("/usermanager/list",302)
}