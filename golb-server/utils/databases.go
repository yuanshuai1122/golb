package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

var err error

func InitDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, // 级别
			Colorful:      true,
		})
	dsn := `root:hh%Gy&rTYM@oepc6@tcp(bj-cynosdbmysql-grp-o9u7ldum.sql.tencentcdb.com:21457)/golb_local?charset=utf8mb4&parseTime=True&loc=Local`
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		fmt.Println(err)
	}
}
