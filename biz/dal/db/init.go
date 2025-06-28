package db

import (
	"Hospital/pkg/utils"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func Init() {
	var err error
	dsn, err := utils.GetMysqlDSN()
	if err != nil {
		panic(err)
	}
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	hlog.Infof("db init suceess")
	if err != nil {
		panic(err)
	}
	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(20 * time.Second)
}
