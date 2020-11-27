package fileserver

import (
	"fmt"
	"net"
	"os"
	"path"
	"remote_fileserver/data_mod"
	pubfunc "remote_fileserver/pubrw"
	"strings"
)

func ls(cmd data_mod.Cmd, conn net.Conn) {
	var sdata data_mod.Data
	dir := path.Join(Basedir, cmd.Arg)
	f, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
		sdata.Data = []byte(fmt.Sprintf("%s 文件或目录不存在.", cmd.Arg))
	} else {
		fd, err := f.Readdirnames(0)
		if err != nil {
			fmt.Printf("%s,read error\n", err)
			sdata.Data = []byte(fmt.Sprintf("%s\n", err))
		} else {
			sdata.Data = []byte(strings.Join(fd, " "))
			sdata.Lenght = len(sdata.Data)

		}
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
