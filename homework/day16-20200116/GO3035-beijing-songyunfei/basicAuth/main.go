package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

var authfile string
func main()  {
	addr := ":8080"
	authfile = "C:\\Users\\yunfei.song\\go\\src\\basicAuth\\authfile"
	//str,_ := genpasswd("123456")
	//fmt.Println(string(str))
	http.HandleFunc("/time",authCheck(web))
	err := http.ListenAndServe(addr,nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}



func web(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.RemoteAddr)
	now := time.Now().Format("2006-01-02 15:04:05")
	_,err := w.Write([]byte(now))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func authCheck(procfunc http.HandlerFunc) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		basestr := r.Header.Get("Authorization")
		if basestr != "" && islogin(basestr) {
			procfunc(w,r)
			return
		}else {
			w.Header().Set("WWW-Authenticate","Basic")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
}


func islogin(basestr string) bool {
	f ,err := os.Open(authfile)
	defer func() {
		_ = f.Close()
	}()

	bf := bufio.NewReader(f)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	split := strings.Split(basestr," ")
	baseStr,err := base64.StdEncoding.DecodeString(split[1])
	if err != nil {
		fmt.Println("Error base64 str.")
		return false
	}
	npslic := strings.Split(string(baseStr),":")
	name := npslic[0]
	passwd := npslic[1]
	for {
		line,_,err := bf.ReadLine()
		if err != nil {
			if err == io.EOF{
				break
			}
			fmt.Println(err)
		}
		upslic := strings.Split(string(line),":")
		u := upslic[0]
		p := upslic[1]

		if name == u  {
			err = bcrypt.CompareHashAndPassword([]byte(p), []byte(passwd))
			if err == nil {
				return true
			}
			return false
		}
	}

	return false
}

func genpasswd(str string) ([]byte,error)  {
	return bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
}
