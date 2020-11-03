package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

func lsdir(p string) {
	var cl int
	f,err := os.Stat(p)
	if err !=nil {
		fmt.Println(err)
	}
	if f.IsDir(){
		nf,err := os.Open(p)
		defer func() {
			_ = nf.Close()
		}()
		if err !=nil{
			fmt.Println(err)
			return
		}
		nd,err := nf.Readdirnames(0)
		if err != nil{
			fmt.Println(err)
			return
		}
		for _,v := range nd{
			np := path.Join(p,v)
			lsdir(np)
		}
	}else {
		ex := filepath.Ext(f.Name())
		if ex == ".go" || ex == ".cgo"{
			cl = count(p)
			fmt.Printf("%s 代码行数:%d\n",p,cl)

		}
	}

}

func count(p string) int  {
	f ,err:= os.Open(p)
	defer func() {
		err := f.Close()
		if err !=nil{
			fmt.Println(err)
		}
	}()
	if err != nil{
		fmt.Println(err)
	}
	nf := bufio.NewReader(f)
	var ct int
	for{
		s,err := nf.ReadBytes('\n')
		if err !=nil{
			break
		}
		if ma,_:=regexp.MatchString("^.*[a-zA-Z0-9].*",string(s));ma{
			ct++
		}
	}
	return ct

}


func main()  {
	//参数
	var d = flag.String("d","","目录路径")
	flag.Parse()
	if d != nil{
		//注释
		lsdir(*d)
	}else {
		fmt.Println("不能为空")
	}

}
