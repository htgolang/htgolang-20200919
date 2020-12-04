package main

import (
	"fmt"
	"os"
	"path/filepath"
	"zhao/server"
)

func main() {
	cmdbin, err := os.Executable()
	if err != nil {
		fmt.Println(err, "[os.Executable]")
		return
	}
	webroot := filepath.Join(filepath.Dir(filepath.Dir(cmdbin)), "webroot")

	server.HttpServer(webroot)

}
