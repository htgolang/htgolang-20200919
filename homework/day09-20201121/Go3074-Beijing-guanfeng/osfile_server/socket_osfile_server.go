package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

// 列出目录下文件
func CmdList(dirname string) (string, error) {
	fileinfo, err := ioutil.ReadDir(dirname)
	var filetype string
	var result string
	for _, v := range fileinfo {
		if v.IsDir() {
			filetype = "dir"
		} else {
			filetype = "f"
		}
		fmt.Printf("%s[%s]\n", v.Name(), filetype)
		a := v.Name() + "[" + filetype + "]" + "\n"
		result += a
	}
	return result, err
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

func write(conn net.Conn, server *Server) error {
	ctx, err := json.Marshal(server)
	length := len(ctx)
	all := length
	var i = 0
	var a = 0
	const maxlength = 99999
	for {
		if length > maxlength {
			_, err = conn.Write([]byte(strconv.Itoa(maxlength)))
			if err != nil {
				fmt.Println(err)
				return err
			}
			//剩余读取长度
			length -= maxlength
			fmt.Println(length)
			//本次读取长度
			a += maxlength
			_, err = conn.Write(ctx[i:a])
			//当前位置
			i += maxlength
			if err != nil {
				fmt.Println(err)
				return err
			}
		} else {
			_, err = conn.Write([]byte(fmt.Sprintf("%05d", length)))
			if err != nil {
				fmt.Println(err)
				return err
			}
			_, err = conn.Write(ctx[i:all])
			if err != nil {
				fmt.Println(err)
				return err
			}
			break
		}
	}
	return nil
}

func read(conn net.Conn) (Client, error) {
	lengthBytes := make([]byte, 5)
	_, err := conn.Read(lengthBytes)
	if err != nil {
		fmt.Println(err)
		return Client{}, err
	}
	length, err := strconv.Atoi(string(lengthBytes))
	if err != nil {
		fmt.Println(err)
		return Client{}, err
	}
	ctx := make([]byte, length)
	var client Client
	_, err = conn.Read(ctx)
	if err != nil {
		return Client{}, err
	}
	err = json.Unmarshal(ctx, &client)
	if err != nil {
		return Client{}, err
	}
	return client, nil
}

func main() {
	addr := "0.0.0.0:8888"
	protocol := "tcp"

	listener, err := net.Listen(protocol, addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil { // 服务器关闭
			fmt.Println(conn)
			continue
		}
		go func() {
			fmt.Println("客户端连接成功: ", conn.RemoteAddr())
			client, err := read(conn)
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				return
			}
			var server *Server
			if client.Cmd == "ls" {
				w, err := CmdList(client.Path)
				if err != nil {
					server = &Server{444, err.Error()}
				} else {
					server = &Server{200, w}
				}
				err = write(conn, server)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			if client.Cmd == "get" {
				w, err := ReadFile(client.Src)
				if err != nil {
					if err != io.EOF {
						server = &Server{444, err.Error()}
					}
				} else {
					server = &Server{200, string(w)}
				}
				err = write(conn, server)
				if err != nil {
					fmt.Println(err)
					return
				}

			}
			if client.Cmd == "put" {
				txt := []byte(client.Content)
				err := WriteFile(client.Dest, txt)
				if err != nil {
					server = &Server{444, err.Error()}
				} else {
					server = &Server{200, "success"}
				}
				err = write(conn, server)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			if client.Cmd == "rm" {
				err := os.Remove(client.Path)
				if err != nil {
					fmt.Println(err)
					server = &Server{444, err.Error()}
				} else {
					server = &Server{200, "remove success"}
				}
				err = write(conn, server)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			err = conn.Close()
			fmt.Println("客户端退出: ", conn.RemoteAddr())
		}()
	}
	//listener.Close()
}
