package repositories

import (
	"wedding/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	FindUserByUsernameOrEmail(username, email string) (models.User, error)
	Login(email string) (models.User, error)
	CheckAuth(ID int) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Preload("Premi").Create(&user).Error
	return user, err
}

func (r *repository) FindUserByUsernameOrEmail(username, email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "username=? OR email=?", username, email).Error
	return user, err
}

func (r *repository) Login(email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email=?", email).Error
	return user, err
}

func (r *repository) CheckAuth(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error
	return user, err
}
