package configs

import (
	"OJ/pkg/global"
	"OJ/platform/databases"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	JWT struct {
		Secret        string
		MinExpires     string
		RefreshKey    string
		RefreshExpire string
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./pkg/configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file:%v", err)
	}

	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	// 初始化 Redis
	//cache.InitRedis()

	// 使用 AppConfig 中的配置初始化数据库
	db, err := databases.InitDB(AppConfig.Database.Host, AppConfig.Database.User, AppConfig.Database.Password, AppConfig.Database.Name, AppConfig.Database.Port)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// 将 db 赋值到 global 包中的 Db 变量（假设 global 包已经正确初始化）
	global.Db = db
}
