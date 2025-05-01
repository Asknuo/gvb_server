package flag

import sys_flag "flag"

type Option struct {
	DB bool
}

//parse解析命令行参数
func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")
	sys_flag.Parse()
	return Option{
		DB: *db,
	}
}

//是否运行web项目
func IsWebStop(option Option) bool {
	if option.DB {
		return true
	}
	return false
}

//根据命令执行函数
func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
	}
}
