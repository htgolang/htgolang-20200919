package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fileserver/process"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	dbtype := "mysql"
	// dsn := "root:admin@ADMIN123456@tcp(192.168.10.128:3306)/log_analysis?parseTime=true&loc=Local&charset=utf8mb4"
	dsn := "test:NDg3NTBi@tcp(127.0.0.1:3306)/usermanager?parseTime=true&loc=Local&charset=utf8mb4"	
	err := process.InitDB(dbtype,dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err := process.CloseDb()
		if err != nil {
			fmt.Println(err)
		}
	}()
	addr := "0.0.0.0:8888"
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/getdata", dataapi)
	http.HandleFunc("/show", showweb)
	fmt.Println("启动完成")
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
	}

}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl, err := template.ParseFiles("./template/upload.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		Fdata := process.Getfile()
		err = tpl.ExecuteTemplate(w, "upload.html",Fdata)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		err := r.ParseMultipartForm(1024 * 1024)
		if err != nil {
			fmt.Println(err)
			return
		}
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}

		defer func() {
			_ = file.Close()
		}()
		id,err := process.Filemeta(handler.Filename,handler.Size)
		if err != nil {
			fmt.Println(err)
			return
		}
		rder := bufio.NewReader(file)
		for {
			data,err:= rder.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println(err)
			}
			if err := process.Loginsert(id,data); err != nil {
				fmt.Println(err)
			}

		}

	}
}

func dataapi(w http.ResponseWriter, r *http.Request)  {
	vars := r.URL.Query()
	fid := vars.Get("fileid")
	if fid == "" {
		rs := strings.NewReader("Id 不能为空")
		_,_ = io.Copy(w,rs)
		return
	}
	id,err := strconv.ParseInt(fid,10,64)
	if err != nil {
		rs := strings.NewReader("非法Id")
		_,_ = io.Copy(w,rs)
		return
	}
	tdata := process.Readformdb(id)
	if b,err := json.Marshal(tdata);err == nil{
		rs := bytes.NewReader(b)
		_,_ = io.Copy(w,rs)
		return
	}
	fmt.Println(err)
	return

}

func showweb(w http.ResponseWriter, r *http.Request)  {
	vars := r.URL.Query()
	fid := vars.Get("fileid")
	if fid == "" {
		rs := strings.NewReader("Id 不能为空")
		_,_ = io.Copy(w,rs)
		return
	}
	id,err := strconv.ParseInt(fid,10,64)
	if err != nil {
		rs := strings.NewReader("非法Id")
		_,_ = io.Copy(w,rs)
		return
	}
	tpl, err := template.ParseFiles("./template/bar.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	type sdata struct {
		Fid string
		Ipdata []map[string]string
	}
	var data sdata
	data.Fid = fid
	data.Ipdata = process.Readformdb(id).Databar
	if err = tpl.ExecuteTemplate(w, "bar.html",data); err != nil{
		fmt.Println(err)
		return
	}
}