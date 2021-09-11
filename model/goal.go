package model

import "gorm.io/gorm"

type Goal struct {
	gorm.Model
	Title       string
	Description string
}
