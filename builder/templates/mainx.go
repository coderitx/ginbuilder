package templates

var MainTemplate = `package main

import (
	"{{.pkgname}}/apps/routers"
	"{{.pkgname}}/common/logx"
	"{{.pkgname}}/config"
	"{{.pkgname}}/global"
	//"{{.pkgname}}/internal"
	"go.uber.org/zap"
)

func init() {
	c, err := config.LoadConfig(global.ConfigPath)
	if err != nil {
		panic(err)
	}
	global.GlobalC = c
	// 如果需要使用数据库取消注视即可
	//internal.InitDB(global.GlobalC.Mysql)
	//internal.InitRedis(global.GlobalC.Redis)
	logx.InitLogger(global.GlobalC.Log)
	zap.S().Infoln("--------所有配置初始化完成---------")	
}


func main(){
	zap.S().Infoln("start server")
	routers.InitApiRouter()
}
`
