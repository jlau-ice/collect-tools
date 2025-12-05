package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jlau-ice/collect/internal/config"
	"github.com/jlau-ice/collect/internal/container"
	"github.com/jlau-ice/collect/internal/runtime"
	"gorm.io/gorm"
)

func main() {
	// 构建依赖注入容器
	c := container.BuildContainer(runtime.GetContainer())
	// 通过依赖注入获取所有需要的服务
	err := c.Invoke(func(
		cfg *config.Config,
		db *gorm.DB,
		router *gin.Engine,
	) error {
		// 确保数据库连接在程序结束时关闭
		defer func() {
			sqlDB, err := db.DB()
			if err != nil {
				log.Printf("获取数据库连接失败: %v", err)
				return
			}
			if err := sqlDB.Close(); err != nil {
				log.Printf("关闭数据库连接失败: %v", err)
			}
		}()
		// 启动服务器
		port := ":" + cfg.Server.Port
		log.Printf("服务器启动在端口 %s", port)
		return router.Run(port)
	})
	if err != nil {
		log.Fatalf("应用启动失败: %v", err)
	}
}
