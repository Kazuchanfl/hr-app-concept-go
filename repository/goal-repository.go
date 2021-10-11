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

func (r *GoalRepository) GetGoalsOfAUser(userId string) ([]model.Goal, error) {
	var goals []model.Goal
	if err := r.DB.Where("UserID = ?", userId).Find(goals).Error; err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
	}
	return goals, nil
}
