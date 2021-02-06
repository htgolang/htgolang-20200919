package config

import (
	"github.com/spf13/viper"
)

type Web struct {
	Addr string `mapstructure:"addr"`
	Auth struct {
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"auth"`
}

type MySQL struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Db       string `mapstructure:"db"`
}

type Logger struct {
	FileName   string `mapstructure:"filename"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	Compress   bool   `mapstructure:"compress"`
}

type Options struct {
	MySQL  MySQL  `mapstructure:"mysql"`
	Web    Web    `mapstructure:"web"`
	Logger Logger `mapstructure:"logger"`
}

func ParseConfig(path string) (*Options, error) {
	conf := viper.New()

	conf.SetDefault("web.addr", ":10001")

	conf.AutomaticEnv()
	conf.SetEnvPrefix("MySQL_EXPORTER")

	conf.SetConfigFile(path)
	if err := conf.ReadInConfig(); err != nil {
		// 允许配置文件不存在，使用默认配置
		// err.(viper.ConfigFileNotFoundError)
		return nil, err
	}

	options := &Options{}

	if err := conf.Unmarshal(options); err != nil {
		return nil, err
	}
	return options, nil
}
