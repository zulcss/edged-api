package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func AddSiteGroup(rg *gin.RouterGroup) {
	site := rg.Group("/site")

	site.GET("/", func(c *gin.Context) {
                c.JSON(http.StatusOK, "ok")
        })
}
