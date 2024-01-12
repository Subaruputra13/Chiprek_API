package controllers

import (
	"Chiprek/models/payload"
	"Chiprek/usecase"

	"github.com/labstack/echo"
)

type CustomerControllers interface {
	CreateCustomerControllers(c echo.Context) error
}

type customerControllers struct {
	customerUsecase usecase.CustomerUsecase
}

func NewCustomerControllers(customerUsecase usecase.CustomerUsecase) *customerControllers {
	return &customerControllers{customerUsecase}
}

// Controller Create Customer
func (cu *customerControllers) CreateCustomerControllers(c echo.Context) error {
	req := payload.CreateCustomerRequest{}

	c.Bind(&req)

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(400, "Field cannot be empty")
	}

	res, err := cu.customerUsecase.CreateCustomer(&req)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Create Customer",
		Data:    res,
	})
}
