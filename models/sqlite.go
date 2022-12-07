package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Setup() {
	DB, err = gorm.Open(sqlite.Open("gintest.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func Close() {
	sqliteDB, err := DB.DB()
	if err != nil {
		panic(err)
	}
	if err := sqliteDB.Close(); err != nil {
		panic(err)
	}
}
