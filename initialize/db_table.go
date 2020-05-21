package initialize

import (
	"gin-example/global"
	"gin-example/models"
	//"gin-example/models"
)

//注册数据库表专用(程序启动的时候会注册表，如果又改动就迁移)
func DBTables() {
	db := global.GVA_DB
	db.AutoMigrate(model.User{})
	global.GVA_LOG.Debug("register table success")
}
