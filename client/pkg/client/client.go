package client

import (
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/zulcss/edged/shared/constants"
)

type Client struct {
	*req.Client
	Host string
}

func NewClient(host string) *Client {
	return &Client{
		Client: req.C(),
		Host: fmt.Sprintf("http://%s:%s%s",host, constants.Port, constants.EndPoint),
	}
}