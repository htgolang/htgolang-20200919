package tasks

import (
	"promagentd/client"
	"promagentd/config"
	"time"
)

type RegisterTask struct {
	option *config.Option
	client *client.Client
}

func NewRegisterTask(option *config.Option, client *client.Client) *RegisterTask {
	return &RegisterTask{option, client}
}

func (t *RegisterTask) Run() {
	ticker := time.NewTicker(time.Hour)
	for {
		event := map[string]interface{}{
			"uuid":     t.option.UUID,
			"addr":     t.option.Addr,
			"hostname": t.option.Hostname,
		}
		t.client.Register(event)
		<-ticker.C
	}
}
