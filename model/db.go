package model

import (
	"Devops/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	db *gorm.DB
	err error
)

func InitDb() {
	db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
		)))
	if err != nil {
		fmt.Printf("连接数据库失败，请检查参数%s", err)
	}

	db.AutoMigrate(&User{}, &Article{}, &Category{})
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)


}
