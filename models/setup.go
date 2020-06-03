package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func SetupModels() *gorm.DB {
	fmt.Println("trying to print")
	db, err := gorm.Open("sqlite3", "orders.db")

	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("made it pass fail")
	db.AutoMigrate(&OrderModel{})

	return db
}
