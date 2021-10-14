package repository

import "gorm.io/gorm"

type ActivityRepository struct {
	DB *gorm.DB
}

func (r *ActivityRepository) CreateAnActivity() error {
	// TODO create an activity and store in DB
	return nil
}
