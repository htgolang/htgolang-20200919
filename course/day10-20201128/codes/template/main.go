package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	// 解析模板 字符串
	text := "我的名字叫 {{.}}\n"
	tpl, err := template.New("test").Parse(text)
	if err != nil {
		fmt.Println(err)
	}

	tpl.Execute(os.Stdout, "kk")
	tpl.Execute(os.Stdout, 1)
	tpl.Execute(os.Stdout, true)
	tpl.Execute(os.Stdout, []int{1, 2, 3})
	tpl.Execute(os.Stdout, map[string]string{"1": "1"})

	text = "性别: {{ if . }} 男 {{ else }} 女 {{ end }}\n"

	tpl, _ = template.New("test").Parse(text)
	tpl.Execute(os.Stdout, false)

	// gt > , lt < , gte >=, lte <=, eq =, neq !=
	text = "性别: {{ if eq 1 . }} 男 {{ else }} 女 {{ end }}\n"

	tpl, _ = template.New("test").Parse(text)
	tpl.Execute(os.Stdout, 1)

	text = "学生列表: {{ range . }} {{.}}| {{ end }}" // 内部的.是range 外部的. 每次遍历的元素
	tpl, _ = template.New("test").Parse(text)
	tpl.Execute(os.Stdout, []string{"aaa", "bbb", "cccc"})

	text = "第一个元素: {{ index . 1 }}"
	tpl, _ = template.New("test").Parse(text)
	tpl.Execute(os.Stdout, []string{"aaa", "bbb", "cccc"})

	text = "name: {{ .name }} addr: {{ .addr }}xxxxxxxx\n"
	tpl, _ = template.New("test").Parse(text)
	tpl.Execute(os.Stdout, map[string]string{"name": "xxx"})

	text = "name: {{ .Name }} addr: {{ .Addr }}xxxxxxxx\n"
	tpl, _ = template.New("test").Parse(text)
	tpl.Execute(os.Stdout, struct{ Name, Addr string }{"xx", "yyy"})

	tpl, _ = template.ParseFiles("user.html")
	tpl.ExecuteTemplate(os.Stdout, "user.html", struct{ Name, Addr string }{"xx", "yyy"})
}
