package repositories

import (
	"wedding/models"

	"gorm.io/gorm"
)

type NavigationRepository interface {
	FindNavigations() ([]models.Navigation, error)
	GetNavigation(ID int) (models.Navigation, error)
	CreateNavigation(navigation models.Navigation) (models.Navigation, error)
	UpdateNavigation(navigation models.Navigation) (models.Navigation, error)
	DeleteNavigation(navigation models.Navigation, ID int) (models.Navigation, error)
}

func RepositoryNavigation(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindNavigations() ([]models.Navigation, error) {
	var navigations []models.Navigation
	err := r.db.Find(&navigations).Error
	return navigations, err
}

func (r *repository) GetNavigation(ID int) (models.Navigation, error) {
	var navigation models.Navigation
	err := r.db.First(&navigation, ID).Error
	return navigation, err
}

func (r *repository) CreateNavigation(navigation models.Navigation) (models.Navigation, error) {
	err := r.db.Create(&navigation).Error
	return navigation, err
}

func (r *repository) UpdateNavigation(navigation models.Navigation) (models.Navigation, error) {
	err := r.db.Debug().Model(&navigation).Updates(navigation).Error
	return navigation, err
}

func (r *repository) DeleteNavigation(navigation models.Navigation, ID int) (models.Navigation, error) {
	err := r.db.Raw("DELETE FROM navigations WHERE id=?", ID).Scan(&navigation).Error
	return navigation, err
}
