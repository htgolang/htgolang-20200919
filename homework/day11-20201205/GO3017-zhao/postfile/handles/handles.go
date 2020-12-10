package handles

import (
	"fmt"
	"net/http"
	"text/template"
	"zhao/utils"
)

func Postfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseMultipartForm(1024 * 1024)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
	// fmt.Println(r.MultipartForm.File)
	fmt.Println(">>>>>>>>", r.Form)
	fmt.Println(r.Header)
	for head, fileHandles := range r.MultipartForm.File {
		fmt.Println("=========", head, fileHandles)
		for _, fileHandle := range fileHandles {
			filename := fileHandle.Filename

			file, err := fileHandle.Open()
			defer file.Close()
			if err != nil {
				fmt.Println(err, "fileHandle open file false")
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}

			if head == "file" {
				//写入到文件
				path := r.Header["Path"] ///////////////////////////////////测试下不带头和带头两种情况
				if len(path) != 1 {
					w.WriteHeader(500)
					w.Write([]byte("上传时带path头[Path:PATH]"))
					file.Close()
					return
				}

				err := utils.StoreToFile(filename, path[0], file)
				if err != nil {
					w.WriteHeader(500)
					w.Write([]byte(err.Error()))
					file.Close()
					return
				}

			}
			if head == "mysql" {
				//写入到数据库
				err := utils.StoreToMysql(filename, file)
				if err != nil {
					w.WriteHeader(500)
					w.Write([]byte(err.Error()))
					file.Close()
					return
				}
			}
		}
	}
	w.Write([]byte("sucess"))
}

func List(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	err = tpl.ExecuteTemplate(w, "index.html", utils.QueryFileSlice())
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
}

func Result(w http.ResponseWriter, r *http.Request) {
	fileName := r.FormValue("file")
	var ipchannel chan [][]string = make(chan [][]string)
	var codechannel chan [][]string = make(chan [][]string)

	go GO(func() {
		utils.QueryField(fileName, "ip", ipchannel)
	})

	go GO(func() {
		utils.QueryField(fileName, "code", codechannel)
	})
	s := make(map[string][][]string, 2)
	s["ip"] = <-ipchannel
	s["code"] = <-codechannel
	tpl, err := template.ParseFiles("template/list.html")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	err = tpl.ExecuteTemplate(w, "list.html", s)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
}

func GO(x func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("-------------------", err)
		}
	}()
	x()
}
