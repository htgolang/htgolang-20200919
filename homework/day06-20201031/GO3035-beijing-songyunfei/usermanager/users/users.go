package users

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
	"usermanager/userutils"
)

const ph  = "./users.csv"

type Userinfo struct {
	Id int
	Name string
	Addr string
	Tel string
	Birthday time.Time
	passwd string
}

type Userdb struct {
	UserSlice []Userinfo
}


//生产ID
func Genid(usersl *Userdb) int {
	if len(usersl.UserSlice) == 0 {
		return 1
	}
	id := 1
	for i := 0; i < len(usersl.UserSlice); i++{
		tid := usersl.UserSlice[i].Id
		if id < tid{
			id = tid
		}
	}
	return id +1
}

//通过ID查找用户
func (u *Userdb) FindByid(id int) (index int,err error) {
	for k,v := range u.UserSlice {
		uid := v.Id
		if uid == id {
			return k,nil
		}
	}
	return -1, fmt.Errorf("未找到")

}
// 新增用户
func (u *Userdb) Add(name,addr,tel,passwd,birthday string  )  error {
	for _,v := range u.UserSlice{
		if v.Name == name {
			return fmt.Errorf("用户名:%s已存在.\n",name)
		}
	}
	b,err:= time.ParseInLocation("2006-01-02", birthday,time.Local)
	if err != nil {
		return err
	}
	newuser := Userinfo{
		Id:       Genid(u),
		Name:     name,
		Addr:     addr,
		Tel:      tel,
		Birthday: b,
		passwd:   userutils.Summd5(passwd),
	}
	u.UserSlice = append(u.UserSlice,newuser)
	u.Sync()
	return nil
}

//通过ID删除用户
func (u *Userdb) Del(id int ) error {
	index, err := u.FindByid(id)
	if err != nil {
		return  fmt.Errorf("未找到用户:%d",id)
	} else {
		if index == 0 && len(u.UserSlice) == 1 {
			u.UserSlice = make([]Userinfo,0)
			return nil
		}
		if index == len(u.UserSlice)-1 {
			u.UserSlice = u.UserSlice[:index]
			return nil
		}
		u.UserSlice = append(u.UserSlice[:index], u.UserSlice[index+1:]...)
	}
	u.Sync()
	return nil
}
// 修改索引为index的用户信息
func (u *Userdb) Modify(index int, name,addr,tel,bri string ) error {
	for _,v := range u.UserSlice{
		if v.Name == name {
			return fmt.Errorf("用户名:%s已存在.\n",name)
		}
	}
	u.UserSlice[index].Name = name
	u.UserSlice[index].Addr = addr
	u.UserSlice[index].Tel = tel
	//u.UserSlice[index].passwd = passwd
	u.Sync()
	return nil
}

//通过关键字查找用户 返回userinfo结构体和error
func (u *Userdb) QueryUser(str string) (f Userinfo, ok bool) {
	for _,v := range u.UserSlice {
		if strings.Contains(v.Name,str) || strings.Contains(v.Addr,str)|| strings.Contains(v.Tel,str) {
			return v,true
		}
	}
	return Userinfo{}, false
}

//认证功能
func (u *Userdb) Auth(username, passwd string) bool {
	for _,v := range u.UserSlice{
		if username == v.Name && userutils.Summd5(passwd) == v.passwd {
			return true
		}
	}
	return false
}

//加载csv
func (u *Userdb) Load()  {
	file,err := os.OpenFile(ph,os.O_CREATE|os.O_RDONLY,os.ModePerm)
	defer func() {
		_ = file.Close()
	}()
	if err != nil{
		fmt.Println(err)
	}
	newrd := csv.NewReader(file)
	for{
		line,err := newrd.Read()
		if err != nil{
			if err == io.EOF{
				break
			}
			fmt.Printf("读取错误:%s",err)
		}
		id ,_:= strconv.Atoi(line[0])
		b,err:= time.ParseInLocation("2006-01-02", line[4],time.Local)
		if err != nil{
			fmt.Println("生日转换错误")
			fmt.Println(err)
		}
		u.UserSlice = append(u.UserSlice,Userinfo{
			Id:       id,
			Name:     line[1],
			Addr:     line[2],
			Tel:      line[3],
			Birthday: b,
			passwd:   line[5],
		})
	}
}

//写入csv
func (u *Userdb) Sync()  {
	file,err := os.OpenFile(ph,os.O_TRUNC|os.O_WRONLY|os.O_CREATE,os.ModePerm)
	if err !=nil{
		fmt.Println(err)
	}
	defer func() {
		_ = file.Close()
	}()
	write := csv.NewWriter(file)
	for _, user := range u.UserSlice{
		id := strconv.Itoa(user.Id)
		b := user.Birthday.Format("2006-1-2")
		record := []string{id,user.Name,user.Addr,user.Tel,b,user.passwd}
		err := write.Write(record)
		if err != nil{
			fmt.Printf("写入错误:%s",err)
		}
	}
	write.Flush()
}