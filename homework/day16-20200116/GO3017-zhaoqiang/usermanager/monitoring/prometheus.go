package monitoring

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestTime prometheus.Summary = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Name:        "request_response_time",
			Help:        "Request response time Summary",
			Objectives:  map[float64]float64{0.95: 0.1, 0.90: 0.1, 0.80: 0.1, 0.5: 0.4},
			ConstLabels: prometheus.Labels{"name": "userManager"},
		},
	)
	RequestCountTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name:        "request_counter_total",
			Help:        "request counter total Counter",
			ConstLabels: prometheus.Labels{"name": "userManager"},
		},
	)
	RequestPathCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:        "request_path_counter",
			Help:        "request path number counter",
			ConstLabels: prometheus.Labels{"name": "userManager"},
		},
		[]string{"URL"},
	)
	RequestStatusCode = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:        "request_status_code",
			Help:        "request status code number counter",
			ConstLabels: prometheus.Labels{"name": "userMnager"},
		},
		[]string{"Code"},
	)
	GorouterNumber = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name:        "goroutiner_Number",
			Help:        "goroutine number gauge",
			ConstLabels: prometheus.Labels{"name": "userManager"},
		},
	)
)

func init() {
	prometheus.MustRegister(RequestTime)
	prometheus.MustRegister(RequestCountTotal)
	prometheus.MustRegister(RequestPathCounter)
	prometheus.MustRegister(RequestStatusCode)
	prometheus.MustRegister(GorouterNumber)
}
