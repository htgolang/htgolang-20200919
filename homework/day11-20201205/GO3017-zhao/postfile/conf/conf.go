package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Conf struct {
	DriverName string `json:"driverName"`
	Host       string `json:"host"`
	Post       string `json:"port"`
}

func LoadConf() (*Conf, error) {
	confBytes, err := ioutil.ReadFile("default.conf")
	if err != nil {
		fmt.Println(err, "Open conf false")
		return nil, err
	}
	conf := new(Conf)
	err = json.Unmarshal(confBytes, conf)
	if err != nil {
		fmt.Println(err, "json Unmarshal false")
		return nil, err
	}
	return conf, nil
}
