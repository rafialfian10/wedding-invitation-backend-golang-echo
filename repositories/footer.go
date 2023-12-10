package repositories

import (
	"wedding/models"

	"gorm.io/gorm"
)

type FooterRepository interface {
	FindFooters() ([]models.Footer, error)
	GetFooter(ID int) (models.Footer, error)
	CreateFooter(footer models.Footer) (models.Footer, error)
	UpdateFooter(footer models.Footer) (models.Footer, error)
	DeleteFooter(footer models.Footer, ID int) (models.Footer, error)
}

func RepositoryFooter(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindFooters() ([]models.Footer, error) {
	var footer []models.Footer
	err := r.db.Find(&footer).Error

	return footer, err
}

func (r *repository) GetFooter(ID int) (models.Footer, error) {
	var footer models.Footer
	err := r.db.First(&footer, ID).Error

	return footer, err
}

func (r *repository) CreateFooter(footer models.Footer) (models.Footer, error) {
	err := r.db.Create(&footer).Error

	return footer, err
}

func (r *repository) UpdateFooter(footer models.Footer) (models.Footer, error) {
	err := r.db.Debug().Model(&footer).Updates(footer).Error

	return footer, err
}

func (r *repository) DeleteFooter(footer models.Footer, ID int) (models.Footer, error) {
	err := r.db.Raw("DELETE FROM footers WHERE id=?", ID).Scan(&footer).Error

	return footer, err
}
