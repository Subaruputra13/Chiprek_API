package controllers

import (
	"Chiprek/middleware"
	"Chiprek/models/payload"
	"Chiprek/usecase"

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
	req := payload.CreateTransactionRequest{}

	CustomerId, err := middleware.IsCustomer(c)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	transaction, err := t.transactionUsecase.CreateTransactionUsecase(CustomerId, &req)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Create Transaction",
		Data:    transaction,
	})
}
