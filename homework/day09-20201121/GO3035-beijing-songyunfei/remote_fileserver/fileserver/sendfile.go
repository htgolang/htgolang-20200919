package fileserver

import (
	"fmt"
	"net"
	"path"
	"remote_fileserver/data_mod"
	pubfunc "remote_fileserver/pubrw"
)

func sendfile(cmd data_mod.Cmd, conn net.Conn) error {
	filepath := path.Join(Basedir, cmd.Arg)
	cmd.Arg = filepath
	err := pubfunc.Sendfile(conn, cmd)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}
