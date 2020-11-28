package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	response, err := http.Get("http://localhost:8080?x=1&y=2")
	fmt.Println(err)
	io.Copy(os.Stdout, response.Body)

	response, err = http.Head("http://localhost:8080?x=1&y=2")
	fmt.Println(err)
	io.Copy(os.Stdout, response.Body)

	// x-www-form-urlencoded
	values := url.Values{}
	values.Add("x", "1")
	values.Add("x", "2")
	values.Set("y", "1")
	values.Set("y", "2")
	response, err = http.PostForm("http://localhost:8080", values)
	fmt.Println(err)
	io.Copy(os.Stdout, response.Body)

	// application/json
	// f, _ := os.Open("test.json")
	reader := strings.NewReader(`
		{"a":"x"}
	`)
	response, err = http.Post("http://localhost:8080", "application/json", reader)
	fmt.Println(err)
	io.Copy(os.Stdout, response.Body)

	buffer := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buffer)

	fileWiter, _ := writer.CreateFormFile("z", "xx.json")

	f, _ := os.Open("test.json")
	io.Copy(fileWiter, f)
	writer.Close()

	response, err = http.Post("http://localhost:8080", "multipart/form-data", buffer)
}
