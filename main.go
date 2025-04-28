package main

import (
	"gvb_server/core"
	"gvb_server/globel"
)

func main() {
	//读取配置文件
	core.InitConfig()
	//连接数据库
	globel.DB = core.InitGorm()
}
