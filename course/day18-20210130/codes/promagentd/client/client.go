package client

import (
	"fmt"
	"promagentd/config"

	"github.com/imroc/req"
	"github.com/sirupsen/logrus"
)

type Client struct {
	option  *config.Option
	request *req.Req
}

func NewClient(option *config.Option) *Client {
	return &Client{
		option:  option,
		request: req.New(),
	}
}

func (c *Client) Register(event map[string]interface{}) {
	path := fmt.Sprintf("%s/agent/register/", c.option.Server)
	response, err := c.request.Post(path, req.BodyJSON(event))
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Debug(response.ToString())
}

func (c *Client) Heartbeat(event map[string]interface{}) {
	path := fmt.Sprintf("%s/agent/heartbeat/", c.option.Server)
	response, err := c.request.Get(path, req.Param(event))
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Debug(response.ToString())
}

func (c *Client) Config(event map[string]interface{}) (string, error) {
	path := fmt.Sprintf("%s/agent/config/", c.option.Server)
	response, err := c.request.Get(path, req.Param(event))
	if err != nil {
		logrus.Error(err)
		return "", nil
	}
	return response.ToString()
}
