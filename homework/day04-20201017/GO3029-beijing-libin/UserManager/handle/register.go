package handle

import (
	"UserManager/model"
	"fmt"
)

func Register(username string, passwd string, userinfo *model.UserInfo, userSlice *model.UserSlice) {

	//生成一个新的userinfo
	userinfo = &model.UserInfo{
		Name:   username,
		Pwd:    passwd,
		Enable: true,
	}
	model.MakeMd5(passwd, userinfo)

	//然后将这个userinfo添加到userslice中
	userSlice.User = append(userSlice.User, userinfo)
	fmt.Printf("Thanks for %s registering", username)

}
