package collectors

import (
	"database/sql"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type CommandCollector struct {
	mysqlCollector
	desc *prometheus.Desc
}

func NewCommandCollector(db *sql.DB) *CommandCollector {
	return &CommandCollector{
		mysqlCollector: mysqlCollector{db},
		desc: prometheus.NewDesc(
			"mysql_global_status_command",
			"MySQL global status command",
			[]string{"command"},
			nil,
		),
	}
}

func (c *CommandCollector) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.desc
}

func (c *CommandCollector) Collect(metrics chan<- prometheus.Metric) {
	names := []string{
		"insert",
		"update",
		"delete",
		"select",
	}

	for _, name := range names {
		sample := c.status(fmt.Sprintf("Com_%s", name))

		logrus.WithFields(logrus.Fields{
			"metric": name,
			"sample": sample,
		}).Debug("command metric")

		metrics <- prometheus.MustNewConstMetric(
			c.desc,
			prometheus.CounterValue,
			sample,
			name,
		)
	}
}
