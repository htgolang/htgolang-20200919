package main

import (
"flag"
"fmt"
	"io"
	"os"
)

func isExists(p string) (bool,error)  {
	_, err := os.Stat(p)
	if err == nil{
		return true,nil
	}
	if os.IsExist(err){
		return false,nil
	}
	return false, err
}

func fsrc(p string) (f *os.File, err error)  {
	f,err = os.Open(p)
	return f,err
}

func fdest(p string) (f *os.File,err  error) {
	if b, _ := isExists(p);b{
		var s string
		fmt.Printf("%s已存在,是否覆盖[y/n]:",p)
		_,_ = fmt.Scanln(&s)
		if s == "y" || s == "Y" {
			f,err = os.OpenFile(p,os.O_TRUNC|os.O_WRONLY,os.ModePerm)
			return f, err
		}
		if s == "n"|| s == "N"{
			return nil,fmt.Errorf("cancel")
		}

	}
	f,err = os.Create(p)
	return f,err
}

func docopy(src,des *os.File) error {
	var b = make([]byte,1024)
	for{
		_,err := src.Read(b)
		if err == io.EOF{
			return nil
		}
		_,werr := des.Write(b)
		if werr !=nil{
			return werr
		}

	}
	// io.copy包
	//_,err := io.Copy(des,src)
	//return err
}

func main()  {
	var d = flag.String("d","","目标路径")
	var s = flag.String("s","","源文件路径")
	flag.Parse()
	sf,serr := fsrc(*s)
	if  serr != nil{
		fmt.Printf("读取错误:%s",serr)
		return
	}
	df,derr := fdest(*d)
	if  derr != nil{
		fmt.Printf("写入错误:%s\n",derr)
		return
	}
	err := docopy(sf,df)
	if err != nil{
		fmt.Println(err)
	}

	defer func() {
		_ = df.Close()
		_ = sf.Close()
	}()

}
