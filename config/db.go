package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"hr-app-go/model"
)

func InitDb() {
	print("This is gorm-test")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Mission{})
}
