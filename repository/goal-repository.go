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
	// TODO takahashikazuaki get goals from DB
	goals := []model.Goal{
		{
			Title:       "メンバーと毎週定例を実施する",
			Description: "その日やっていたこと・今後やっていこうとしていること・課題などを共有し、解決に向けて前進させる。",
		},
	}
	return goals, nil
}
