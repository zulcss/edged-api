package api

import (
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"

	"github.com/zulcss/edged/shared/constants"
)

type Node struct {
	Status string `json:"status"`
	Version string `json:"version"`
	Hostname string `json:"hostname"`
}

func NodeStatus(c *gin.Context) {
	node := Node{}

	node.Status = "ok"
	node.Version = constants.Version
	
	hostname, err := os.Hostname()
	if err != nil {
		log.Println("Failed to get hostname")
	}
	node.Hostname = hostname

	c.JSON(http.StatusOK, node)
}
