package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Options struct {
	MySQL struct {
		Host     string
		Password string
	} `mapstructure:"db"`
}

func main() {

	viper.SetDefault("mysql.port", 3306)

	viper.AutomaticEnv()
	viper.SetEnvPrefix("mysql_exporter")

	viper.SetConfigType("yaml")
	viper.SetConfigName("test")
	viper.AddConfigPath(".")

	// viper.SetConfigFile("./test.yaml")

	err := viper.ReadInConfig()

	fmt.Println(err)
	fmt.Println(viper.Get("mysql"))
	fmt.Println(viper.GetString("mysql.host"))
	fmt.Println(viper.GetInt("mysql.port"))
	fmt.Println(viper.GetString("mysql.username"))
	fmt.Println(viper.GetString("mysql.password"))
	fmt.Println(viper.GetString("mysql.db"))

	fmt.Println(viper.GetString("redis.host"))

	options := &Options{}
	err = viper.Unmarshal(&options)

	fmt.Println(err, options)

	viper.SetDefault("redis.port", 6379)
	viper.Set("redis.host", "1.1.1.1")
	viper.WriteConfigAs("./test2.yaml")
}
