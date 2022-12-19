package models

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kk-rschain.com/gin-rest-auth/config"
)

var (
	DB  *gorm.DB
	err error
)

func Setup() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
		config.Database.Charset,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database error: %v", err)
	}
}

func Close() {
	mysqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("database error: %v", err)
	}
	if err := mysqlDB.Close(); err != nil {
		log.Fatalf("database error: %v", err)
	}
}
