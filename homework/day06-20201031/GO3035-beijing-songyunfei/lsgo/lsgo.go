package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func lsdir(p string)  {
	f,err := os.Stat(p)
	if err !=nil {
		fmt.Println(err)
	}
	if f.IsDir(){
		nf,err := os.Open(p)
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
			fmt.Println(p)
		}
	}

}

func main()  {
	var d = flag.String("d","","目录路径")
	flag.Parse()
	lsdir(*d)
}
