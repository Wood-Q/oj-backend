package databases
import (
	"fmt"
	"time"
	"OJ/pkg/global"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := "host=localhost user=postgres password=246266262 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed to connect to database:%v", err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		fmt.Printf("failed to connect to database:%v", err)

	}
	fmt.Println("Database connection successfully established")
	global.Db = db
}
