package main

import (
	"go-logistics/model"
	"go-logistics/routes"
	"go-logistics/utils"
)

// @title Logistics Query Api
// @version 1.0
// @description 聚合物流服务

// @contact.name   martin
// @contact.url
// @contact.email

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func main() {
	utils.InitConfig() // 加载配置文件
	model.InitDB()     // 初始化数据库
	routes.InitRouter()
}
