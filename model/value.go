package model

import "gorm.io/gorm"

type Value struct {
	gorm.Model
	Name        string
	Description string
}
