package repositories

import (
	"wedding/models"

	"gorm.io/gorm"
)

type FeatureRepository interface {
	FindFeatures() ([]models.Feature, error)
	GetFeature(ID int) (models.Feature, error)
	CreateFeature(feature models.Feature) (models.Feature, error)
	UpdateFeature(feature models.Feature) (models.Feature, error)
	DeleteFeature(feature models.Feature, ID int) (models.Feature, error)
}

func RepositoryFeature(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindFeatures() ([]models.Feature, error) {
	var features []models.Feature
	err := r.db.Find(&features).Error

	return features, err
}

func (r *repository) GetFeature(ID int) (models.Feature, error) {
	var feature models.Feature
	err := r.db.First(&feature, ID).Error

	return feature, err
}

func (r *repository) CreateFeature(feature models.Feature) (models.Feature, error) {
	err := r.db.Create(&feature).Error

	return feature, err
}

func (r *repository) UpdateFeature(feature models.Feature) (models.Feature, error) {
	err := r.db.Debug().Model(&feature).Updates(feature).Error

	return feature, err
}

func (r *repository) DeleteFeature(feature models.Feature, ID int) (models.Feature, error) {
	err := r.db.Raw("DELETE FROM features WHERE id=?", ID).Scan(&feature).Error

	return feature, err
}
