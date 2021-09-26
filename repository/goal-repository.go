package repository

import (
	"gorm.io/gorm"
	"hr-app-go/model"
)

type GoalRepository struct {
	DB *gorm.DB
}

func (r *GoalRepository) GetAGoal(id string) (model.Goal, error) {
	// TODO takahashikazuaki not yet implemented
	return model.Goal{}, nil
}
