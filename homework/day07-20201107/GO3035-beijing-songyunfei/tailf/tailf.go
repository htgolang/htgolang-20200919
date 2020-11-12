package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	n := flag.Int("n",5,"最后行数")
	f := flag.String("f","","文件")
	flag.Parse()
	file,err := os.OpenFile(*f,os.O_RDONLY,os.ModePerm)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer func() {
		_ = file.Close()
	}()
	rd := bufio.NewReader(file)
	s := getseek(file,*n)
	_,err = file.Seek(s+1,2)
	if err != nil{
		fmt.Println(err)
	}
	for {
		line,_,err := rd.ReadLine()
		if err != nil {
			if err == io.EOF{
				fmt.Println("xx结束了")
				return
			}
			fmt.Println(err)
			break
		}
		fmt.Println(string(line))
	}

}
//获取最后N行位置
func getseek(f *os.File, lastn int) int64  {
	var tmp int64 = -1
	c:=0
	for c<lastn {
		ret,err := f.Seek(tmp,2)
		p := make([]byte,1)
		_,err = f.Read(p)
		if err != nil {
			if err == io.EOF{
				return ret
			}
			fmt.Println(err)
			break
		}
		if string(p) == "\n" {
			c++
		}
		tmp--
	}
	return tmp
}
