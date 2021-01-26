说明:
说明: 导入usermanager.sql建立表 修改config.ini中的配置 启动服务 访问：http://127.0.0.1:8888/metrics/

```
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 2
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
# HELP userwebCode_total_v1 Status_Code_Count
# TYPE userwebCode_total_v1 counter
userwebCode_total_v1{code="200",name="status_code"} 2
# HELP userweb_Goroutine_Num Goroutine_Num
# TYPE userweb_Goroutine_Num gauge
userweb_Goroutine_Num{name="GoroutineNum"} 10
# HELP userweb_time_total_v1 Request_Time_ms
# TYPE userweb_time_total_v1 histogram
userweb_time_total_v1_bucket{name="request_time",request_path="/favicon.ico",le="0"} 0
userweb_time_total_v1_bucket{name="request_time",request_path="/favicon.ico",le="10"} 1
userweb_time_total_v1_bucket{name="request_time",request_path="/favicon.ico",le="20"} 2
userweb_time_total_v1_bucket{name="request_time",request_path="/favicon.ico",le="+Inf"} 2
userweb_time_total_v1_sum{name="request_time",request_path="/favicon.ico"} 16.002000000000002
userweb_time_total_v1_count{name="request_time",request_path="/favicon.ico"} 2
# HELP userweb_total_v1 Request_Path_Count
# TYPE userweb_total_v1 counter
userweb_total_v1{name="request",path="/favicon.ico"} 2
userweb_total_v1{name="request",path="sum"} 2
```
