package settings_api

import (
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.OK(map[string]string{}, "xxx", c)
}
