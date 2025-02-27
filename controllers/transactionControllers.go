package controllers

import (
	"Chiprek/middleware"
	"Chiprek/models/payload"
	"Chiprek/usecase"
	"strconv"

	"github.com/labstack/echo"
)

type TransactionController interface {
	CreateTransactionController(c echo.Context) error
}

type transactionController struct {
	transactionUsecase usecase.TransactionUsecase
}

func NewTransactionController(transactionUsecase usecase.TransactionUsecase) *transactionController {
	return &transactionController{transactionUsecase}
}

// Controller Create Transaction
func (t *transactionController) CreateTransactionController(c echo.Context) error {
	// req := payload.CreateTransactionRequest{}

	CustomerId, err := middleware.IsCustomer(c)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	// c.Bind(&req)

	// if err := c.Validate(req); err != nil {
	// 	return echo.NewHTTPError(400, err.Error())
	// }

	transaction, err := t.transactionUsecase.CreateTransactionUsecase(CustomerId)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Create Transaction",
		Data:    transaction,
	})
}

// Controller Get All Transaction
func (t *transactionController) GetAllTransactionController(c echo.Context) error {
	transaction, err := t.transactionUsecase.GetAllTransaction()
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Get All Transaction",
		Data:    transaction,
	})
}

// Controller Get Transaction By Id
func (t *transactionController) GetTransactionByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := t.transactionUsecase.GetTransactionById(id)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Get Transaction By Id",
		Data:    transaction,
	})
}

// Controller Get Transaction By Customer Id
func (t *transactionController) GetTransactionByCustomerIdController(c echo.Context) error {
	CustomerId, err := middleware.IsCustomer(c)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	transaction, err := t.transactionUsecase.GetTransactionByCustomerId(CustomerId)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Get Transaction By Customer Id",
		Data:    transaction,
	})
}

func (t *transactionController) GetNotificationController(c echo.Context) error {
	payloadNotification := payload.TransactionNotificationInput{}

	c.Bind(&payloadNotification)

	if err := c.Validate(payloadNotification); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	err := t.transactionUsecase.ProcessPayment(&payloadNotification)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Process Payment",
		Data:    nil,
	})
}
