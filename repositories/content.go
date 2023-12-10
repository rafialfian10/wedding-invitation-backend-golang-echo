package repositories

import (
	"wedding/models"

	"gorm.io/gorm"
)

type ContentRepository interface {
	FindContents() ([]models.Content, error)
	GetContent(ID int) (models.Content, error)
	CreateContent(content models.Content) (models.Content, error)
	UpdateContent(content models.Content) (models.Content, error)
	DeleteContent(content models.Content, ID int) (models.Content, error)
}

func RepositoryContent(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindContents() ([]models.Content, error) {
	var contents []models.Content
	err := r.db.Preload("Feature").Find(&contents).Error

	return contents, err
}

func (r *repository) GetContent(ID int) (models.Content, error) {
	var content models.Content
	err := r.db.Preload("Feature").First(&content, ID).Error

	return content, err
}

func (r *repository) CreateContent(content models.Content) (models.Content, error) {
	err := r.db.Create(&content).Error

	return content, err
}

func (r *repository) UpdateContent(content models.Content) (models.Content, error) {
	err := r.db.Debug().Model(&content).Updates(content).Error

	return content, err
}

func (r *repository) DeleteContent(content models.Content, ID int) (models.Content, error) {
	err := r.db.Preload("Feature").First(&content, ID).Error
	if err != nil {
		return content, err
	}

	// Delete related features
	for _, feature := range content.Feature {
		if err := r.db.Delete(&feature).Error; err != nil {
			return content, err
		}
	}

	// delete content
	err = r.db.Delete(&content, ID).Error
	return content, err
}
