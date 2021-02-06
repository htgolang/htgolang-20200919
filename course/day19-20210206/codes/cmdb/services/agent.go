package services

import (
	base "cmdb/base/models"
	"cmdb/models"
	"time"

	"github.com/astaxie/beego/orm"
	"gopkg.in/yaml.v2"
)

type agentService struct {
}

var AgentService = new(agentService)

func (s *agentService) Register(agent *models.Agent) {
	now := time.Now()
	ormer := orm.NewOrm()
	tmpAgent := models.Agent{
		UUID: agent.UUID,
	}
	if err := ormer.Read(&tmpAgent, "UUID"); err == nil {
		// 存在
		tmpAgent.Hostname = agent.Hostname
		tmpAgent.Addr = agent.Addr
		tmpAgent.Heartbeat = &now
		tmpAgent.DeletedAt = nil
		ormer.Update(&tmpAgent, "Hostname", "Addr", "Heartbeat", "DeletedAt")
	} else if err == orm.ErrNoRows {
		// 不存在
		agent.Heartbeat = &now
		ormer.Insert(agent)
	}
}

func (s *agentService) Heartbeat(uuid string) {
	now := time.Now()
	tmpAgent := models.Agent{
		UUID: uuid,
	}

	ormer := orm.NewOrm()
	if err := ormer.Read(&tmpAgent, "UUID"); err == nil {
		tmpAgent.Heartbeat = &now
		tmpAgent.DeletedAt = nil
		ormer.Update(&tmpAgent, "Heartbeat", "DeletedAt")
	}
}

func (s *agentService) GetConfig(uuid string, version int64) (base.Jobs, int64) {
	ormer := orm.NewOrm()
	tmpAgent := models.Agent{UUID: uuid}
	var jobs base.Jobs
	var configVersion int64
	if err := ormer.Read(&tmpAgent, "UUID"); err == nil {
		if tmpAgent.ConfigVersion > version {
			yaml.Unmarshal([]byte(tmpAgent.Config), &jobs)
			configVersion = tmpAgent.ConfigVersion
		}
	}
	return jobs, configVersion
}

func (s *agentService) Query() []*models.Agent {
	var agents []*models.Agent
	queryset := orm.NewOrm().QueryTable(models.Agent{})
	queryset.Filter("DeletedAt__isnull", true).All(&agents)
	return agents
}

func (s *agentService) GetByPk(pk int64) *models.Agent {
	ormer := orm.NewOrm()
	tmpAgent := models.Agent{ID: pk}
	if err := ormer.Read(&tmpAgent); err == nil {
		return &tmpAgent
	}
	return nil
}

func (s *agentService) Modify(agent *models.Agent) {
	ormer := orm.NewOrm()
	tmpAgent := models.Agent{ID: agent.ID}
	if err := ormer.Read(&tmpAgent); err == nil {
		if tmpAgent.Config != agent.Config {
			tmpAgent.ConfigVersion += 1
		}
		tmpAgent.Description = agent.Description
		tmpAgent.Config = agent.Config
		ormer.Update(&tmpAgent, "Description", "Config", "ConfigVersion")
	}
}
