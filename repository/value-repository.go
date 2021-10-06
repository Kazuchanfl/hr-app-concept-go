package repository

import (
	"gorm.io/gorm"
	"hr-app-go/model"
)

type ValueRepository struct {
	DB *gorm.DB
}

func (r *ValueRepository) GetAllValues() ([]model.Value, error) {
	// TODO takahashikazuaki handle DB error
	var values []model.Value
	r.DB.Find(&values)
	return values, nil
}
