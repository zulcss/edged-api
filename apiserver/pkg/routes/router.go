package routes

import (
	"github.com/gin-gonic/gin"
)


func Run() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	GetRoutes(router)
	router.Run(":8080")
}

func GetRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")

	// Root
	AddRootGroup(v1)
}
