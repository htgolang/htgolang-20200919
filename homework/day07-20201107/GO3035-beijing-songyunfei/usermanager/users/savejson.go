package users

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

type JsonUserDb struct {
	Baseuser //继承baseuser
}

//增加load方法
func (u *JsonUserDb) Load() error {

	userdecoder := json.NewDecoder(Filefd)
	err := userdecoder.Decode(&u.UserSlice)
	if err != nil {
		return err
	}
	return nil
}

//写
func (u *JsonUserDb) Sync() error {
	//压入切片队列
	queue = append(queue,u.UserSlice)
	return nil
}

// 滚动存储
func (u *JsonUserDb) RotateSave() error {
	l := len(queue)
	if l <= QueueLen{
		for i:=0; i< l; i++{
			if i == l-1 {
				write := json.NewEncoder(Filefd)
				err := write.Encode(queue[i])
				if err != nil {
					return fmt.Errorf("写入失败:%s",err)
				}
				return nil
			}
			name := path.Dir(Savepath)
			filename := path.Join(name,fmt.Sprintf("user.json.%d",i))
			file,err := os.OpenFile(filename,os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
			if err != nil {
				return fmt.Errorf("创建失败:%s",err)
			}
			write := json.NewEncoder(file)
			err = write.Encode(queue[i])
			if err != nil {
				return fmt.Errorf("写入失败:%s",err)
			}
			_ = file.Close()
		}
		return nil
	}
	if l > QueueLen {
		for k,v := range queue[l-QueueLen:]{
			if k == QueueLen-1 {
				write := json.NewEncoder(Filefd)
				err := write.Encode(v)
				if err != nil {
					return fmt.Errorf("写入失败:%s", err)
				}
				continue
			}
			name := path.Dir(Savepath)
			filename := path.Join(name,fmt.Sprintf("user.json.%d",k))
			file,err := os.OpenFile(filename,os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
			if err != nil {
				return fmt.Errorf("创建失败:%s",err)
			}
			write := json.NewEncoder(file)
			err = write.Encode(v)
			if err != nil {
				return fmt.Errorf("写入失败:%s",err)
			}
			_ = file.Close()

		}
		return nil

	}
	return fmt.Errorf("未知错误")

}