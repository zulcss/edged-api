package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"
	"github.com/zulcss/edged/apiserver/pkg/constants"
)


func Run() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	GetRoutes(router)
	router.Run(fmt.Sprintf("%s:%s", viper.GetString("host.server"), constants.Port))
}

func GetRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")

	// Root
	AddRootGroup(v1)
	// Site
	AddSiteGroup(v1)
}
