package config

import (
	"gorm.io/gorm"
	"hr-app-go/repository"
)

type Repositories struct {
	MissionR *repository.MissionRepository
	UserR    *repository.UserRepository
}

func InitRepositories(db *gorm.DB) Repositories {
	missionR := &repository.MissionRepository{DB: db}
	userR := &repository.UserRepository{DB: db}

	return Repositories{MissionR: missionR, UserR: userR}
}
