package database

import (
	"database/sql"
	"go-server-template/pkg/logger"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type GormModel struct {
	Id int64 `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
}

var (
	db    *gorm.DB
	sqlDB *sql.DB
)

func OpenDB(dsn string, logPath string, maxIdleConns, maxOpenConns int, models ...interface{}) (err error) {
	config := &gorm.Config{}

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Errorf("opens database failed: %s", err.Error())
		return
	}

	config.Logger = gormLogger.New(log.New(file, "\r\n", log.LstdFlags), gormLogger.Config{
		SlowThreshold: time.Second,
		Colorful:      true,
		LogLevel:      gormLogger.Info,
	})

	if config.NamingStrategy == nil {
		config.NamingStrategy = schema.NamingStrategy{
			TablePrefix:   "template_",
			SingularTable: true,
		}
	}

	if db, err = gorm.Open(mysql.Open(dsn), config); err != nil {
		logger.Errorf("opens database failed: %s", err.Error())
		return
	}

	if sqlDB, err = db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(maxIdleConns)
		sqlDB.SetMaxOpenConns(maxOpenConns)
	} else {
		logger.Error(err)
	}

	if err = db.AutoMigrate(models...); nil != err {
		logger.Errorf("auto migrate tables failed: %s", err.Error())
	}
	return
}

// 获取数据库链接
func DB() *gorm.DB {
	return db
}

// 关闭连接
func CloseDB() {
	if sqlDB == nil {
		return
	}
	if err := sqlDB.Close(); nil != err {
		logger.Errorf("Disconnect from database failed: %s", err.Error())
	}
}
