package controller

import (
	"log"
	"net/http"
	"time"
)

func wlog(r *http.Request)  {
	log.Printf("%s %s %s %s %s\n",time.Now().Format("2006-01-02 15:04:05"),r.RemoteAddr,r.Method,r.RequestURI,r.UserAgent())
	//fmt.Printf("%s %s %s %s %s\n",time.Now(),r.RemoteAddr,r.Method,r.RequestURI,r.UserAgent())
}
