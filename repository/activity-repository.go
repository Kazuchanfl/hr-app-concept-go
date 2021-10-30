package repository

import (
	"gorm.io/gorm"
	"hr-app-go/model"
)

type ActivityRepository struct {
	DB *gorm.DB
}

func (r *ActivityRepository) CreateAnActivity(activity model.Activity) error {
	result := r.DB.Create(&activity)
	return result.Error
}
