package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
)

type IPTOP struct {
	IP  string
	Num int
}

type IPSlice []IPTOP

func (a IPSlice) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a IPSlice) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a IPSlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].Num < a[i].Num
}

func FileList() []string {
	fileinfo, _ := ioutil.ReadDir("./logs")
	var result []string
	for _, v := range fileinfo {
		if v.IsDir() {
			return nil
		} else {
			result = append(result, v.Name())
		}
	}
	return result
}

func iplist(filename string) []IPTOP {
	var IPList map[string]int
	IPList = make(map[string]int)
	f, _ := os.Open("./logs/" + filename)
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		s := strings.Split(line, " ")
		IP := s[0]
		num, ok := IPList[IP]
		if ok {
			IPList[IP] = num + 1
		} else {
			IPList[IP] = 1
		}
	}
	result := []IPTOP{}
	for k, v := range IPList {
		result = append(result, IPTOP{k, v})
	}

	sort.Sort(IPSlice(result))
	if a := len(result); a >= 10 {
		result = result[0:10]
	} else {
		result = result[0:a]
	}
	return result
}

func main() {
	addr := ":8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			f := FileList()
			tpl := template.Must(template.ParseFiles("template/index.html"))
			err := tpl.ExecuteTemplate(w, "index.html", f)
			if err != nil {
				fmt.Println(err)
			}
		}
		if r.Method == "POST" {
			http.Redirect(w, r, "/table", 302)
		}
	})
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			f := FileList()
			tpl := template.Must(template.ParseFiles("template/upload.html"))
			err := tpl.ExecuteTemplate(w, "upload.html", f)
			if err != nil {
				fmt.Println(err)
			}
		}
		if r.Method == "POST" {
			_ = r.ParseMultipartForm(1024 * 1024)
			if fileheaders, ok := r.MultipartForm.File["uploadfile"]; ok {
				for _, fileheader := range fileheaders {
					nfile, err := os.Create("./logs/" + fileheader.Filename)
					if err != nil {
						fmt.Println(err)
						return
					}
					file, err := fileheader.Open()
					if err != nil {
						fmt.Println(err)
						return
					}
					_, err = io.Copy(nfile, file)
					if err != nil {
						fmt.Println(err)
						return
					}
				}
			}
			http.Redirect(w, r, "/", 302)
		}
	})
	http.HandleFunc("/table", func(w http.ResponseWriter, r *http.Request) {
		msg := iplist(r.FormValue("filename"))
		tpl := template.Must(template.ParseFiles("template/table.html"))
		err := tpl.ExecuteTemplate(w, "table.html", msg)
		if err != nil {
			fmt.Println(err)
		}
	})
	_ = http.ListenAndServe(addr, nil)
}
