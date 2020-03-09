package main

import (
	"gin_web/config"
	"gin_web/model"
	"gin_web/router"
	"github.com/gin-gonic/gin"
)

func main() {

	//初始化配置
	config.InitCnf("online")
	// 初始化数据库
	model.InitDB()
	defer model.DB.Close()
	// 设置gin开发模式
	gin.SetMode(config.Configs.Key("gin_mode").String())
	// 初始化路由
	r := router.InitRouter()
	// 启动服务器
	r.Run(":5000")
}
