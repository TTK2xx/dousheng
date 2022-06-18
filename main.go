package main

import (
	"dousheng/config"
	"dousheng/database"
	"dousheng/logger"
	"dousheng/router"
	"dousheng/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置文件
	cfg, err := config.ParseConfig("./config/init.json")
	if err != nil {
		panic("Configuration file parse error! " + err.Error())
	}
	// 初始化日志
	if err := logger.InitLogger(cfg); err != nil {
		fmt.Printf("init logger failed! \n", err.Error())
		return
	}
	// 连接Mysql
	if err := database.InitMySQL(cfg); err != nil {
		fmt.Printf("MySQL connect error! \n" + err.Error())
		return
	}
	sqlDB, _ := database.MySQLDB.DB()
	defer sqlDB.Close()
	// 连接redis
	if err := database.InitRedisClient(cfg); err != nil {
		fmt.Printf("Reids connect error! \n" + err.Error())
		return
	}
	// 加载 雪花 id
	if err := utils.Init(); err != nil {
		fmt.Printf("init redis failed, err: \n" + err.Error())
		return
	}

	// 初始化路由
	r := gin.Default()
	// 注册zap日志相关中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	router.InitRouter(r)

	r.Run(cfg.AppPort)
}
