package container

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jlau-ice/collect/internal/config"
	"github.com/jlau-ice/collect/internal/models"
	"github.com/jlau-ice/collect/internal/router"
	"go.uber.org/dig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// BuildContainer 构建依赖注入容器
// 注册所有需要的依赖：配置、数据库、路由等
func BuildContainer(container *dig.Container) *dig.Container {
	// 注册配置加载器
	must(container.Provide(config.LoadConfig))
	// 注册数据库初始化器
	must(container.Provide(initDatabase))
	// 注册 Gin 引擎创建器（包含路由设置）
	must(container.Provide(initGinEngine))
	return container
}

// must 辅助函数，如果错误则 panic
func must(err error) {
	if err != nil {
		panic(err)
	}
}

// initDatabase 初始化数据库连接
// 依赖注入：需要 *config.Config
// 返回：*gorm.DB
func initDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.Database.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// GORM 日志配置
		//Logger: logger.Default.LogMode(logger.Info),
		// ❗ 注意：如果您在 config 包中设置了 Schema，
		// 这里 GORM 默认就会使用该 Schema 进行操作。
	})
	if err != nil {
		return nil, fmt.Errorf("数据库连接失败: %w", err)
	}
	log.Println("数据库连接成功")
	// 自动迁移
	err = db.AutoMigrate(
		&models.Department{},
		&models.User{},
		&models.Task{},
		&models.Upload{},
	)
	if err != nil {
		return nil, fmt.Errorf("数据库表迁移失败: %w", err)
	}
	log.Println("数据库表迁移完成")
	// 设置连接池参数
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("获取底层数据库连接失败: %w", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Duration(10) * time.Minute)
	return db, nil
}

// initGinEngine 初始化 Gin 引擎并设置路由
// 返回：*gin.Engine
func initGinEngine() (*gin.Engine, error) {
	// 设置Gin模式
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	engine := gin.Default()
	// 设置路由
	router.SetupRoutes(engine)
	return engine, nil
}
