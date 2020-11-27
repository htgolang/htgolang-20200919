package pubfunc

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"remote_fileserver/data_mod"
	"strconv"
)

// Readdata 读取数据返回字节切片
func Readdata(conn net.Conn) ([]byte, error) {
	tmplen := make([]byte, 5)
	_, err := conn.Read(tmplen)
	if err != nil {
		return nil,err
	}
	length, err := strconv.Atoi(string(tmplen))
	if err != nil {
		fmt.Println(err)
	}
	data := make([]byte, length)
	_, err = conn.Read(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, nil
}

// Senddata 发送数据返回错误或者nil
func Senddata(conn net.Conn, data []byte) error {

	length := len(data)
	_, err := conn.Write([]byte(fmt.Sprintf("%05d", length)))
	if err != nil {
		return err
	}
	_, err = conn.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// Unmarshalcmd 解析cmd
func Unmarshalcmd(src []byte) (data_mod.Cmd, error) {
	var cmd data_mod.Cmd
	err := json.Unmarshal(src, &cmd)
	if err != nil {
		fmt.Printf("json unmarshal error:%s", err)
		return data_mod.Cmd{}, err
	}
	return cmd, nil
}

// Unmarshaldata 解析data
func Unmarshaldata(src []byte) (data_mod.Data, error) {
	var data data_mod.Data
	err := json.Unmarshal(src, &data)
	if err != nil {
		fmt.Printf("json unmarshal error:%s", err)
		return data_mod.Data{}, err
	}
	return data, nil
}

//MarshalData 反解析data
func MarshalData(data *data_mod.Data) ([]byte, error) {
	bytedata, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return bytedata, nil
}

// Input 输入函数
func Input(str string) data_mod.Cmd {
	var tcmd string
	var arg string
	fmt.Printf("%s", str)
	_, _ = fmt.Scanf("%s %s\n", &tcmd, &arg)
	cmd := data_mod.Cmd{
		Cmd: tcmd,
		Arg: arg,
	}
	return cmd
}

//Sendfile 发送文件
func Sendfile(conn net.Conn, c data_mod.Cmd) error {
	fmt.Println(c)
	f, err := os.Open(c.Arg)
	if err != nil {
		fmt.Println(err)
		return err
	}
	sd, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	err = Senddata(conn, sd)
	if err != nil {
		return err
	}
	ackd, err := Readdata(conn)
	if err != nil {
		return err
	}
	data, err := Unmarshaldata(ackd)
	if err != nil {
		return err
	}
	if data.Ack {
		for {
			var td data_mod.Data
			var b []byte = make([]byte, 1024)
			n, err := f.Read(b)
			if err != nil {
				if err == io.EOF {
					td.Error = "io.EOF"
					sd, err := MarshalData(&td)
					if err != nil {
						fmt.Println(err)
						return err
					}
					err = Senddata(conn, sd)
					fmt.Println("上传结束..")
					return nil
				}
				fmt.Println(err)
			}
			td.Data = b[:n]
			sd, err := MarshalData(&td)
			if err != nil {
				fmt.Println(err)
				return err
			}
			err = Senddata(conn, sd)
			if err != nil {
				fmt.Println(err)
				return err
			}

		}
	}
	return nil

}

// Getfile 下载文件
func Getfile(cmd data_mod.Cmd, conn net.Conn) error {
	var sd data_mod.Data
	filename := cmd.Arg
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	defer f.Close()
	if err != nil {
		sd.Error = fmt.Sprintf("%s create error", cmd.Arg)
		data, err := MarshalData(&sd)
		if err != nil {
			fmt.Println(err)
		}
		err = Senddata(conn, data)
		return err

	}
	ack := data_mod.Data{
		Ack: true,
	}
	ackdata, err := MarshalData(&ack)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("发送ack")
	err = Senddata(conn, ackdata)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("开始接收...")
	for {
		wdata, err := Readdata(conn)
		if err != nil {
			return err
		}
		rdata, err := Unmarshaldata(wdata)
		if err != nil {
			return err
		}
		if rdata.Error == "io.EOF" {
			fmt.Println("接收完成")
			return nil
		}
		_, err = f.Write(rdata.Data)
		if err != nil {
			return err
		}

	}
}
