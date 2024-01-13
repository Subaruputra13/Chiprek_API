package usecase

import (
	"Chiprek/config"
	"Chiprek/models"
	"Chiprek/models/payload"
	"Chiprek/repository/database"
	"errors"
	"time"

	"github.com/google/uuid"
)

type TransactionUsecase interface {
	CreateTransactionUsecase(id int, req *payload.CreateTransactionRequest) (res models.Transaction, err error)
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
	err = t.transactionRepository.CreateTransactionRepository(nil, &transaction)
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
