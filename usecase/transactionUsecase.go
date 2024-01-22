package usecase

import (
	"Chiprek/config"
	"Chiprek/models"
	"Chiprek/models/payload"
	"Chiprek/repository/database"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type TransactionUsecase interface {
	CreateTransactionUsecase(id int, req *payload.CreateTransactionRequest) (res models.Transaction, err error)
	GetAllTransaction() ([]models.Transaction, error)
	GetTransactionById(id int) (*models.Transaction, error)
	GetTransactionByCustomerId(id int) (*models.Transaction, error)
}

type transactionUsecase struct {
	transactionRepository database.TransactionRepository
	cartRespository       database.CartRepository
}

func NewTransactionUsecase(transactionRepository database.TransactionRepository, cartRespository database.CartRepository) *transactionUsecase {
	return &transactionUsecase{transactionRepository, cartRespository}
}

// Create transaction
func (t *transactionUsecase) CreateTransactionUsecase(id int, req *payload.CreateTransactionRequest) (res models.Transaction, err error) {
	// get cart by customer id
	cart, err := t.cartRespository.GetCartByCustomerID(id)
	if err != nil {
		return res, err
	}

	// create generate transaction id ex: MRGD-0001 increment
	transactionId := uuid.New().String()
	transactionId = "MRGD" + "-" + transactionId[0:4]

	// create transaction
	transaction := models.Transaction{
		CustomerID:    cart.CustomerID,
		Customer:      cart.Customer,
		CartID:        int(cart.ID),
		Cart:          *cart,
		CustomerName:  cart.Customer.Name,
		PhoneNumber:   cart.Customer.PhoneNumber,
		TranscationId: transactionId,
		OrderType:     "Dine In",
		OrderTime:     time.Now(),
		PaymentType:   req.PaymentType,
		Status:        false,
		TotalPrice:    cart.TotalPrice,
	}

	tx := config.DB.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// create transaction
	err = t.transactionRepository.CreateTransaction(nil, &transaction)
	if err != nil {
		return res, err
	}

	err = tx.Commit().Error
	if err != nil {
		errors.New("Failed to commit transaction")
		return
	}

	return transaction, nil
}

// Get all transaction
func (t *transactionUsecase) GetAllTransaction() ([]models.Transaction, error) {
	transaction, err := t.transactionRepository.GetAllTransactionRepository()
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

// Get transaction by id
func (t *transactionUsecase) GetTransactionById(id int) (*models.Transaction, error) {
	transaction, err := t.transactionRepository.GetTransactionById(id)
	if err != nil {
		return nil, echo.NewHTTPError(400, err.Error())
	}

	return transaction, nil
}

// Get transaction by customer id
func (t *transactionUsecase) GetTransactionByCustomerId(id int) (*models.Transaction, error) {
	transaction, err := t.transactionRepository.GetTransactionByCustomerId(id)
	if err != nil {
		return nil, echo.NewHTTPError(400, err.Error())
	}

	return transaction, nil
}
