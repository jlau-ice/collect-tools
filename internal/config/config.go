package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config 应用配置
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Upload   UploadConfig   `mapstructure:"upload"`
	MinIo    MinIoConfig    `mapstructure:"minio"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string
	Mode string // debug, release, test
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	TimeZone string
	Schema   string
}

// UploadConfig 上传配置
type UploadConfig struct {
	BasePath string // 文件存储基础路径
	MaxSize  int64  // 最大文件大小（字节）
}

type MinIoConfig struct {
	EndPoint   string
	BucketName string
	AccessKey  string
	SecretKey  string
}

// LoadConfig 加载配置
// 返回 *Config 以便依赖注入容器使用
func LoadConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")   // 配置文件名称(不带扩展名)
	v.SetConfigType("yaml")     // 配置文件类型
	v.AddConfigPath("./config") // config子目录
	v.AddConfigPath(".")        // 当前目录

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("加载配置文件失败: %w", err)
	}
	// 支持环境变量覆盖配置（可选）
	v.AutomaticEnv()

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("配置绑定失败: %w", err)
	}
	return &cfg, nil
}

// GetDSN 获取数据库连接字符串
func (c *DatabaseConfig) GetDSN() string {
	// 使用标准的 key=value 格式，不需要额外的 'options=' 关键字包裹。
	// search_path 应该直接作为 DSN 的一个参数。
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s search_path=%s",
		c.Host,
		c.User,
		c.Password,
		c.DBName,
		c.Port,
		c.SSLMode,
		c.TimeZone,
		c.Schema,
	)
}
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
