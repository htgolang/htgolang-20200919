package config

import (
	"fmt"
	"net"
	"os"
	"promagentd/utils"
	"strings"

	"github.com/google/uuid"
)

type Option struct {
	Server   string
	Conf     string
	UUID     string
	Hostname string
	Addr     string
	Promtool string
}

func GetUUID() string {
	path := "promeagentd.uuid"
	if agentID := utils.ReadFile(path); agentID != "" {
		return agentID
	}
	agentID := strings.ReplaceAll(uuid.New().String(), "-", "")

	utils.WriteFile(path, agentID)
	return agentID
}

func GetAddr() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if strings.Index(addr.String(), ":") >= 0 {
			continue
		}
		nodes := strings.SplitN(addr.String(), "/", 2)
		if len(nodes) != 2 {
			continue
		}
		return nodes[0], nil
	}
	return "", fmt.Errorf("no net addr")

}

func NewOption(server string, conf string) (*Option, error) {
	uuid := GetUUID()
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	addr, err := GetAddr()
	if err != nil {
		return nil, err
	}
	return &Option{
		Server:   server,
		Conf:     conf,
		UUID:     uuid,
		Hostname: hostname,
		Addr:     addr,
		Promtool: "/opt/prometheus/prometheus/promtool",
	}, nil
}
