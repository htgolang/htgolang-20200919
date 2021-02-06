package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func ReadFile(path string) string {
	if ctx, err := ioutil.ReadFile(path); err != nil {
		return ""
	} else {
		return string(ctx)
	}
}

func WriteFile(path string, txt string) {
	ioutil.WriteFile(path, []byte(txt), os.ModePerm)
}

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

func main() {
	txt := ReadFile("prometheus.yml")
	var config struct {
		Global        interface{} `yaml:"global"`
		Alerting      interface{} `yaml:"alerting"`
		RuleFiles     interface{} `yaml:"rule_files"`
		ScrapeConfigs Jobs        `yaml:"scrape_configs"`
	}
	yaml.Unmarshal([]byte(txt), &config)

	fmt.Println(config)

	jobs := Jobs{}
	jobs = append(jobs, Job{
		JobName: "aaaaaaa",
	})
	jobs = append(jobs, Job{
		JobName: "bbb",
	})
	config.ScrapeConfigs = jobs
	content, _ := yaml.Marshal(&config)
	WriteFile("prometheus.yml2", string(content))
}
