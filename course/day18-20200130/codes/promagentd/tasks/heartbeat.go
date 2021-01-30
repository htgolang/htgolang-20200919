package tasks

import (
	"promagentd/client"
	"promagentd/config"
	"time"
)

type HeartbeatTask struct {
	option *config.Option
	client *client.Client
}

func NewHeartbeatTask(option *config.Option, client *client.Client) *HeartbeatTask {
	return &HeartbeatTask{option, client}
}

func (t *HeartbeatTask) Run() {
	ticker := time.NewTicker(10 * time.Second)
	for {
		event := map[string]interface{}{
			"uuid": t.option.UUID,
		}
		t.client.Heartbeat(event)
		<-ticker.C
	}
}
