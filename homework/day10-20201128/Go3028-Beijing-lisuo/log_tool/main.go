package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl, err := template.New("home").ParseFiles("./home.html")
		if err != nil {
			panic(err)
		}
		if err := tpl.ExecuteTemplate(w, "home.html", ""); err != nil {
			panic(err)
		}

	})

	http.HandleFunc("/uploadPage/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("uploadPage==========start")
		tpl, err := template.New("upload").ParseFiles("./upload.html")
		if err != nil {
			panic(err)
		}
		if err := tpl.ExecuteTemplate(w, "upload.html", ""); err != nil {
			panic(err)
		}
		fmt.Println("uploadPage==========end")

	})

	http.HandleFunc("/upload/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("upload==========start")
		fmt.Println("r.URL: ", r.URL)
		// file size limited to 1G
		r.ParseMultipartForm(1000 << 20)
		fmt.Println("r.Form: ", r.Form)
		fmt.Println("r.Postform: ", r.PostForm)
		fmt.Println("r.MultipartForm: ", r.MultipartForm)
		if fileHeaders, ok := r.MultipartForm.File[r.Form["filename"][0]]; ok {
			for _, fileHeader := range fileHeaders {
				fmt.Println("fileHeader.Filename: ", fileHeader.Filename)
				fmt.Println("fileHeader.Size: ", fileHeader.Size)
				uploadFile = uploadDir + fileHeader.Filename
				newFile, err := os.Create(uploadFile)
				if err != nil {
					panic(err)
				}
				file, err := fileHeader.Open()
				if err != nil {
					panic(err)
				}
				io.Copy(newFile, file)
			}
		}
		http.Redirect(w, r, "/", 302)
		fmt.Println("upload==========end")

	})

	http.HandleFunc("/resultPage/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("resultPage==========start")
		tpl, err := template.New("result").ParseFiles("./result.html")
		if err != nil {
			panic(err)
		}
		if err := tpl.ExecuteTemplate(w, "result.html", ""); err != nil {
			panic(err)
		}
		fmt.Println("resultPage==========end")
	})

	// process the log file and  display result
	http.HandleFunc("/display/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("displayPage==========start")
		tpl, err := template.New("display").ParseFiles("./display.html")
		if err != nil {
			panic(err)
		}
		// calculate the result
		ipList, statusList, URLList, err = GetList(uploadFile)
		if err != nil {
			fmt.Fprint(w, "Have you upload a file?")
			http.Redirect(w, r, "/", 302)
			return
		}
		ipNum := GetNum(ipList)
		ipRankList = GetRank(ipNum)
		fmt.Println("ipRankList: ", ipRankList)
		statusNum := GetNum(statusList)
		statusRankList := GetRank(statusNum)
		fmt.Println("statusRankList: ", statusRankList)
		// get methodList from URLList(request Line)
		for _, url := range URLList {
			reqLine := strings.Split(url, " ")
			if strings.Contains(reqLine[0], `\x`) || strings.Contains(reqLine[0], `tmp`) ||
				strings.Contains(reqLine[0], `WIN`) {
				continue
			}
			methodList = append(methodList, reqLine[0])
		}
		methodNum := GetNum(methodList)
		methodRankList := GetRank(methodNum)
		fmt.Println("methodRankList: ", methodRankList)

		resultList := NewResult()

		// put ip to rank
		for i := 1; i <= rankLen; i++ {
			for k, v := range ipRankList {
				resultList[rankLen-i].IP = v
				resultList[rankLen-i].IPNum = k
			}

		}
		// put status code to rank
		for i := 1; i <= rankLen; i++ {
			for k, v := range statusRankList {
				s, _ := strconv.Atoi(v)
				resultList[rankLen-i].Status = s
				resultList[rankLen-i].StatusNum = k
			}
		}

		// put method to rank
		for i := 1; i <= rankLen; i++ {
			for k, v := range methodRankList {
				resultList[rankLen-i].Method = v
				resultList[rankLen-i].MethodNum = k
			}
		}
		if err := tpl.ExecuteTemplate(w, "display.html", resultList); err != nil {
			panic(err)
		}
		fmt.Println("displayPage==========end")

	})

	if err := http.ListenAndServe(":8889", nil); err != nil {
		panic(err)
	}

}

func NewResult() []Result {
	var rList = []Result{}
	for i := 1; i <= rankLen; i++ {
		r := Result{Rank: i}
		rList = append(rList, r)
	}
	return rList
}

func GetList(file string) (ipList, statusLit, URLList []string, err error) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
		return nil, nil, nil, err
	}
	//io.Copy(os.Stdout, f)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		ipLine := strings.Split(line, " ")

		ipList = append(ipList, ipLine[0])
		statusList = append(statusList, ipLine[8])

		URLLine := strings.Split(line, "\"")
		URLList = append(URLList, URLLine[1])
	}
	return ipList, statusList, URLList, nil
}

func GetNum(itermList []string) map[string]int {
	var itermNum = make(map[string]int)
	for _, iterm := range itermList {
		if _, ok := itermNum[iterm]; ok {
			itermNum[iterm]++
		} else {
			itermNum[iterm] = 1
		}
	}
	return itermNum
}

func GetRank(itermNum map[string]int) map[int]string {
	var rankList = make(map[int]string)
	for i := 0; i < rankLen; i++ {
		var max int
		var keyMax string
		// get max occured ip and the occurence
		for k, v := range itermNum {
			if v > max {
				max = v
				keyMax = k
			}
		}
		// put the max to rankList
		rankList[max] = keyMax
		// make the max to zero
		itermNum[keyMax] = 0
	}
	return rankList
}
