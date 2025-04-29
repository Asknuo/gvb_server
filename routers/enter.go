package routers

import (
	"gvb_server/global"

	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	settingsRouter(router)
	return router
}
