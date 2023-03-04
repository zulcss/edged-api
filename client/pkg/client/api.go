package client

import (
	"fmt"
)

type NodeInfo struct {
	Status string `json:"status"`
    Hostname string `json:"hostname"`
    Version string `json:"version"`
}

func (c *Client) Status() (*NodeInfo, error) {	
	node := &NodeInfo{}

	_, err := c.R().
			SetSuccessResult(&node).
			Get(fmt.Sprintf("%s/health", c.Host))
	
	return node, err
}