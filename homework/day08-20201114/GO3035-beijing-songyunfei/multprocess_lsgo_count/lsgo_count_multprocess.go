package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sync"
)

var workerQueue = make(chan int,4)

func lsdir(p string,wg *sync.WaitGroup) {
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
			lsdir(np,wg)
		}
	}else {
		ex := filepath.Ext(f.Name())
		if ex == ".go" || ex == ".cgo"{
			workerQueue <- 1
			wg.Add(1)
			go count(p,wg)
		}
	}

}

func count(p string, wg *sync.WaitGroup)   {
	_ = <- workerQueue
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
	fmt.Printf("%s 代码行数:%d\n",p,ct)
	wg.Done()

}


func main()  {
	//参数
	var d = flag.String("d","","目录路径")
	flag.Parse()
	var wg sync.WaitGroup
	if d != nil{
		//注释
		lsdir(*d,&wg)
		wg.Wait()
	}else {
		fmt.Println("不能为空")
	}

}
