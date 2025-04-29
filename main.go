package main

import (
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	//读取配置文件
	core.InitConfig()
	//初始化日志
	global.Log = core.InitLog()
	global.Log.Warnln("warn")
	global.Log.Error("error")
	global.Log.Info("info")
	//连接数据库
	global.DB = core.InitGorm()
}
