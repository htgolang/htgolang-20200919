package users

import (
	"encoding/csv"
	"fmt"
	"io"
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
	//write := csv.NewWriter(Filefd)
	//for _, user := range u.UserSlice{
	//	id := strconv.Itoa(user.Id)
	//	b := user.Birthday.Format("2006-1-2")
	//	record := []string{id,user.Name,user.Addr,user.Tel,b,user.Passwd}
	//	err := write.Write(record)
	//	if err != nil{
	//		fmt.Printf("写入错误:%s",err)
	//		return err
	//	}
	//}
	//write.Flush()
	//return nil
}
// 滚动存储
func (u *CsvUserDb) RotateSave() error {
	l := len(queue)
	if l <= QueueLen {

	}
	fmt.Println("test")
	return nil

}

