package usecase

import (
	"Chiprek/middleware"
	"Chiprek/models"
	"Chiprek/models/payload"
	"Chiprek/repository/database"

	"github.com/labstack/echo"
)

type CustomerUsecase interface {
	CreateCustomer(req *payload.CreateCustomerRequest) (res payload.CreateCustomerResponse, err error)
}

type customerUsecase struct {
	CustomerRepository database.CustomerRepository
}

func NewCustomerUsecase(customerRepository database.CustomerRepository) *customerUsecase {
	return &customerUsecase{customerRepository}
}

// Create Customer withs token
func (c *customerUsecase) CreateCustomer(req *payload.CreateCustomerRequest) (res payload.CreateCustomerResponse, err error) {
	customerReq := models.Customer{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
	}

	customer, err := c.CustomerRepository.CreateCustomer(&customerReq)
	if err != nil {
		echo.NewHTTPError(400, "Failed to create customer")
		return
	}

	token, err := middleware.CreateTokenUser(int(customer.ID), customer.Name)
	if err != nil {
		echo.NewHTTPError(400, "Failed to generate token")
		return
	}

	customer.Token = token

	res = payload.CreateCustomerResponse{
		Name:        customer.Name,
		PhoneNumber: customer.PhoneNumber,
		Token:       token,
	}

	return

}
