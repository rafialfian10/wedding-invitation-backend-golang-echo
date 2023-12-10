package repositories

import (
	"wedding/models"

	"gorm.io/gorm"
)

type OptionRepository interface {
	FindOptions() ([]models.Option, error)
	GetOption(ID int) (models.Option, error)
	CreateOption(option models.Option) (models.Option, error)
	UpdateOption(option models.Option) (models.Option, error)
	DeleteOption(option models.Option, ID int) (models.Option, error)
}

func RepositoryOption(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindOptions() ([]models.Option, error) {
	var options []models.Option
	err := r.db.Find(&options).Error

	return options, err
}

func (r *repository) GetOption(ID int) (models.Option, error) {
	var option models.Option
	err := r.db.First(&option, ID).Error

	return option, err
}

func (r *repository) CreateOption(option models.Option) (models.Option, error) {
	err := r.db.Create(&option).Error

	return option, err
}

func (r *repository) UpdateOption(option models.Option) (models.Option, error) {
	err := r.db.Debug().Model(&option).Updates(option).Error

	return option, err
}

func (r *repository) DeleteOption(option models.Option, ID int) (models.Option, error) {
	err := r.db.Raw("DELETE FROM options WHERE id=?", ID).Scan(&option).Error

	return option, err
}
