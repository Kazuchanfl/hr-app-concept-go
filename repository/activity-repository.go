package repository

import (
	"gorm.io/gorm"
	"hr-app-go/model"
)

type ActivityRepository struct {
	DB *gorm.DB
}

// TODO takahashikazuaki Get activity object and create record with it
func (r *ActivityRepository) CreateAnActivity() error {
	activity := model.Activity{
		Title:       "1日に３つLGTMを出した",
		Description: "レビューが必要なプルリクエストが多かったので３つ分レビューし、LGTMを出した。",
		UserID:      1,
		GoalID:      1,
	}
	result := r.DB.Create(&activity)
	return result.Error
}
