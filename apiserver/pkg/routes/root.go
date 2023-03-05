package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/zulcss/edged/apiserver/pkg/api"
)

func AddRootGroup(rg *gin.RouterGroup) {
	root := rg.Group("/")
    root.GET("/", api.NodeStatus)
}
