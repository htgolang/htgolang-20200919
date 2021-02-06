package models

/*

scrape_configs: # 抓取配置
  - job_name: 'prometheus' #任务 采集目标分类
    basic_auth:
      username: admin
      password: 123123
    static_configs: # 抓取目标静态配置
    - targets:
      - "localhost:9090" #抓取目标

  - job_name: "node"
    static_configs:
    - targets:
      - "localhost:9100"
      - "localhost:9101"

*/

type BasicAuth struct {
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

type StaticConfig struct {
	Targets []string `yaml:"targets" json:"target"`
}

type Job struct {
	JobName       string         `yaml:"job_name" json:"job_name"`
	BasicAuth     BasicAuth      `yaml:"basic_auth" json:"basic_auth"`
	StaticConfigs []StaticConfig `yaml:"static_configs" json:"static_configs"`
}

type Jobs []Job
