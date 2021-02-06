package domain

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
