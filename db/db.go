package db

import (
	// _ "github.com/go-sql-driver/mysql"

	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var user, pass, host string

func GetDB() *gorm.DB {
	return db
}

func dbConnect(dialector gorm.Dialector, config gorm.Option, count uint) (err error) {
	// NOTE: 指定回数だけDB接続を試みる
	for count > 1 {
		DB, err := gorm.Open(dialector, config)
		if err != nil {
			time.Sleep(time.Second * 2)
			count--
			log.Printf("retry... count:%v\n", count)
			continue
		} else {
			DB.Logger.LogMode(logger.Info)
			db = DB
			return nil
		}

	}

	return err
}

func InitDB() {
	user = os.Getenv("MYSQL_USER")
	pass = os.Getenv("MYSQL_PASSWORD")
	host = os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, dbname)

	err := dbConnect(mysql.Open(dsn), &gorm.Config{}, 10)
	if err != nil {
		log.Fatalln(err)
	}
}
