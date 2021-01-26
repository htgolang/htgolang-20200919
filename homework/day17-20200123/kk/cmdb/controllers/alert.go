package controllers

import (
	base "cmdb/base/controllers"
	"cmdb/services"
	"strings"
)

type AlertController struct {
	base.LayoutController
}

func (c *AlertController) Query() {
	q := strings.TrimSpace(c.GetString("q"))

	c.Data["q"] = q
	c.Data["alerts"] = services.AlertService.Query(q)
	c.TplName = "alert/query.html"
	c.LayoutSections["LayoutSectionTitle"] = "alert/query_title.html"
}
