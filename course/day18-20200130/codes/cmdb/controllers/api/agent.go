package api

import (
	base "cmdb/base/controllers"
	"encoding/json"

	"cmdb/forms"
	"cmdb/services"
)

type AgentController struct {
	base.ApiController
}

func (c *AgentController) Register() {
	c.Ctx.Input.CopyBody(1024 * 1024)
	var form forms.RegisterForm
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err == nil {
		services.AgentService.Register(form.ToModel())
	}

	c.Data["json"] = map[string]string{
		"code": "200",
	}
	c.ServeJSON()
}

func (c *AgentController) Heartbeat() {
	uuid := c.GetString("uuid")
	services.AgentService.Heartbeat(uuid)
	c.Data["json"] = map[string]string{
		"code": "200",
	}
	c.ServeJSON()
}

func (c *AgentController) Config() {
	uuid := c.GetString("uuid")
	version, _ := c.GetInt64("version")
	// Todolist
	config, version := services.AgentService.GetConfig(uuid, version)
	c.Data["json"] = map[string]interface{}{
		"code":    "200",
		"config":  config,
		"version": version,
	}
	c.ServeJSON()
}
