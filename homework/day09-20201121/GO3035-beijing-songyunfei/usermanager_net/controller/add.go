package controller

import (
	"fmt"
	"net"
	"usermanager_net/users"
	"usermanager_net/userutils"
)

func Add(u userutils.Cmd,udb users.Mydb, conn net.Conn){
	un := u.Userinfo["username"]
	ua := u.Userinfo["addr"]
	ut := u.Userinfo["tel"]
	up := u.Userinfo["passwd"]
	ub := u.Userinfo["brth"]
	if err := udb.Add(un,ua,ut,up,ub); err != nil{
		fmt.Println(err)
	}else {
		if err = udb.Sync(); err == nil{
			var m = userutils.Msg{
				Status: 200,
				Ack:    true,
				Data:   "添加成功.",
			}
			err = userutils.Sendmes(conn,m)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

	}
}
