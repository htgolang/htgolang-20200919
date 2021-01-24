package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

var Pathcounter *prometheus.CounterVec
var StatusCodecounter *prometheus.CounterVec
var RequestTimeProc *prometheus.HistogramVec

func Initmetrics(namespace, subsystem, name string )  {
	Pathcounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   namespace,
			Subsystem:   subsystem,
			Name:        name,
			Help:        "Request_Path_Count",
			ConstLabels: prometheus.Labels{"name": "request"},
		},
		[]string{"path"},
	)
	StatusCodecounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   namespace+"Code",
			Subsystem:   subsystem,
			Name:        name,
			Help:        "Status_Code_Count",
			ConstLabels: prometheus.Labels{"name": "status_code"},
		},
		[]string{"code"},
	)
	RequestTimeProc = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace:   namespace+"_time",
			Subsystem:   subsystem,
			Name:        name,
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

//继承ResponseWriter用于获取响应状态码
type response struct {
	http.ResponseWriter
	StatusCode int
}

func (w *response) WriteHeader(code int)  {
	w.StatusCode = code
	w.ResponseWriter.WriteHeader(code)
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

//PathCounter
func MetriscProc(handl http.HandlerFunc) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		//初始设置为200
		rw := &response{
			w,
			200,
		}
		Pathcounter.WithLabelValues(r.RequestURI).Inc()
		Pathcounter.WithLabelValues("sum").Inc()
		start := time.Now()
		handl(rw,r)
		code := strconv.Itoa(rw.StatusCode)
		StatusCodecounter.WithLabelValues(code).Inc()
		ms := time.Now().Sub(start).Nanoseconds()
		rtime,_ := strconv.ParseFloat(fmt.Sprintf("%.3f", float64(ms)/1000/1000), 64)
		RequestTimeProc.WithLabelValues(r.RequestURI).Observe(rtime)
	}

}
