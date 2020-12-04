package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type file struct {
	Filename string
	Path     string
}

type Count struct {
	Text  string
	Count int
}

type Slice []Count

func (s Slice) Len() int {
	return len(s)
}

func (s Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Slice) Less(i, j int) bool {
	return s[i].Count > s[j].Count
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u := make([]*file, 0, 20)
		if r.Method == "GET" {
			filepath.Walk("file", func(path string, info os.FileInfo, err error) error {
				if !info.IsDir() {
					u = append(u, &file{info.Name(), path})
				}
				return nil
			})

			t, err := template.ParseFiles("index.html")
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			err = t.ExecuteTemplate(w, "index.html", u)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
			}

		}
	})
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		path := r.Form.Get("file")
		fmt.Println(path)
		file, err := os.Open(path)
		defer file.Close()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		ipres, coderes := sortSlice(file)
		fmt.Println(ipres)
		t, err := template.ParseFiles("list.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		u := struct {
			IP   Slice
			Code Slice
		}{
			IP:   ipres,
			Code: coderes,
		}
		fmt.Println(t.ExecuteTemplate(w, "list.html", u))
	})

	http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseMultipartForm(1024 * 1024)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			for k, fileHandles := range r.MultipartForm.File {
				for _, fileHandle := range fileHandles {
					f, err := fileHandle.Open()
					defer f.Close()
					if err != nil {
						w.WriteHeader(500)
						w.Write([]byte(err.Error()))
						return
					}
					file, err := os.Create(filepath.Join("file", k, fileHandle.Filename))
					defer file.Close()
					if err != nil {
						w.WriteHeader(500)
						w.Write([]byte(err.Error()))
						return
					}

					_, err = io.Copy(file, f)
					if err != nil {
						w.WriteHeader(500)
						w.Write([]byte(err.Error()))
						return
					}
				}
			}
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func sortSlice(r io.Reader) (i Slice, c Slice) {

	ipcount := make(map[string]int)
	statuscode := make(map[string]int)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		res := strings.Fields(scanner.Text())
		ipcount[res[0]]++
		statuscode[res[8]]++
	}
	ip := make(Slice, 0, 100)
	st := make(Slice, 0, 100)
	for k, v := range ipcount {
		ip = append(ip, Count{k, v})
	}
	for k, v := range statuscode {
		st = append(st, Count{k, v})
	}
	// sort.Slice(ip, func(i, j int) bool {
	// 	return ip[i].Count > ip[j].Count
	// })
	sort.Sort(ip)
	sort.Sort(st)
	fmt.Println(ip)
	return ip[:10], st[:10]
}
