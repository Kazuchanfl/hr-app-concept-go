package model

import "gorm.io/gorm"

type Mission struct {
	gorm.Model
	Statement string
}
