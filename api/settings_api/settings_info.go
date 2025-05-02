package settings_api

import (
	"gvb_server/global"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.OKWithData(global.Config.SiteInfo, c)
}
