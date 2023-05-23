package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

var err error

/**
* 模块名: db.go
* 代码描述: 初始化数据库
* 作者:lizibin
* 创建时间:2022/12/23 14:38:35
 */
func InitDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, // 级别
			Colorful:      true,
		})
	dsn := `root:123456@tcp(localhost:3306)/golb?charset=utf8mb4&parseTime=True&loc=Local`
	// dsn := viper.GetString("mysql.user") + ":" +
	// 	viper.GetString("mysql.password") +
	// 	"@tcp(" + viper.GetString("mysql.databaseIP") + ": " +
	// 	viper.GetString("mysql.databasePort") + ")/" +
	// 	viper.GetString("mysql.databaseName") + "?charset=utf8mb4&parseTime=True&loc=Local"
	// fmt.Println(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		fmt.Println(err)
	}
}
