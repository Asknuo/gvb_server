package globel

import (
	"gvb_server/config"

	"gorm.io/gorm"
)

var (
	// Config 全局配置变量
	Config *config.Config
	DB     *gorm.DB
)
