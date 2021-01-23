package collectors

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type ConnectionCollector struct {
	mysqlCollector
	maxConnectionDesc    *prometheus.Desc
	threadsConnectedDesc *prometheus.Desc
}

func NewConnectionCollector(db *sql.DB) *ConnectionCollector {

	return &ConnectionCollector{
		mysqlCollector: mysqlCollector{db},
		maxConnectionDesc: prometheus.NewDesc(
			"mysql_global_variables_max_connections",
			"MySQL global variables max connections",
			nil,
			nil,
		),
		threadsConnectedDesc: prometheus.NewDesc(
			"mysql_global_status_threads_connected",
			"MySQL global status threads connected",
			nil,
			nil,
		),
	}
}

func (c *ConnectionCollector) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.maxConnectionDesc
	descs <- c.threadsConnectedDesc
}

func (c *ConnectionCollector) Collect(metrics chan<- prometheus.Metric) {
	maxConnections := c.variables("max_connections")
	logrus.WithFields(logrus.Fields{
		"metric": "max_connections",
		"sample": maxConnections,
	}).Debug("command metric")

	metrics <- prometheus.MustNewConstMetric(
		c.maxConnectionDesc,
		prometheus.CounterValue,
		maxConnections,
	)

	threadsConnected := c.status("threads_connected")

	logrus.WithFields(logrus.Fields{
		"metric": "max_connections",
		"sample": threadsConnected,
	}).Debug("command metric")

	metrics <- prometheus.MustNewConstMetric(
		c.threadsConnectedDesc,
		prometheus.CounterValue,
		threadsConnected,
	)
}
