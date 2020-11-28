package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// 提交数据常用规则
	/*
		1. 一个名字对应一个值
		2. 提交数据方式
			浏览器
				GET		参数常在URL
				POST
					x-www-form-urlencoded
					multipart/form-data
			第三方工具: curl, client
				GET
					参数常在URL
					某些情况参数可放在Body(取决于client工具/服务是否支持)
				POST
					x-www-form-urlencoded
					multipart/form-data
					application/json
				DELETE
					参数常在URL(取决于client工具/服务是否支持)
				PUT
					body
				HEAD
					参数常在URL(取决于client工具/服务是否支持)

	*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.FormValue("x"))
		fmt.Println(r.FormValue("y"))
		fmt.Println(r.PostFormValue("y"))
		if file, fileHeader, err := r.FormFile("z"); err == nil {
			fmt.Println(fileHeader.Filename, fileHeader.Size)
			io.Copy(os.Stdout, file)
		}
	})
	http.ListenAndServe("0.0.0.0:8080", nil)
}
