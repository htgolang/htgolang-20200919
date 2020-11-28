package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	addr := "0.0.0.0:8080"
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(strings.Repeat("-", 30))

		// 请求行
		fmt.Println("method:", r.Method)
		fmt.Println("url:", r.URL)
		fmt.Println("protocol:", r.Proto)

		// 请求头
		fmt.Println(r.Host)

		header := r.Header
		fmt.Println(header.Get("User-Agent"))

		fmt.Println(header.Get("Token"))

		// 请求体
		fmt.Println("body:")
		io.Copy(os.Stdout, r.Body)

		fmt.Fprint(w, time.Now().Format("2006-01-02 15:04:05"))

	})
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err)
}
