package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func SetupOrderModel() *gorm.DB {
	db, err := gorm.Open("sqlite3", "orders.db")
	if err != nil {
		panic("Failed to connect to orders database")
	}
	db.AutoMigrate(&OrderModel{})
	return db
}

func SetupUserModel() *gorm.DB {
	db, err := gorm.Open("sqlite3", "users.db")
	if err != nil {
		panic("Failed to connect to users database")
	}
	db.AutoMigrate(&UserModel{})
	return db
}
