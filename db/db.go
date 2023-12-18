package db

import (
	// _ "github.com/go-sql-driver/mysql"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var user, pass, host string

func GetDB() *gorm.DB {
	return db
}

func InitDB() {
	user = os.Getenv("MYSQL_USER")
	pass = os.Getenv("MYSQL_PASSWORD")
	host = os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	db.Logger.LogMode(logger.Info)
}
