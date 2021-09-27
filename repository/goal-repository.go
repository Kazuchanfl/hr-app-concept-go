package repository

import (
	"errors"
	"gorm.io/gorm"
	"hr-app-go/model"
)

type GoalRepository struct {
	DB *gorm.DB
}

func (r *GoalRepository) GetAGoal(id string) (model.Goal, error) {
	var goal model.Goal
	if err := r.DB.First(&goal, id).Error; err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
	}
	return goal, nil
}
