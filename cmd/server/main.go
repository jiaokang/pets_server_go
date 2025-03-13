package main

import (
	"github.com/gin-gonic/gin"
	"jk.com/pets/config"
	"jk.com/pets/internal/database"
	"jk.com/pets/internal/handlers"
)

func main() {
	// 加载配置文件
	cfg, err := config.LoadConfig("config.yml")
	if err != nil {
		panic("failed to load config")
	}

	// 初始化数据库
	database.InitDB(cfg)

	r := gin.Default()

	handlers.RegisterAuthRoutes(r)

	r.Run(":9001")
}
