package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/google/uuid"

	"github.com/imroc/req"
)

func main() {
	// 1. 配置
	// 2. UUID
	// 检查文件, 无, 保存
	uid := strings.ReplaceAll(uuid.New().String(), "-", "")
	fmt.Println(uid)

	// 3. 通信(http)

	client := req.New()
	agentAddr := ""
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if strings.Index(addr.String(), ":") > 0 {
			continue
		}

		nodes := strings.SplitN(addr.String(), "/", 2)
		if len(nodes) != 2 {
			continue
		}
		agentAddr = nodes[0]

	}

	hostname, _ := os.Hostname()

	agent := map[string]string{
		"UUID":     uid,
		"hostname": hostname,
		"addr":     agentAddr,
	}
	response, err := client.Post("http://localhost:8888/agent/register", req.BodyJSON(agent))
	fmt.Println(err)
	fmt.Println(response.ToString())

	// 4. 更新配置 yaml
	// 5. 应用配置
	//    a. promtool check
	//    b. 覆盖文件
	//    c. 热加载
}
