package controller

import (
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/prometheus/client_golang/prometheus"
	"runtime"
	"strconv"
	"time"
)

var (
	Subsystem       string
	Namespace		string
	Name 			string
	Pathcounter *prometheus.CounterVec
	StatusCodecounter *prometheus.CounterVec
	RequestTimeProc *prometheus.HistogramVec
)


func InitMertrics()  {
	Pathcounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        Name,
			Help:        "Request_Path_Count",
			ConstLabels: prometheus.Labels{"name": "request"},
		},
		[]string{"path"},
	)
	StatusCodecounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   Namespace+"Code",
			Subsystem:   Subsystem,
			Name:        Name,
			Help:        "Status_Code_Count",
			ConstLabels: prometheus.Labels{"name": "status_code"},
		},
		[]string{"code"},
	)
	RequestTimeProc = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace:   Namespace+"_time",
			Subsystem:   Subsystem,
			Name:        Name,
			Help:        "Request_Time_ms",
			ConstLabels: prometheus.Labels{"name":"request_time"},
			Buckets:     prometheus.LinearBuckets(0,10,3),
		},
		[]string{"request_path"},
	)

	prometheus.MustRegister(Pathcounter)
	prometheus.MustRegister(StatusCodecounter)
	prometheus.MustRegister(RequestTimeProc)
	prometheus.MustRegister(NewgoroutineCollector())
}

//GoroutineNum
type GoroutineCollector struct {
	GoroutineSum *prometheus.Desc
}

func (g *GoroutineCollector) Describe(descs chan<- *prometheus.Desc)  {
	descs <- g.GoroutineSum
}
func (g *GoroutineCollector) Collect(metrics chan<- prometheus.Metric)  {
	metrics <- prometheus.MustNewConstMetric(
		g.GoroutineSum,
		prometheus.GaugeValue,
		float64(runtime.NumGoroutine()),
		"GoroutineNum",
	)
}

func NewgoroutineCollector() *GoroutineCollector{
	return &GoroutineCollector{
		GoroutineSum:prometheus.NewDesc(
			"userweb_Goroutine_Num",
			"Goroutine_Num",
			[]string{"name"},
			nil,
		),
	}
}



func Before(c *context.Context)  {
	Pathcounter.WithLabelValues(c.Input.URL()).Inc()
	Pathcounter.WithLabelValues("sum").Inc()
	c.Input.SetData("stime",time.Now())
}


func After(c *context.Context)  {
	code := strconv.Itoa(c.ResponseWriter.Status)
	StatusCodecounter.WithLabelValues(code).Inc()
	stime := c.Input.GetData("stime")
	if stime != nil {
		if t,ok := stime.(time.Time);ok {
			rtime := time.Now().Sub(t).Nanoseconds()
			totaltime,_ := strconv.ParseFloat(fmt.Sprintf("%.3f", float64(rtime)/1000/1000), 64)
			RequestTimeProc.WithLabelValues(c.Input.URL()).Observe(totaltime)
		}

	}
}