package forms

import (
	"fmt"

	"gopkg.in/yaml.v2"

	base "cmdb/base/models"
	"cmdb/models"

	"github.com/astaxie/beego/validation"
)

type RegisterForm struct {
	UUID     string `json:"uuid"`
	Addr     string `json:"addr"`
	Hostname string `json:"hostname"`
}

func (f *RegisterForm) ToModel() *models.Agent {
	return &models.Agent{
		UUID:     f.UUID,
		Addr:     f.Addr,
		Hostname: f.Hostname,
	}
}

type ModifyAgentForm struct {
	ID          int64  `form:"id"`
	Description string `form:"description"`
	Config      string `form:"config"`
}

func (f *ModifyAgentForm) ToModel() *models.Agent {
	return &models.Agent{
		ID:          f.ID,
		Description: f.Description,
		Config:      f.Config,
	}

}

func (f *ModifyAgentForm) FromModel(agent *models.Agent) {
	if agent == nil {
		return
	}
	f.ID = agent.ID
	f.Description = agent.Description
	f.Config = agent.Config
}

func (f *ModifyAgentForm) Valid(v *validation.Validation) {
	var jobs base.Jobs
	if err := yaml.UnmarshalStrict([]byte(f.Config), &jobs); err != nil {
		v.SetError("config", err.Error())
	}
	fmt.Println(jobs)
}
