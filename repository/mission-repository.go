package repository

import (
	"gorm.io/gorm"
	"hr-app-go/model"
)

type MissionRepository struct {
	DB *gorm.DB
}

func (r *MissionRepository) GetAllMissions() []model.Mission {
	// TODO takahashikazuaki get Mission from DB
	return []model.Mission{
		{Statement: "人の成長を支えることを通じて世の中の成長を実現する。"},
		{Statement: "生き物への慈愛の念を広め、人と動物が共生できる社会を実現する。"},
	}
}
