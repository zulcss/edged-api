package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/zulcss/edged/apiserver/pkg/api"
)

func AddRootGroup(rg *gin.RouterGroup) {
	root := rg.Group("/")

	root.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"server": "localhost"})
	})

	// Check for server is up
    	root.GET("/health", api.NodeStatus)
}
