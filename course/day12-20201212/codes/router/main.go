package main

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {
	c.Ctx.WriteString("index")
}

func main() {
	// url 控制器关系
	// url =>控制器, 请求方式-> 控制器函数
	// beego.Router("/home/", &HomeController{}, "*:Index")
	// url第一个参数 => 控制器 url第二个参数 => 控制器函数
	beego.AutoRouter(&HomeController{})
	// beego.Router("/home/index", &HomeController{}, "*:Index")

	beego.Run()
}
