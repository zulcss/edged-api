package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Node struct {
	Status string `json:"status"`
}

func NodeStatus(c *gin.Context) {
	node := Node{}
	node.Status = "ok"

	c.JSON(http.StatusOK, node)
}
