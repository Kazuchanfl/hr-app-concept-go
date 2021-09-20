package repository

import (
	"errors"
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

func (r *MissionRepository) GetAMission(id string) model.Mission {
	var mission model.Mission
	if err := r.DB.First(&mission, id).Error; err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
	}
	return mission
}
