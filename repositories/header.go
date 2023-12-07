package repositories

import (
	"wedding/models"

	"gorm.io/gorm"
)

type HeaderRepository interface {
	FindHeaders() ([]models.Header, error)
	GetHeader(ID int) (models.Header, error)
	CreateHeader(header models.Header) (models.Header, error)
	UpdateHeader(header models.Header) (models.Header, error)
	DeleteHeader(header models.Header, ID int) (models.Header, error)
	DeleteHeaderImage(ID int) error
}

func RepositoryHeader(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindHeaders() ([]models.Header, error) {
	var headers []models.Header
	err := r.db.Find(&headers).Error
	return headers, err
}

func (r *repository) GetHeader(ID int) (models.Header, error) {
	var header models.Header
	err := r.db.First(&header, ID).Error
	return header, err
}

func (r *repository) CreateHeader(header models.Header) (models.Header, error) {
	err := r.db.Create(&header).Error
	return header, err
}

func (r *repository) UpdateHeader(header models.Header) (models.Header, error) {
	err := r.db.Debug().Model(&header).Updates(header).Error
	return header, err
}

func (r *repository) DeleteHeader(header models.Header, ID int) (models.Header, error) {
	err := r.db.Raw("DELETE FROM headers WHERE id=?", ID).Scan(&header).Error
	return header, err
}

func (r *repository) DeleteHeaderImage(ID int) error {
	return r.db.Model(&models.Header{}).Where("id = ?", ID).UpdateColumn("header", gorm.Expr("NULL")).Error
}
