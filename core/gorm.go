package core

import (
	"gvb_server/global"
	"log"
	"time"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		log.Println("没有配置MySQL,取消连接")
		return nil
	}
	dsn := global.Config.Mysql.DSN()
	var mysqllogger logger.Interface
	if global.Config.System.Env == "debug" {
		//开发环境显示所有mysql
		mysqllogger = logger.Default.LogMode(logger.Info)
	} else {
		//只打印错误mysql
		mysqllogger = logger.Default.LogMode(logger.Error)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqllogger,
	})
	if err != nil {
		log.Println("连接MySQL失败", err)
		return nil
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour * 4)
	return db
}
