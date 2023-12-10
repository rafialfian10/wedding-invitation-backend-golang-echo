package repositories

import (
	"wedding/models"

	"gorm.io/gorm"
)

type FagRepository interface {
	FindFags() ([]models.Fag, error)
	GetFag(ID int) (models.Fag, error)
	CreateFag(fag models.Fag) (models.Fag, error)
	UpdateFag(fag models.Fag) (models.Fag, error)
	DeleteFag(fag models.Fag, ID int) (models.Fag, error)
}

func RepositoryFag(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindFags() ([]models.Fag, error) {
	var fags []models.Fag
	err := r.db.Preload("FagContent.Option").Find(&fags).Error

	return fags, err
}

func (r *repository) GetFag(ID int) (models.Fag, error) {
	var fag models.Fag
	err := r.db.Preload("FagContent.Option").First(&fag, ID).Error

	return fag, err
}

func (r *repository) CreateFag(fag models.Fag) (models.Fag, error) {
	err := r.db.Create(&fag).Error

	return fag, err
}

func (r *repository) UpdateFag(fag models.Fag) (models.Fag, error) {
	err := r.db.Debug().Model(&fag).Updates(fag).Error

	return fag, err
}

func (r *repository) DeleteFag(fag models.Fag, ID int) (models.Fag, error) {
	err := r.db.Preload("FagContent.Option").First(&fag, ID).Error
	if err != nil {
		return fag, err
	}

	// delete related fag contents and options
	for _, fagContent := range fag.FagContent {
		for _, option := range fagContent.Option {
			if err := r.db.Delete(&option).Error; err != nil {
				return fag, err
			}
		}

		if err := r.db.Delete(&fagContent).Error; err != nil {
			return fag, err
		}
	}

	// delete fag
	err = r.db.Delete(&fag, ID).Error
	return fag, err
}
