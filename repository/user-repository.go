package repository

import (
	"errors"
	"gorm.io/gorm"
	"hr-app-go/model"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) GetAUser(id string) model.User {
	var user model.User
	if err := r.DB.First(&user, id).Error; err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
	}
	return user
}
