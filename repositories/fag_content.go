package repositories

import (
	"wedding/models"

	"gorm.io/gorm"
)

type FagContentRepository interface {
	FindFagContents() ([]models.FagContent, error)
	GetFagContent(ID int) (models.FagContent, error)
	CreateFagContent(fagContent models.FagContent) (models.FagContent, error)
	UpdateFagContent(fagContent models.FagContent) (models.FagContent, error)
	DeleteFagContent(fagContent models.FagContent, ID int) (models.FagContent, error)
}

func RepositoryFagContent(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindFagContents() ([]models.FagContent, error) {
	var fagContents []models.FagContent
	err := r.db.Preload("Option").Find(&fagContents).Error

	return fagContents, err
}

func (r *repository) GetFagContent(ID int) (models.FagContent, error) {
	var fagContent models.FagContent
	err := r.db.Preload("Option").First(&fagContent, ID).Error

	return fagContent, err
}

func (r *repository) CreateFagContent(fagContent models.FagContent) (models.FagContent, error) {
	err := r.db.Create(&fagContent).Error

	return fagContent, err
}

func (r *repository) UpdateFagContent(fagContent models.FagContent) (models.FagContent, error) {
	err := r.db.Debug().Model(&fagContent).Updates(fagContent).Error

	return fagContent, err
}

func (r *repository) DeleteFagContent(fagContent models.FagContent, ID int) (models.FagContent, error) {
	err := r.db.Preload("Option").First(&fagContent, ID).Error
	if err != nil {
		return fagContent, err
	}

	// Delete related options
	for _, option := range fagContent.Option {
		if err := r.db.Delete(&option).Error; err != nil {
			return fagContent, err
		}
	}

	// delete fag content
	err = r.db.Delete(&fagContent, ID).Error
	return fagContent, err
}
