package settings_api

import (
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr config.SiteInfo
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArguementError, c)
	}
	global.Config.SiteInfo = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OKWith(c)
}
