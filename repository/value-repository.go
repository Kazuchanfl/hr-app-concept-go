package repository

import "gorm.io/gorm"

type ValueRepository struct {
	DB *gorm.DB
}
