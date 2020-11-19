package users

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"time"
)

type CsvUserDb struct {
	Baseuser //继承baseuser
}

//增加load方法
func (u *CsvUserDb) Load() error {
	//file,err := os.OpenFile(Savepath,os.O_CREATE|os.O_RDONLY,os.ModePerm)
	//defer func() {
	//	_ = file.Close()
	//}()
	//if err != nil{
	//	fmt.Println(err)
	//	return err
	//}
	newrd := csv.NewReader(Filefd)
	for{
		line,err := newrd.Read()
		if err != nil{
			if err == io.EOF{
				break
			}
			fmt.Printf("读取错误:%s",err)
			return err
		}
		id ,_:= strconv.Atoi(line[0])
		b,err:= time.ParseInLocation("2006-01-02", line[4],time.Local)
		if err != nil{
			fmt.Println("生日转换错误",err)
			return err
		}
		u.UserSlice = append(u.UserSlice,Userinfo{
			Id:       id,
			Name:     line[1],
			Addr:     line[2],
			Tel:      line[3],
			Birthday: b,
			Passwd:   line[5],
		})

	}
	if len(u.UserSlice) == 0 {
		return io.EOF
	}
	return nil
}

//写入csv
func (u *CsvUserDb) Sync() error {
	queue = append(queue,u.UserSlice)
	return nil
}

func writetocsv(u []Userinfo, f *os.File) error {
	w := csv.NewWriter(f)
	for _, user := range u{
		id := strconv.Itoa(user.Id)
		b := user.Birthday.Format("2006-1-2")
		record := []string{id,user.Name,user.Addr,user.Tel,b,user.Passwd}
		err := w.Write(record)
		if err != nil{
			fmt.Printf("写入错误:%s",err)
			return err
		}
	}
	w.Flush()
	return nil
}
// 滚动存储
func (u *CsvUserDb) RotateSave() error {
	l := len(queue)
	if l <= QueueLen {
		for i:=0;i<l;i++{
			if i == l-1{
				err := writetocsv(queue[i],Filefd)
				if err != nil{
					fmt.Println(err)
					return err
				}
				continue
			}
			name := path.Dir(Savepath)
			filename := path.Join(name,fmt.Sprintf("user.csv.%d",i))
			file,err := os.OpenFile(filename,os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
			if err != nil {
				return fmt.Errorf("创建失败:%s",err)
			}
			err = writetocsv(queue[i],file)
			if err != nil{
				fmt.Println(err)
				return err
			}
		}

	}
	if l > QueueLen{
		for k,v := range queue[l-QueueLen:]{
			if k == QueueLen -1 {
				write := csv.NewWriter(Filefd)
				for _, user := range v{
					id := strconv.Itoa(user.Id)
					b := user.Birthday.Format("2006-1-2")
					record := []string{id,user.Name,user.Addr,user.Tel,b,user.Passwd}
					err := write.Write(record)
					if err != nil{
						fmt.Printf("写入错误:%s",err)
						return err
					}
				}
				write.Flush()
				continue
			}
			name := path.Dir(Savepath)
			filename := path.Join(name,fmt.Sprintf("user.csv.%d",k))
			file,err := os.OpenFile(filename,os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
			if err != nil {
				return fmt.Errorf("创建失败:%s",err)
			}
			err = writetocsv(v,file)
			if err != nil{
				fmt.Println(err)
				return err
			}
		}
		return nil
	}
	return fmt.Errorf("未知错误")

}

