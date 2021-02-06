package controllers

import (
	base "cmdb/base/controllers"
	"cmdb/forms"
	"cmdb/services"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type AgentController struct {
	base.LayoutController
}

func (c *AgentController) Query() {
	c.Data["agents"] = services.AgentService.Query()
	c.TplName = "agent/query.html"
	c.LayoutSections["LayoutSectionTitle"] = "agent/query_title.html"
}

func (c *AgentController) Modify() {
	form := &forms.ModifyAgentForm{}
	valid := &validation.Validation{}

	if c.Ctx.Input.IsPost() {
		c.ParseForm(form)
		if success, err := valid.Valid(form); err == nil && success {
			services.AgentService.Modify(form.ToModel())
			c.Redirect(beego.URLFor("AgentController.Query"), http.StatusFound)
		}
	} else {
		pk, _ := c.GetInt64("pk")
		agent := services.AgentService.GetByPk(pk)
		form.FromModel(agent)
	}

	c.Data["form"] = form
	c.Data["errors"] = valid.ErrorMap()
	c.TplName = "agent/modify.html"
	c.LayoutSections["LayoutSectionTitle"] = "agent/modify_title.html"
	c.LayoutSections["LayoutSectionStyle"] = "agent/modify_style.html"
	c.LayoutSections["LayoutSectionScript"] = "agent/modify_script.html"
}
