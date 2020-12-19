package controllers

import (
	"fmt"
	base "user/base/controllers"
)

type HomeController struct {
	base.RequiredAuthController
}

func (c *HomeController) Index() {
	fmt.Println("home")
	c.Ctx.WriteString("home")
}
