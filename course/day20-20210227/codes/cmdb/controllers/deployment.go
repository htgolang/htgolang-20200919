package controllers

import (
	base "cmdb/base/controllers"
	"cmdb/forms"
	"cmdb/services"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type DeploymentController struct {
	base.LayoutController
}

func (c *DeploymentController) Add() {
	var form forms.AddDeploymentForm
	valid := validation.Validation{}

	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(&form); err == nil {
			services.DeploymentService.Create(form.Namespace, form.Name, form.Image, form.Labels(), form.Exposes(), form.Replicas)
		}
		c.Redirect(beego.URLFor("DeploymentController.Query"), 302)
	}
	c.Data["errors"] = valid.ErrorMap()
	c.Data["form"] = &form
	c.Data["namespaces"] = services.NamespaceService.Query()

	c.TplName = "deployment/add.html" // LayoutContent内容替代的文件路径
	c.Layout = "base/layout.html"
	c.LayoutSections["LayoutSectionTitle"] = "deployment/add_title.html"
}

func (c *DeploymentController) Modify() {
	var form forms.ModifyDeploymentForm
	valid := validation.Validation{}

	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(&form); err == nil {
			services.DeploymentService.Modify(form.Namespace, form.Name, form.Image, form.Exposes(), form.Replicas)
		}
		c.Redirect(beego.URLFor("DeploymentController.Query"), 302)
	} else {
		deployment := services.DeploymentService.Get(c.GetString("namespace"), c.GetString("name"))
		form.FromModel(deployment)
	}
	c.Data["errors"] = valid.ErrorMap()
	c.Data["form"] = &form

	c.TplName = "deployment/modify.html" // LayoutContent内容替代的文件路径
	c.Layout = "base/layout.html"
	c.LayoutSections["LayoutSectionTitle"] = "deployment/modify_title.html"
}

func (c *DeploymentController) Query() {
	c.Data["objects"] = services.DeploymentService.Query()

	c.TplName = "deployment/query.html"
	c.LayoutSections["LayoutSectionTitle"] = "deployment/query_title.html"
}

func (c *DeploymentController) Delete() {
	name := c.GetString("name")
	namespace := c.GetString("namespace")
	services.DeploymentService.Delete(name, namespace)
	c.Redirect(beego.URLFor("DeploymentController.Query"), 302)
}
