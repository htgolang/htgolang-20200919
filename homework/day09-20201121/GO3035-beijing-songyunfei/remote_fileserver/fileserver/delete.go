package fileserver

import (
	"fmt"
	"net"
	"os"
	"path"
	"remote_fileserver/data_mod"
	pubfunc "remote_fileserver/pubrw"
)

func deletefile(cmd data_mod.Cmd, conn net.Conn) {
	var sdata data_mod.Data
	dir := path.Join(Basedir, cmd.Arg)
	err := os.Remove(dir)
	if err != nil {
		fmt.Println(err)
		sdata.Data = []byte(fmt.Sprintf("%s 文件或目录不存在.", cmd.Arg))
	} else {
		sdata.Data = []byte(fmt.Sprintf("%s 已删除.", cmd.Arg))

	}
	send, err := pubfunc.MarshalData(&sdata)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = pubfunc.Senddata(conn, send); err != nil {
		fmt.Println(err)
		return
	}
}
