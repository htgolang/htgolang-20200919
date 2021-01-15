package controller

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/prometheus/common/log"
	"strconv"
	"usermanager_beego/users"
)

var Udb users.Mydb


type Usermanager struct {
	Authcontroller
}


func (c *Usermanager) Prepare() {
	c.Islogin()
	c.CheckPermission()
	c.GenNav()
	beego.ReadFromRequest(&c.Controller)

}


func (c *Usermanager) List()  {
	c.Data["users"] = Udb.Getall()
	bt := fmt.Sprintf("%d",c.GetSession("roleid"))
	rid,err := strconv.Atoi(bt)
	if err != nil {
		fmt.Println(err)
		return
	}
	if rid >= 2 {
		c.Data["isdisable"] = true
	}
	c.Layout = "base/layout.html"
	c.TplName = "list.html"
}

func (c *Usermanager) Add()  {

	if c.Ctx.Input.IsPost()  {
		un := c.Input().Get("username")
		ua := c.Input().Get("addr")
		us := c.Input().Get("sex")
		ut := c.Input().Get("tel")
		up := c.Input().Get("password")
		ub := c.Input().Get("brth")
		ur := c.Input().Get("RoleId")
		if err := Udb.Add(un, ua,us, ut, up, ub,ur); err != nil {
			c.Ctx.WriteString(fmt.Sprintf("%s",err))
		} else {
			c.Redirect("/",302)
		}

	}else {
		rd,err := Udb.GetRole()
		if err != nil {
			fmt.Println(err)
			return
		}
		c.Data["Role"] = rd
		c.TplName = "add.html"
	}
}

func (c *Usermanager) Del()  {

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
			c.Ctx.WriteString(fmt.Sprintf("非法ID:%s\n",id))
		}
		err = Udb.Modify(uid,name,addr,sex,tel,ub,up)
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
		if str == "" {
			flash := beego.NewFlash()
			flash.Set("warn", "关键字不能为空")
			flash.Store(&c.Controller)
			c.Redirect(beego.URLFor("Usermanager.Query"),302)
			return
		}
		u,ok := Udb.QueryUser(str)
		fmt.Println(u,ok)
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
	roleid,ok := c.GetSession("roleid").(int)
	if !ok {
		log.Error("非法roleid:",roleid)
	}
	if roleid >= 3 {
		c.Redirect(beego.URLFor("LogAnalysis.Upload"),302)
	}
	c.Redirect(beego.URLFor("Usermanager.List"),302)
}

func (c *Usermanager) Role()  {
	if c.Ctx.Input.IsGet(){
		c.Data["roleall"] = QueryRole()
		c.Layout = "base/layout.html"
		c.TplName = "role.html"
		roleid := c.Input().Get("id")
		if c.Input().Get("id") != "" {
			rid,err := strconv.Atoi(roleid)
			if err != nil {
				c.Ctx.WriteString("非法roleID")
			}
			perSli,err := Querypermission(rid)
			c.Data["roleper"] = perSli
		}
	}

}