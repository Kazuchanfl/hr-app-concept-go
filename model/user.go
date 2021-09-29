package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string
	Profile string
	Goals   []Goal
}
