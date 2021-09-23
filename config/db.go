package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"hr-app-go/model"
)

func InitDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Mission{})
	db.AutoMigrate(&model.Value{})
	db.AutoMigrate(&model.Goal{})

	return db
}
