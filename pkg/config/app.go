package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	dbmysql, err := gorm.Open("mysql", "root:123456@(127.0.0.1)/gin_info?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("数据库初始化失败")
	}
	db = dbmysql
}

func GetConnect() *gorm.DB {
	return db
}
