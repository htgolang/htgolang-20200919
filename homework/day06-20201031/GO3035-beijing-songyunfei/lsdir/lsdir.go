package main

import (
	"flag"
	"fmt"
	"os"
	"path"
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
		fmt.Println(f.Name())
	}

}

func main()  {
	var d = flag.String("d","","目录路径")
	flag.Parse()
	lsdir(*d)
}
