package config

import (
	"gorm.io/gorm"
	"hr-app-go/repository"
)

type Repositories struct {
	MissionR *repository.MissionRepository
}

func InitRepositories(db *gorm.DB) Repositories {
	missionR := &repository.MissionRepository{DB: db}

	return Repositories{MissionR: missionR}
}
