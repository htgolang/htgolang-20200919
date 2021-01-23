package api

import (
	"cmdb/base/controllers"
	"cmdb/forms"
	"encoding/json"
	"fmt"
)

type PrometheusController struct {
	controllers.ApiController
}

func (c *PrometheusController) Alert() {
	c.Ctx.Input.CopyBody(1024 * 1024)
	// c.Ctx.Input.RequestBody => 结构体
	// 定义表
	// 存储
	fmt.Println("alert")
	fmt.Println(string(c.Ctx.Input.RequestBody))
	form := forms.AlertForm{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err == nil {
		for _, alert := range form.Alerts {
			fmt.Printf("%#v\n", alert)
		}
	} else {
		fmt.Println(err)
	}

	c.Data["json"] = map[string]string{"code": "200"}
	c.ServeJSON()
}
