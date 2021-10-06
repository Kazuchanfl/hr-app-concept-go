package config

import (
	"gorm.io/gorm"
	"hr-app-go/repository"
)

type Repositories struct {
	MissionR *repository.MissionRepository
	ValueR   *repository.ValueRepository
	UserR    *repository.UserRepository
}

func InitRepositories(db *gorm.DB) Repositories {
	missionR := &repository.MissionRepository{DB: db}
	userR := &repository.UserRepository{DB: db}
	valueR := &repository.ValueRepository{DB: db}

	return Repositories{MissionR: missionR, ValueR: valueR, UserR: userR}
}
