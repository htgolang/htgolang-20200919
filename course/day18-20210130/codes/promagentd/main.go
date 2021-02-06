package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"promagentd/client"
	"promagentd/config"
	"promagentd/tasks"

	"github.com/imroc/req"
	"github.com/sirupsen/logrus"
)

func main() {
	// 1. 配置
	var (
		server  = ""
		conf    = ""
		help, h bool
	)
	flag.StringVar(&server, "server", "http://10.0.0.1:8888", "cmdb server")
	flag.StringVar(&conf, "conf", "/opt/prometheus/prometheus/prometheus.yml", "cmdb server")
	flag.BoolVar(&help, "help", false, "help")
	flag.BoolVar(&h, "h", false, "help")
	flag.Usage = func() {
		fmt.Println("usage: promagentd --server http://localhost:9999")
		flag.PrintDefaults()
	}

	flag.Parse()

	if h || help {
		flag.Usage()
		os.Exit(0)
	}
	req.Debug = false
	option, err := config.NewOption(server, conf)
	if err != nil {
		logrus.Fatal(err)
	}

	client := client.NewClient(option)

	configTask := tasks.NewConfigTask(option, client)
	heartbeatTask := tasks.NewHeartbeatTask(option, client)
	registerTask := tasks.NewRegisterTask(option, client)

	// 3. 通信(http)
	// 心跳
	go configTask.Run()
	go heartbeatTask.Run()
	go registerTask.Run()
	// 注册
	// 更新配置 yaml
	// 应用配置
	//    a. promtool check
	//    b. 覆盖文件
	//    c. 热加载
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt
}
