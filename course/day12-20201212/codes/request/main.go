package main

import (
	"github.com/astaxie/beego"
)

type LoginForm struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}

type RequestController struct {
	beego.Controller
}

// 用户提交的数据
/*
Context => c.Ctx
	请求数据 c.Ctx
			c.Ctx.Request => http.Requst
				URL
					ParseForm + Form
					FormValue
				BODY
					x-www-form-urlencoded
						ParseForm + Form
						ParseForm + PostForm
						FormValue
						PostFormValue
					其他
						Body
			c.Ctx.Input

Controller =>
	请求数据
		Get*
*/

func (c *RequestController) Header() {
	// // 请求控制器和动作
	// fmt.Println(c.GetControllerAndAction())

	// // 请求头信息
	// // 请求行
	// ctx := c.Ctx
	// input := ctx.Input
	// fmt.Println(input.Method(), input.Protocol(), input.URI(), input.URL())

	// // 请求头信息
	// fmt.Println(input.Header("User-Agent"))

	// // url数据
	// fmt.Println(input.Query("id"))
	// var id int
	// input.Bind(&id, "id")
	// fmt.Println(id)

	// fmt.Println(c.GetInt("id"))
	// // GetBool, GetString, GetStrings
	// fmt.Println(c.Input())

	// var form LoginForm

	// err := c.ParseForm(&form)
	// fmt.Println(err, form)

	// fmt.Println(input.CopyBody(1024 * 1024))
	// // fmt.Println(string(input.RequestBody))
	// file, header, err := c.GetFile("x")
	// fmt.Println(file, header, err)
	c.SaveToFile("x", "test.txt")

	c.Ctx.Output.Body([]byte("header"))
}

func main() {
	beego.AutoRouter(&RequestController{})
	beego.Run()
}
