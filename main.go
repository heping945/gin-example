package main

import (
	"gin-example/core"
	"gin-example/global"
	"gin-example/initialize"
)

func main() {

	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		initialize.Mysql()
	//case "sqlite":
	//	initialize.Sqlite()  // sqlite需要gcc支持 windows用户需要自行安装gcc 如需使用打开注释即可
	default:
		initialize.Mysql()
	}
	initialize.DBTables()
	// 程序结束前关闭数据库链接
	defer global.GVA_DB.Close()

	core.RunServer()
}
