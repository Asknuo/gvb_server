package main

import (
	"gvb_server/core"
	"gvb_server/flag"
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
	//命令行参数绑定
	op := flag.Parse()
	if flag.IsWebStop(op) {
		flag.SwitchOption(op)
		return
	}
	//初始化路由
	router := routers.InitRouters()
	addr := global.Config.System.Addr()
	global.Log.Infoln("服务启动成功,地址为:", addr)
	//启动服务
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatal(err.Error())
	}
}
