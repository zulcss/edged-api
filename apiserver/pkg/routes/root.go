package routes

import (
	"net/http"
        "github.com/gin-gonic/gin"
)

func AddRootGroup(rg *gin.RouterGroup) {
	root := rg.Group("/")

	// Check for server is up
        root.GET("/health", func(c *gin.Context) {
                c.JSON(http.StatusOK, "ok")
        })
}
