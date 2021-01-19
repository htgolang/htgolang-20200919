说明:
导入usermanager.sql建立表
修改config.ini中的配置
启动服务
访问：http://127.0.0.1:8888/metrics/

metrics示例:
# HELP userwebCode_total_v1 Status_Code_Count
# TYPE userwebCode_total_v1 counter
userwebCode_total_v1{code="200",name="status_code"} 13
# HELP userweb_time_total_v1 Request_Time_ms
# TYPE userweb_time_total_v1 histogram
userweb_time_total_v1_bucket{name="request_time",request_path="/",le="0"} 0
userweb_time_total_v1_bucket{name="request_time",request_path="/",le="10"} 1
userweb_time_total_v1_bucket{name="request_time",request_path="/",le="20"} 1
userweb_time_total_v1_bucket{name="request_time",request_path="/",le="+Inf"} 1
userweb_time_total_v1_sum{name="request_time",request_path="/"} 1.998
userweb_time_total_v1_count{name="request_time",request_path="/"} 1
userweb_time_total_v1_bucket{name="request_time",request_path="/favicon.ico",le="0"} 0
userweb_time_total_v1_bucket{name="request_time",request_path="/favicon.ico",le="10"} 8
userweb_time_total_v1_bucket{name="request_time",request_path="/favicon.ico",le="20"} 8
userweb_time_total_v1_bucket{name="request_time",request_path="/favicon.ico",le="+Inf"} 8
userweb_time_total_v1_sum{name="request_time",request_path="/favicon.ico"} 19.988
userweb_time_total_v1_count{name="request_time",request_path="/favicon.ico"} 8
userweb_time_total_v1_bucket{name="request_time",request_path="/modify?Id=8",le="0"} 0
userweb_time_total_v1_bucket{name="request_time",request_path="/modify?Id=8",le="10"} 1
userweb_time_total_v1_bucket{name="request_time",request_path="/modify?Id=8",le="20"} 1
userweb_time_total_v1_bucket{name="request_time",request_path="/modify?Id=8",le="+Inf"} 1
userweb_time_total_v1_sum{name="request_time",request_path="/modify?Id=8"} 3.996
userweb_time_total_v1_count{name="request_time",request_path="/modify?Id=8"} 1
userweb_time_total_v1_bucket{name="request_time",request_path="/query",le="0"} 1
userweb_time_total_v1_bucket{name="request_time",request_path="/query",le="10"} 3
userweb_time_total_v1_bucket{name="request_time",request_path="/query",le="20"} 3
userweb_time_total_v1_bucket{name="request_time",request_path="/query",le="+Inf"} 3
userweb_time_total_v1_sum{name="request_time",request_path="/query"} 2.955
userweb_time_total_v1_count{name="request_time",request_path="/query"} 3
# HELP userweb_total_v1 Request_Path_Count
# TYPE userweb_total_v1 counter
userweb_total_v1{name="request",path="/"} 1
userweb_total_v1{name="request",path="/favicon.ico"} 8
userweb_total_v1{name="request",path="/modify?Id=8"} 1
userweb_total_v1{name="request",path="/query"} 3
userweb_total_v1{name="request",path="sum"} 13

