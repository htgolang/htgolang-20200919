package collectors

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type SlowQueriesCollector struct {
	mysqlCollector
	desc *prometheus.Desc
}

func NewSlowQueriesCollector(db *sql.DB) *SlowQueriesCollector {
	return &SlowQueriesCollector{
		mysqlCollector: mysqlCollector{db},
		desc: prometheus.NewDesc(
			"mysql_global_status_slow_queries",
			"MySQL global status slow queries",
			nil,
			nil,
		),
	}
}

func (c *SlowQueriesCollector) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.desc
}

func (c *SlowQueriesCollector) Collect(metrics chan<- prometheus.Metric) {
	sample := c.status("slow_queries")
	logrus.WithFields(logrus.Fields{
		"metric": "slow_queries",
		"sample": sample,
	}).Debug("command metric")

	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, sample)
}

type QpsCollector struct {
	mysqlCollector
	desc *prometheus.Desc
}

func NewQpsCollector(db *sql.DB) *QpsCollector {
	return &QpsCollector{
		mysqlCollector: mysqlCollector{db},
		desc: prometheus.NewDesc(
			"mysql_global_status_queries",
			"MySQL global status Queries",
			nil,
			nil,
		),
	}
}

func (c *QpsCollector) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.desc
}

func (c *QpsCollector) Collect(metrics chan<- prometheus.Metric) {
	sample := c.status("queries")
	logrus.WithFields(logrus.Fields{
		"metric": "qps",
		"sample": sample,
	}).Debug("command metric")
	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, sample)
}
