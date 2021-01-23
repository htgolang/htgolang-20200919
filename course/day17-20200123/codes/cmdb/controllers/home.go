package controllers

import (
	base "cmdb/base/controllers"
	"fmt"
)

type HomeController struct {
	base.RequiredAuthController
}

func (c *HomeController) Index() {
	fmt.Println("home")
	c.Ctx.WriteString("home")
}
