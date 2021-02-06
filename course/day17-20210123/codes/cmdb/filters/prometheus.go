package filters

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/context"

	"github.com/prometheus/client_golang/prometheus"
)

// 总请求次数 Counter
// 每个URL请求次数 Counter 带可变Label
// 状态码统计 Counter 带可变Label
// 每个URL请求时间 Histogram 带可变Label

// 存在Controller
var (
	totalRequest = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "cmdb_request_total",
		Help: "",
	})
	urlRequest = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cmdb_request_url_total",
		Help: "",
	}, []string{"url"})

	statusCode = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cmdb_status_code_total",
		Help: "",
	}, []string{"code"})

	elapsedTime = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "cmdb_request_url_elapsed_time",
		Help: "",
	}, []string{"url"})
)

func init() {
	prometheus.MustRegister(totalRequest, urlRequest, statusCode, elapsedTime)
}

func BeferExec(ctx *context.Context) {
	totalRequest.Inc()
	urlRequest.WithLabelValues(ctx.Input.URL()).Inc()
	ctx.Input.SetData("stime", time.Now())
}

func AfterExec(ctx *context.Context) {
	statusCode.WithLabelValues(strconv.Itoa(ctx.ResponseWriter.Status)).Inc()
	stime := ctx.Input.GetData("stime")
	if stime != nil {
		if t, ok := stime.(time.Time); ok {
			elapsed := time.Now().Sub(t)
			elapsedTime.WithLabelValues(ctx.Input.URL()).Observe(float64(elapsed))
		}
	}
}
