package repositories

import (
	"wedding/models"

	"gorm.io/gorm"
)

type PricingRepository interface {
	FindPricings() ([]models.Pricing, error)
	GetPricing(ID int) (models.Pricing, error)
	CreatePricing(pricing models.Pricing) (models.Pricing, error)
	UpdatePricing(pricing models.Pricing) (models.Pricing, error)
	DeletePricing(pricing models.Pricing, ID int) (models.Pricing, error)
	DeleteImage(ID int) error
}

func RepositoryPricing(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindPricings() ([]models.Pricing, error) {
	var pricings []models.Pricing
	err := r.db.Preload("Content.Feature").Find(&pricings).Error

	return pricings, err
}

func (r *repository) GetPricing(ID int) (models.Pricing, error) {
	var pricing models.Pricing
	err := r.db.Preload("Content.Feature").First(&pricing, ID).Error

	return pricing, err
}

func (r *repository) CreatePricing(pricing models.Pricing) (models.Pricing, error) {
	err := r.db.Create(&pricing).Error

	return pricing, err
}

func (r *repository) UpdatePricing(pricing models.Pricing) (models.Pricing, error) {
	err := r.db.Debug().Model(&pricing).Updates(pricing).Error

	return pricing, err
}

func (r *repository) DeletePricing(pricing models.Pricing, ID int) (models.Pricing, error) {
	err := r.db.Preload("Content.Feature").First(&pricing, ID).Error
	if err != nil {
		return pricing, err
	}

	// delete related features and content
	for _, content := range pricing.Content {
		for _, feature := range content.Feature {
			if err := r.db.Delete(&feature).Error; err != nil {
				return pricing, err
			}
		}
		if err := r.db.Delete(&content).Error; err != nil {
			return pricing, err
		}
	}

	// delete pricing
	err = r.db.Delete(&pricing, ID).Error
	return pricing, err

}

func (r *repository) DeleteImage(ID int) error {
	return r.db.Model(&models.Pricing{}).Where("id = ?", ID).UpdateColumn("image", gorm.Expr("NULL")).Error
}
