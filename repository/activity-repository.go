package repository

import "gorm.io/gorm"

type ActivityRepository struct {
	DB *gorm.DB
}
