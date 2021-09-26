package repository

import (
	"errors"
	"gorm.io/gorm"
	"hr-app-go/model"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) GetAUser(id string) (model.User, error) {
	var user model.User
	if err := r.DB.First(&user, id).Error; err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
	}
	return user, nil
}

func (r *UserRepository) UpdateAUser(user model.User) (model.User, error) {
	// TODO takahashikazuaki Not yet implemented
	return model.User{}, nil
}
