// 初始化mysql

package core

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"grs-server/global"
	"time"
)

func InitMysql() *gorm.DB {

	if global.Config.Mysql.Host == "" {
		logrus.Warn("未配置mysql，取消gorm连接")
		return nil
	}

	dsn := global.Config.Mysql.Dsn() // Mysql数据库

	var mysqlLogger logger.Interface

	var logLevel = logger.Error
	switch global.Config.Mysql.LogLevel {
	case "info":
		logLevel = logger.Info
	case "warn":
		logLevel = logger.Warn
	}

	mysqlLogger = logger.Default.LogMode(logLevel)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 mysqlLogger,
		SkipDefaultTransaction: true,
	})

	if err != nil {
		logrus.Fatalf(fmt.Sprintf("[%s] mysql 连接失败, err:%s", dsn, err.Error()))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大连接数量
	sqlDB.SetMaxOpenConns(100)              // 最大带宽
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间，不能超过mysql的wait_timeout
	return db
}
