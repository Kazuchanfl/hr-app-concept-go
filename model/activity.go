package model

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
	Title       string
	Description string
	UserID      uint
	GoalID      uint
}
