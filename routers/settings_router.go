package routers

import (
	"gvb_server/api"

	"github.com/gin-gonic/gin"
)

func settingsRouter(router *gin.Engine) {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("", settingsApi.SettingsInfoView)
}
