package controller

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/prometheus/common/log"
	"strings"
	"usermanager_beego/users"
)

type Authcontroller struct {
	beego.Controller
}


func (c *Authcontroller) GenNav()  {
	//Gen navlist
	var Nav map[string][]users.Permission
	Nav = make(map[string][]users.Permission)
	supadmin := []users.Permission{
		{
			Id:    0,
			PName: "添加用户",
			Path:  "/usermanager/add",
		},
		{
			Id:    0,
			PName: "查找用户",
			Path:  "/usermanager/query",
		},
		{
			Id:    0,
			PName: "角色管理",
			Path:  "/usermanager/role",
		},
		{
			Id:    0,
			PName: "报警管理",
			Path:  "/alert/query",
		},
	}
	admin := []users.Permission{
		{
			Id:    0,
			PName: "添加用户",
			Path:  "/usermanager/add",
		},
		{
			Id:    0,
			PName: "查找用户",
			Path:  "/usermanager/query",
		},
		{
			Id:    0,
			PName: "报警管理",
			Path:  "/alert/query",
		},
	}
	operator := []users.Permission{
		{
			Id:    0,
			PName: "日志管理",
			Path:  "/loganalysis/upload",
		},
	}
	Nav["1"] = append(Nav["1"], supadmin...)
	Nav["2"] = append(Nav["2"], admin...)
	Nav["3"] = append(Nav["3"], operator...)
	roleid := fmt.Sprintf("%d",c.GetSession("roleid"))
	_,ok := Nav[roleid]
	if ok {
		c.Data["navlist"] = Nav[roleid]
	}else {
		log.Error("未获取nav")
	}

}

func (c *Authcontroller) Login()  {
	if user := c.GetSession("User"); user != nil {
		c.Data["cuser"] = user
		c.Redirect("/",302)
		return
	}
	if c.Ctx.Input.IsGet(){
		beego.ReadFromRequest(&c.Controller)
		c.TplName = "login.html"
		return
	}else {
		username := c.Input().Get("username")
		passwd := c.Input().Get("password")
		if Udb.Auth(username,passwd) {
			c.SetSession("User",username)
			c.Data["cuser"] = username
			c.Redirect("/",302)
			return
		}else {
			flash := beego.NewFlash()
			flash.Set("error", "用户名或密码错误")
			flash.Store(&c.Controller)
			c.Data["form"] = username
			c.TplName = "login.html"
			c.Redirect("/authcontroller/login/",302)
			return
		}
	}
}

func (c *Authcontroller) Logout() {
	c.DestroySession()
	c.Redirect("/authcontroller/login/", 302)
	return
}

func (c *Authcontroller) Islogin()  {
	if user := c.GetSession("User"); user != nil {
		c.Data["cuser"] = user
		return
	}else {
		c.Redirect("/authcontroller/login/", 302)
		return
	}
}

func (c *Authcontroller) CheckPermission()  {
	c.Islogin()
	user := c.GetSession("User")
	roleid := GetRoleId(fmt.Sprintf("%s",user))
	userPermission,err := Querypermission(roleid)
	if err != nil {
		log.Error(err)
		c.Abort("403")
		return
	}
	hasper := false
	for _,v := range userPermission {
		requrl := strings.Split(c.Ctx.Request.RequestURI,"?")
		if requrl[0] == v.Path {
			hasper = true
			break
		}
	}
	if hasper {
		c.SetSession("roleid",roleid)
	}else {
		c.DestroySession()
		c.Abort("403")
	}
	return
	
}

func GetRoleId(username string) int {
	user := &users.Userinfo{Name:username}
	qs := orm.NewOrm()
	err := qs.Read(user,"name")
	if err != nil {
		return 0
	}
	role := &users.UserRole{UserId:user.Id}
	err = qs.Read(role,"user_id")
	if err != nil {
		return 0
	}
	return role.RoleId
}

func Querypermission(roleid int) (userPermission []users.Permission,err error) {
	var perSli []users.RolePermission
	qs := orm.NewOrm()
	_,err = qs.QueryTable("role_permission").Filter("RoleId__exact",roleid).All(&perSli)
	if err != nil {
		return nil,err
	}
	for _,v := range perSli {
		var onePer users.Permission
		onePer.Id = v.PermissionId
		err := qs.Read(&onePer)
		if err != nil {
			return nil,err
		}
		userPermission = append(userPermission,onePer)
	}
	return userPermission,nil
}

func QueryRole() (roleall []users.RoleInfo) {
	qs := orm.NewOrm()
	_,err := qs.QueryTable("RoleInfo").All(&roleall)
	if err != nil {
		fmt.Println(err)
	}
	return roleall
}