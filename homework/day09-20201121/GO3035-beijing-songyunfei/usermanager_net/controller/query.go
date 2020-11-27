package controller

import (
	"bytes"
	"fmt"
	"net"
	"usermanager_net/users"
	"usermanager_net/userutils"
)

func queryuser(c userutils.Cmd,udb users.Mydb,conn net.Conn) {
	str := c.Data
	u,ok := udb.QueryUser(str)
	var mes = userutils.Msg{
		Status: 200,
		Ack:    true,
		Data:   "",
	}
	if ok {
		header := []string{"用户名","地址","联系电话","生日"}
		data := [][]string{{u.Name,u.Addr,u.Tel,u.Birthday.Format("2006-1-2")}}
		var buf = bytes.NewBufferString("")
		userutils.Gentable(header,data,buf)
		mes.Data = buf.String()

	}else {
		mes.Data = "未找到"
	}
	err := userutils.Sendmes(conn,mes)
	if err !=nil {
		fmt.Println(err)
		return
	}
}
