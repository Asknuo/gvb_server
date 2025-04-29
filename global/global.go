package global

import (
	"gvb_server/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	// Config 全局配置变量
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)
