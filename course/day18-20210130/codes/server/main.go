package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	addr := ":9999"
	// http.ListenAndServe(addr, nil)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%d", time.Now().Unix())
	})

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		interrupt := make(chan os.Signal)
		signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-interrupt

		ctx, cancel := context.WithCancel(context.Background())
		server.Shutdown(ctx)
		cancel()
	}()
	fmt.Println(server.ListenAndServe())
	// http.ListenAndServe(addr, mux) // mux

	// go gin, beego, httprouter radix-tree, 所有的框架
	// FormValue
	// 响应 html/template
	//
}
