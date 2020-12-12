package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header.Get("Cookie"))

		// Response Set-Cookie
		// w.Header().Add("set-cookie", "xxx=xxxxxx")
		sid := &http.Cookie{
			Name:  "sid",
			Value: "xxxx",
		}
		http.SetCookie(w, sid)
	})
	http.ListenAndServe(":9999", nil)
}
