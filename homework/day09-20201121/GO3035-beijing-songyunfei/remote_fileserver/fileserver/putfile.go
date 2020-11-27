package fileserver

import (
	"fmt"
	"net"
	"path"
	"remote_fileserver/data_mod"
	pubfunc "remote_fileserver/pubrw"
)

func putfile(cmd data_mod.Cmd, conn net.Conn) error {
	newpath := path.Join(Basedir, cmd.Arg)
	cmd.Arg = newpath
	err := pubfunc.Getfile(cmd, conn)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
