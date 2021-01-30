package controllers

import (
	"strings"

	"github.com/astaxie/beego"
)

type LayoutController struct {
	AuthorizationController
}

func (c *LayoutController) Prepare() {
	c.AuthorizationController.Prepare()

	if c.Data["currentUser"] == nil {
		return
	}

	beego.ReadFromRequest(&c.Controller)

	controllerName, _ := c.GetControllerAndAction()

	controllerName = strings.ToLower(controllerName)
	controllerName = strings.ReplaceAll(controllerName, "controller", "")

	c.Layout = "base/layout.html"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["LayoutSectionTitle"] = ""
	c.LayoutSections["LayoutSectionStyle"] = ""
	c.LayoutSections["LayoutSectionFooter"] = ""
	c.LayoutSections["LayoutSectionScript"] = ""

	c.Data["navKey"] = controllerName
}
