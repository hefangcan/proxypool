package database

import (
	"github.com/bh-qt/proxypool/log"
	"os"

	"github.com/bh-qt/proxypool/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func connect() (err error) {
	// mysql url
	dsn := "hefangcan:hefangcan@tcp(192.168.2.251:3306)/proxy?charset=utf8mb4&parseTime=True&loc=Local"
	if url := config.Config.DatabaseUrl; url != "" {
		dsn = url
	}
	if url := os.Getenv("DATABASE_URL"); url != "" {
		dsn = url
	}
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err == nil {
		log.Infoln("database: successfully connected to: %s", DB.Name())
	} else {
		DB = nil
		log.Warnln("database connection info: %s \n\t\tUse cache to store proxies", err.Error())
	}
	return
}
