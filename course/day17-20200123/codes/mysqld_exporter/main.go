package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"mysqld_exporter/collectors"
	"mysqld_exporter/config"
	"mysqld_exporter/handler"

	_ "github.com/go-sql-driver/mysql"
)

func initLogger(options config.Logger) func() {
	logger := lumberjack.Logger{
		Filename:   options.FileName,
		MaxSize:    options.MaxSize,
		MaxAge:     options.MaxSize, //days
		MaxBackups: options.MaxBackups,
		Compress:   options.Compress,
	}

	logrus.SetOutput(&logger)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)

	return func() {
		logger.Close()
	}
}

func initDb(options config.MySQL) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		options.Username,
		options.Password,
		options.Host,
		options.Port,
		options.Db,
	)

	return sql.Open("mysql", dsn)
}

func initMetrics(options *config.Options, db *sql.DB) {
	// a. 定义指标
	// mysql_up
	prometheus.MustRegister(prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "mysql_up",
			Help: "MySQL UP Info",
			ConstLabels: prometheus.Labels{
				"addr": net.JoinHostPort(
					options.MySQL.Host,
					strconv.Itoa(options.MySQL.Port),
				),
			},
		},
		func() float64 {
			if err := db.Ping(); err == nil {
				return 1
			} else {
				logrus.WithFields(logrus.Fields{
					"metric": "mysql_up",
				}).Error(err)
			}
			return 0
		},
	))

	prometheus.MustRegister(collectors.NewSlowQueriesCollector(db))
	prometheus.MustRegister(collectors.NewQpsCollector(db))
	prometheus.MustRegister(collectors.NewCommandCollector(db))
	prometheus.MustRegister(collectors.NewConnectionCollector(db))
	prometheus.MustRegister(collectors.NewTrafficCollector(db))

}

func main() {
	// 配置文件通过命令行参数指定
	// flag
	options, err := config.ParseConfig("./etc/exporter.yaml")
	if err != nil {
		logrus.Error(err)
	}

	// 指标项
	// 指标类型/触发时间
	//

	// 0. 存活状态 Gauge 0 1
	// 监控的mysql地址信息
	// 带标签的 固定

	// 1. 慢查询数量 Counter
	// 不带label

	// 2. 执行查询数量 QPS Counter
	// 不带label

	// 3. 执行操作数量 TPS Counter
	// 		SELECT, INSERT, DELETE, UPDATE
	// 带label
	// 4. 流量 Counter
	// 带label

	// 5. 连接数量 Gauge
	// max_conections
	// threads_connected
	// 定义两种指标

	// 采集API时触发

	close := initLogger(options.Logger)
	defer close()

	db, err := initDb(options.MySQL)
	if err != nil && db.Ping() != nil {
		logrus.Fatal(err)
	}

	initMetrics(options, db)

	// b. 暴露指标
	http.Handle("/metrics", handler.Auth(
		promhttp.Handler(),
		handler.AuthSecrets{
			options.Web.Auth.Username: options.Web.Auth.Password,
		},
	))
	logrus.Error(http.ListenAndServe(options.Web.Addr, nil))
}
