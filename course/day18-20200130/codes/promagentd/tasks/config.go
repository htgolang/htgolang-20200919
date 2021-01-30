package tasks

import (
	"encoding/json"
	"fmt"
	"os"
	"promagentd/client"
	"promagentd/config"
	"promagentd/domain"
	"promagentd/utils"
	"time"

	"gopkg.in/yaml.v2"
)

type ConfigTask struct {
	option *config.Option
	client *client.Client
}

func NewConfigTask(option *config.Option, client *client.Client) *ConfigTask {
	return &ConfigTask{option, client}
}

func (t *ConfigTask) Run() {
	ticker := time.NewTicker(time.Minute)
	var version int64
	for {
		event := map[string]interface{}{
			"uuid":    t.option.UUID,
			"version": version,
		}

		if txt, err := t.client.Config(event); err == nil {
			var result struct {
				Code    string
				Config  domain.Jobs
				Version int64
			}
			err := json.Unmarshal([]byte(txt), &result)
			if err == nil {
				version = result.Version
				if len(result.Config) > 0 {
					// 更改配置
					t.handle(result.Config)
				}
			}
		}
		<-ticker.C
	}

}

func (t *ConfigTask) handle(jobs domain.Jobs) {
	fmt.Println(jobs)
	txt := utils.ReadFile(t.option.Conf)

	var conf struct {
		Global       interface{} `yaml:"global"`
		Alerting     interface{} `yaml:"alerting"`
		RuleFiles    interface{} `yaml:"rule_files"`
		ScrapeConfig domain.Jobs `yaml:"scrape_configs"`
	}
	err := yaml.Unmarshal([]byte(txt), &conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	conf.ScrapeConfig = jobs
	content, err := yaml.Marshal(&conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	path := t.option.Conf + ".bak"
	utils.WriteFile(path, string(content))

	cmd := fmt.Sprintf("%s check config %s", t.option.Promtool, path)
	if output, err := utils.Run(cmd); err == nil {
		os.Rename(t.option.Conf, fmt.Sprintf("%s.%d", t.option.Conf, time.Now().Unix()))
		os.Rename(path, t.option.Conf)

		fmt.Println(utils.Run("systemctl reload prometheus"))
	} else {
		fmt.Println(output, err)
	}
}
