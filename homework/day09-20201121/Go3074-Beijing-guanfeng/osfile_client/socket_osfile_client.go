package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

type Server struct {
	Status  int
	Content string
}

type Client struct {
	Cmd     string
	Src     string
	Dest    string
	Path    string
	Content string
}

func write(conn net.Conn, client *Client) error {
	ctx, err := json.Marshal(client)
	length := len(ctx)
	_, err = conn.Write([]byte(fmt.Sprintf("%05d", length)))
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = conn.Write(ctx)
	if err != nil {
		return err
	}
	return nil
}

func read(conn net.Conn) (Server, error) {
	var server Server
	var buffer bytes.Buffer
	lengthBytes := make([]byte, 5)
	const maxlength = 99999

	for {
		_, err := conn.Read(lengthBytes)
		if err != nil {
			if err != io.EOF {

				return Server{}, err
			}
		}
		fmt.Println(string(lengthBytes))
		length, err := strconv.Atoi(string(lengthBytes))
		if err != nil {
			fmt.Println("line 57", err)
			return Server{}, err
		}
		if length == maxlength {
			ctx := make([]byte, maxlength)
			_, err = conn.Read(ctx)
			buffer.Write(ctx)
			if err != nil {
				return Server{}, err
			}
		}
		if length < maxlength && length != 0 {
			ctx := make([]byte, length)
			_, err = conn.Read(ctx)
			buffer.Write(ctx)
			if err != nil {
				if err != io.EOF {
					return Server{}, err
				}
			}
			break
		}
		if length == 0 {
			break
		}
	}

	b3 := buffer.Bytes()

	//fmt.Println(len(b3))
	err := json.Unmarshal(b3, &server)
	if err != nil {
		return Server{}, err
	}
	return server, nil
}

func ReadFile(f string) ([]byte, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	context := make([]byte, 256)
	var all []byte
	for {
		n, err := file.Read(context)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		aa := context[:n]
		all = append(all, aa...)
	}
	return all, nil
}

func WriteFile(f string, all []byte) error {
	file, err := os.OpenFile(f, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0660)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(all)
	return err
}

func CmdList(conn net.Conn, client *Client) {
	err := write(conn, client)
	if err != nil {
		fmt.Println(err)
	}
	server, err := read(conn)
	if err != nil {
		if err != io.EOF {
			fmt.Println(err)
			return
		}
		return
	}

	status := server.Status
	txt := server.Content
	fmt.Printf("status: %d\n%s", status, txt)

}

func CmdGet(conn net.Conn, client *Client) {
	fmt.Println(client)
	err := write(conn, client)
	if err != nil {
		fmt.Println(err)
	}
	server, err := read(conn)

	if err != nil {
		if err != io.EOF {
			fmt.Println(err)
			return
		}
		return
	}
	var txt []byte
	status := server.Status
	txt = []byte(server.Content)
	err = WriteFile(client.Dest, txt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("status: %d", status)
}

func CmdPut(conn net.Conn, client *Client) {
	fmt.Println(client)
	txt, err := ReadFile(client.Src)
	client.Content = string(txt)
	if err != nil {
		fmt.Println(err)
	}
	err = write(conn, client)
	if err != nil {
		fmt.Println(err)
	}
	server, err := read(conn)

	if err != nil {
		if err != io.EOF {
			fmt.Println(err)
			return
		}
		return
	}
	fmt.Printf("status: %d %s", server.Status, server.Content)
}

func CmdRm(conn net.Conn, client *Client) {
	err := write(conn, client)
	if err != nil {
		fmt.Println(err)
	}
	server, err := read(conn)
	if err != nil {
		if err != io.EOF {
			fmt.Println(err)
			return
		}
		return
	}

	status := server.Status
	txt := server.Content
	fmt.Printf("status: %d\n%s", status, txt)
}

func command() (string, string, string, string) {
	var cmd string
	var path string
	var src string
	var dest string
	var help bool
	flag.StringVar(&cmd, "cmd", "", "cmd")
	flag.StringVar(&path, "path", "", "path")
	flag.StringVar(&src, "s", "", "src")
	flag.StringVar(&dest, "d", "", "dest")
	flag.BoolVar(&help, "h", false, "help")
	flag.Usage = func() {
		fmt.Println("Usage:")
		flag.PrintDefaults()
	}
	flag.Parse()
	if help || cmd != "ls" && cmd != "put" && cmd != "get" && cmd != "rm" {
		flag.Usage()
		os.Exit(0)
	}
	return cmd, path, src, dest
}

func Connection() (net.Conn, error) {
	addr := "127.0.0.1:8888"
	protocol := "tcp"
	conn, err := net.Dial(protocol, addr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return conn, nil
}

func main() {
	cmd, path, src, dest := command()
	client := &Client{cmd, src, dest, path, ""}
	conn, err := Connection()
	if err != nil {
		fmt.Println(err)
		return
	}
	if client.Cmd == "ls" {
		CmdList(conn, client)
	}
	if client.Cmd == "get" {
		CmdGet(conn, client)
	}
	if client.Cmd == "put" {
		CmdPut(conn, client)
	}
	if client.Cmd == "rm" {
		CmdRm(conn, client)
	}
	err = conn.Close()
}
