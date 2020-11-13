package users

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonUserDb struct {
	Baseuser //继承baseuser
}

//增加load方法
func (u *JsonUserDb) Load() error {
	file,err := os.OpenFile(Savepath,os.O_CREATE|os.O_RDONLY,os.ModePerm)
	defer func() {
		_ = file.Close()
	}()
	if err != nil{
		return fmt.Errorf("打开文件失败:%s",err)
	}
	userdecoder := json.NewDecoder(file)
	err = userdecoder.Decode(&u.UserSlice)
	if err != nil {
		return err
	}
	return nil
}

//写入csv
func (u *JsonUserDb) Sync() error {
	file,err := os.OpenFile(Savepath,os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	if err !=nil{
		return fmt.Errorf("打开文件失败:%s",err)
	}
	defer func() {
		_ = file.Close()
	}()
	write := json.NewEncoder(file)
	err = write.Encode(u.UserSlice)
	if err != nil {
		return fmt.Errorf("写入失败:%s",err)
	}
	return nil
}
