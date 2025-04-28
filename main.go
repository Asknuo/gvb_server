package main

import (
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	//读取配置文件
	core.InitConfig()
	//连接数据库
	global.DB = core.InitGorm()
}
