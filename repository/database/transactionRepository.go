package database

import (
	"Chiprek/config"
	"Chiprek/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetAllTransactionRepository() (transaction []models.Transaction, err error)
	CreateTransactionRepository(tx *gorm.DB, transaction *models.Transaction) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

// Get All Transaction by new date
func (r *transactionRepository) GetAllTransactionRepository() (transaction []models.Transaction, err error) {
	err = r.db.Order("order_time desc").Find(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// create new transaction
func (r *transactionRepository) CreateTransactionRepository(tx *gorm.DB, transaction *models.Transaction) error {
	db := config.DB
	if tx != nil {
		db = tx
	}

	err := db.Preload("Cart").Create(&transaction).Error
	if err != nil {
		return err
	}

	return nil
}
