package database

import (
	"Chiprek/config"
	"Chiprek/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetAllTransactionRepository() (transaction []models.Transaction, err error)
	GetTransactionById(id int) (transaction *models.Transaction, err error)
	GetTransactionByCustomerId(id int) (transaction *models.Transaction, err error)
	CreateTransaction(tx *gorm.DB, transaction *models.Transaction) error
	UpdateTransaction(tx *gorm.DB, transaction *models.Transaction) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

// Get All Transaction by new date
func (t *transactionRepository) GetAllTransactionRepository() (transaction []models.Transaction, err error) {
	err = t.db.Preload("Customer").Preload("Cart.CartItem.Menu").Order("order_time asc").Find(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// Get Transaction by id
func (t *transactionRepository) GetTransactionById(id int) (transaction *models.Transaction, err error) {
	err = t.db.Preload("Customer").Preload("Cart.CartItem.Menu").Where("id = ?", id).First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// Get Transaction by customer id
func (t *transactionRepository) GetTransactionByCustomerId(id int) (transaction *models.Transaction, err error) {
	err = t.db.Preload("Customer").Preload("Cart.CartItem.Menu").Where("customer_id = ?", id).First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// create new transaction
func (t *transactionRepository) CreateTransaction(tx *gorm.DB, transaction *models.Transaction) error {
	db := config.DB
	if tx != nil {
		db = tx
	}

	err := db.Preload("Customer").Preload("Cart").Create(&transaction).Error
	if err != nil {
		return err
	}

	return nil
}

// update transaction
func (t *transactionRepository) UpdateTransaction(tx *gorm.DB, transaction *models.Transaction) error {
	db := config.DB
	if tx != nil {
		db = tx
	}

	err := db.Updates(&transaction).Error
	if err != nil {
		return err
	}

	return nil
}
