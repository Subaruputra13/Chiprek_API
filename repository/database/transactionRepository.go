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
	GetTransactionByOrderId(id string) (transaction *models.Transaction, err error)
	CreateTransaction(tx *gorm.DB, transaction *models.Transaction) error
	UpdateTransaction(tx *gorm.DB, transaction *models.Transaction) error
	UpdateTransactionById(id uint, transaction *models.Transaction) error
	SumTransactionsAmount() (income int, err error)
	TotalOrder() (total int64, err error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

// Get All Transaction by new date
func (t *transactionRepository) GetAllTransactionRepository() (transaction []models.Transaction, err error) {
	err = t.db.Preload("Customer").Preload("Cart.CartItem.Menu").Order("created_at desc").Find(&transaction).Error
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

// Get Transaction by transaction id
func (t *transactionRepository) GetTransactionByOrderId(id string) (transaction *models.Transaction, err error) {
	err = t.db.Preload("Customer").Preload("Cart.CartItem.Menu").Where("order_id = ?", id).Find(&transaction).Error
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

// Update Transaction By Id
func (t *transactionRepository) UpdateTransactionById(id uint, transaction *models.Transaction) error {
	err := t.db.Model(&models.Transaction{}).Where("id = ?", id).Updates(&transaction).Error
	if err != nil {
		return err
	}
	return nil
}

// Total Transation
func (t *transactionRepository) SumTransactionsAmount() (income int, err error) {
	err = t.db.Table("transactions").Select("COALESCE(sum(total_price), 0)").Where("payment_status = ?", "Paid").Row().Scan(&income)
	if err != nil {
		return income, err
	}
	return income, nil
}

// Total Order
func (t *transactionRepository) TotalOrder() (total int64, err error) {
	Transactions := []models.Transaction{}
	if err := t.db.Model(&Transactions).Where("payment_status = ?", "Paid").Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}
