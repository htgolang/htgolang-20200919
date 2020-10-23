package functions

import (
	"fmt"
	"strconv"
	"usermanager/users"
)

//通过ID查找用户
func findByid(id int) (index int,err error) {
	for k,v := range users.Users {
		uid,err := strconv.Atoi(v["id"])
		if err != nil {
			_ = fmt.Errorf("id转换错误:%s",err)
			return -1, err
		}
		if uid == id {
			return k,nil
		}
	}
	return -1, fmt.Errorf("未找到")

}
