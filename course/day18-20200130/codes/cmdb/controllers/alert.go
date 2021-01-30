package controllers

import (
	base "cmdb/base/controllers"
	"cmdb/services"
	"strings"

	"github.com/astaxie/beego/utils/pagination"
)

type AlertController struct {
	base.LayoutController
}

func (c *AlertController) Query() {
	pageSize := 10
	q := strings.TrimSpace(c.GetString("q"))

	total := services.AlertService.Total(q)
	paginator := pagination.SetPaginator(c.Ctx, pageSize, total)

	c.Data["q"] = q
	c.Data["alerts"] = services.AlertService.Query(q, paginator.Offset(), pageSize)
	c.TplName = "alert/query.html"
	c.LayoutSections["LayoutSectionTitle"] = "alert/query_title.html"
}
