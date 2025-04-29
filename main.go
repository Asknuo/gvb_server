package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/routers"
)

func main() {
	//读取配置文件
	core.InitConfig()
	//初始化日志
	global.Log = core.InitLog()
	//连接数据库
	global.DB = core.InitGorm()
	//初始化路由
	router := routers.InitRouters()
	addr := global.Config.System.Addr()
	global.Log.Infoln("服务启动成功,地址为:", addr)
	//启动服务
	router.Run(addr)
}
