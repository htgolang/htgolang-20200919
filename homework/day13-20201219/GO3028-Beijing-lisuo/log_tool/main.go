package main

import (
	"fmt"
	"net/http"
	"webtool/route"
)

type Result struct {
	Rank      int
	IP        string
	IPNum     int
	Status    int
	StatusNum int
	Method    string
	MethodNum int
}

var (
	uploadDir  = "/tmp/upload/"
	uploadFile string
	rankLen    = 10
	ipList     []string
	ipRankList = make(map[int]string)
	methodList []string
	URLList    []string
	statusList []string
)

func main() {
	route.Route()
	err := http.ListenAndServe(":8086", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
