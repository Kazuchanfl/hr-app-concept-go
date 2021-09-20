package repository

import (
	"gorm.io/gorm"
	"hr-app-go/model"
)

type MissionRepository struct {
	DB *gorm.DB
}

func (r *MissionRepository) GetAllMissions() ([]model.Mission, error) {
	var missions []model.Mission
	if err := r.DB.Find(&missions).Error; err != nil {
		return nil, err
	}
	return missions, nil
}
