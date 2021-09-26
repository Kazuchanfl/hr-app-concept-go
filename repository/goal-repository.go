package repository

import "gorm.io/gorm"

type GoalRepository struct {
	DB *gorm.DB
}
