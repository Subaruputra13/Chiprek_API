package usecase

import (
	"Chiprek/config"
	"Chiprek/models"
	"Chiprek/models/payload"
	"Chiprek/repository/database"
	"Chiprek/util"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type TransactionUsecase interface {
	CreateTransactionUsecase(id int) (res models.Transaction, err error)
	GetAllTransaction() ([]models.Transaction, error)
	GetTransactionById(id int) (*models.Transaction, error)
	GetTransactionByCustomerId(id int) (*models.Transaction, error)
	ProcessPayment(req *payload.TransactionNotificationInput) error
	UpdateTransactionStatusById(transaction *models.Transaction, req *payload.UpdateTransactionRequest) error
}

type transactionUsecase struct {
	transactionRepository database.TransactionRepository
	cartRespository       database.CartRepository
}

func NewTransactionUsecase(transactionRepository database.TransactionRepository, cartRespository database.CartRepository) *transactionUsecase {
	return &transactionUsecase{transactionRepository, cartRespository}
}

// Create transaction
func (t *transactionUsecase) CreateTransactionUsecase(id int) (res models.Transaction, err error) {
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
		OrderID:      transactionId,
		CustomerID:   cart.CustomerID,
		Customer:     cart.Customer,
		CartID:       int(cart.ID),
		Cart:         *cart,
		CustomerName: cart.Customer.Name,
		PhoneNumber:  cart.Customer.PhoneNumber,
		OrderType:    "Dine In",
		// OrderTime:     time.Now(),
		// PaymentType:   req.PaymentType,
		Status:     "Waiting for Payment",
		TotalPrice: cart.TotalPrice,
	}

	tx := config.DB.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// create transaction
	err = t.transactionRepository.CreateTransaction(tx, &transaction)
	if err != nil {
		return res, err
	}

	// Midtrans
	responseMidtrans, err := util.GetPaymentURL(&transaction, &cart.Customer)
	if err != nil {
		return res, err
	}

	// update payment url
	transaction.PaymentURL = responseMidtrans.RedirectURL

	// update transaction
	err = t.transactionRepository.UpdateTransaction(tx, &transaction)
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

// Midtrans Notificaiton
func (t *transactionUsecase) ProcessPayment(req *payload.TransactionNotificationInput) error {
	transaction, err := t.transactionRepository.GetTransactionByOrderId(req.OrderID)
	if err != nil {
		return err
	}

	tx := config.DB.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// update payment Method
	transaction.PaymentMethod = req.PaymentType

	if req.TransactionStatus == "settlement" || req.TransactionStatus == "capture" {
		transaction.Status = "On Process"
		transaction.PaymentStatus = "Paid"

		date, err := time.Parse("2006-01-02 15:04:05", req.TransactionTime)
		if err != nil {
			return err
		}

		transaction.PaymentDate = &date
		err = t.transactionRepository.UpdateTransaction(tx, transaction)
		if err != nil {
			return err
		}
	} else if req.TransactionStatus != "pending" {
		transaction.Status = "Cancelled"
		transaction.PaymentStatus = "Cancelled"

		err = t.transactionRepository.UpdateTransaction(tx, transaction)
		if err != nil {
			return err
		}
	}

	err = tx.Commit().Error
	if err != nil {
		errors.New("Failed to commit transaction")
		return err
	}

	return nil
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

// Update Transaction Status By Id
func (t *transactionUsecase) UpdateTransactionStatusById(transaction *models.Transaction, req *payload.UpdateTransactionRequest) error {
	// Check Transaction By Id
	transaction, err := t.transactionRepository.GetTransactionById(int(transaction.ID))
	if err != nil {
		return err
	}

	transaction.Status = req.Status

	tx := config.DB.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	err = t.transactionRepository.UpdateTransaction(tx, transaction)
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		errors.New("Failed to commit transaction")
		return err
	}

	return nil

}
