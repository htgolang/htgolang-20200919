package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// 路由控制器
type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	// 没有指定响应 views/homecontroller/get.tpl
	c.Ctx.Output.Body([]byte("HomeController.Get"))
}

func (c *HomeController) Post() {
	c.Ctx.Output.Body([]byte("HomeController.Post"))
}

type UserController struct {
	beego.Controller
}

// /user/?:id/
func (c *UserController) Get() {
	c.Ctx.Output.Body([]byte(c.Ctx.Input.Param(":id")))
}

func (c *UserController) Create() {
	// Post
	c.Ctx.Output.Body([]byte("Create"))
}

func (c *UserController) Modify() {
	// Put
	c.Ctx.Output.Body([]byte("Modify"))
}

func (c *UserController) Delete() {
	// Delete
	c.Ctx.Output.Body([]byte("Delete"))
}

func (c *UserController) Any() {
	// Other
	c.Ctx.Output.Body([]byte("Any"))
}

// /Controller/CMethod
// Auth/Login
// Auth/Logout
type AuthController struct {
	beego.Controller
}

func (c *AuthController) Login() {
	c.Ctx.Output.Body([]byte("Login"))
}

func (c *AuthController) Logout() {
	c.Ctx.Output.Body([]byte("Logout"))

}

func main() {
	// 2.0.0 beta
	// 1.12.*
	// 定义处理器
	// 绑定URL和处理器的功能 路由
	// 基本路由
	// 为 URL 绑定某个请求方法的 路由函数
	// 为根路径的GET方法绑定函数
	// Post, Delete, Head, PUT, OPTIONS, Patch
	beego.Get("/", func(ctx *context.Context) {
		//
		ctx.Output.Body([]byte("hi, beego"))
	})

	beego.Post("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("post"))
	})

	beego.Delete("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("delete"))
	})

	beego.Any("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("any"))
	})

	// 提交参数
	// ?a=b
	// body
	// delete/1/
	// 正则路由 /delete/数值/ => 数值解析到参数中
	// 正则
	beego.Any(`/delete/:id(\d{1,8})/`, func(ctx *context.Context) {
		ctx.Output.Body([]byte("delete:" + ctx.Input.Param(":id")))
	})

	beego.Any(`/get/:id:int/`, func(ctx *context.Context) {
		ctx.Output.Body([]byte("get:" + ctx.Input.Param(":id")))
	})

	beego.Any(`/user/*`, func(ctx *context.Context) {
		ctx.Output.Body([]byte("user:" + ctx.Input.Param(":splat")))
	})

	beego.Any(`/file/*.*`, func(ctx *context.Context) {
		ctx.Output.Body([]byte("file:" + ctx.Input.Param(":path") + "," + ctx.Input.Param(":ext")))
	})

	// 路由控制器
	beego.Router("/home/", &HomeController{})
	beego.Router(`/muser/?:id(\d+)/`, &UserController{}, "POST:Create;PUT:Modify;GET:Get;DELETE:Delete;*:Any")

	// 自动路由
	beego.AutoRouter(&AuthController{})

	// 注解路由

	// 启动服务
	beego.Run()
}
