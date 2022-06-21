package main

import (
	"go-logistics/global"
	"go-logistics/model"
	"go-logistics/routes"
	"go-logistics/utils"
)

// @title Logistics Open Api
// @version 1.0
// @description 聚合物流服务

// @contact.name   Mzihao
// @contact.url
// @contact.email pimpkin_mak@163.com

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func main() {
	utils.InitConfig() // 加载配置文件
	model.InitDB()     // 初始化数据库
	global.InitHttpClient()
	routes.InitRouter()
}
