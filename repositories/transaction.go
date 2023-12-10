package repositories

import (
	"wedding/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
	FindTransactionsByUser(UserID int) ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(status string, ID int) (models.Transaction, error)
	UpdateTokenTransaction(token string, ID int) (models.Transaction, error)
	DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("Pricing.Content.Feature").Preload("User").Find(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Pricing.Content.Feature").Preload("User").First(&transaction, ID).Error

	return transaction, err
}

func (r *repository) FindTransactionsByUser(UserID int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("Pricing.Content.Feature").Preload("User").Where("user_id = ?", UserID).Order("booking_date desc").Find(&transaction).Error

	return transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error

	return transaction, err
}

func (r *repository) UpdateTransaction(status string, Id int) (models.Transaction, error) {
	var transaction models.Transaction
	r.db.Preload("Pricing.Content.Feature").Preload("User").First(&transaction, "id = ?", Id)

	transaction.Status = status
	err := r.db.Model(&transaction).Updates(transaction).Error

	return transaction, err
}

func (r *repository) UpdateTokenTransaction(token string, Id int) (models.Transaction, error) {
	var transaction models.Transaction
	r.db.Preload("Pricing.Content.Feature").Preload("User").First(&transaction, "id = ?", Id)

	transaction.Token = token
	err := r.db.Model(&transaction).Updates(transaction).Error

	return transaction, err
}

func (r *repository) DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Raw("DELETE FROM transactions WHERE id=?", ID).Scan(&transaction).Error

	return transaction, err
}
