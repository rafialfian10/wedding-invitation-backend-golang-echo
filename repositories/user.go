package repositories

import (
	"wedding/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(user models.User, ID int) (models.User, error)
	GetProfile(userId int) (models.User, error)
	DeletePhoto(ID int) error
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error
	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Debug().Model(&user).Updates(user).Error
	return user, err
}

func (r *repository) DeleteUser(user models.User, ID int) (models.User, error) {
	err := r.db.Raw("DELETE FROM users WHERE id=?", ID).Scan(&user).Error
	return user, err
}

func (r *repository) GetProfile(userId int) (models.User, error) {
	var profile models.User
	err := r.db.First(&profile, userId).Error
	return profile, err
}

func (r *repository) DeletePhoto(ID int) error {
	return r.db.Model(&models.User{}).Where("id = ?", ID).UpdateColumn("photo", gorm.Expr("NULL")).Error
}
