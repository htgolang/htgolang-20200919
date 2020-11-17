package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

const bfsize = 1024
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
	line := make([]byte,bfsize)
	for {
		n,err := rd.Read(line)
		if err != nil {
			if err == io.EOF{
				//等待1s
				time.Sleep(time.Second * 1)
				continue
			}
			fmt.Println(err)
			break
		}else {
			fmt.Print(string(line[:n]))
		}

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
