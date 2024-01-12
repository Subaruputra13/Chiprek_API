package controllers

import (
	"Chiprek/middleware"
	"Chiprek/models/payload"
	"Chiprek/usecase"

	"github.com/labstack/echo"
)

type CartControllers interface {
	GetCartByCustomerIDControllers(c echo.Context) error
	AddMenuToCartControllers(c echo.Context) error
}

type cartControllers struct {
	cartUsecase usecase.CartUsecase
}

func NewCartControllers(cartUsecase usecase.CartUsecase) *cartControllers {
	return &cartControllers{cartUsecase}
}

// Controller Get Cart By Customer ID
func (ca *cartControllers) GetCartByCustomerIDControllers(c echo.Context) error {
	CustomerId, err := middleware.IsCustomer(c)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	cart, err := ca.cartUsecase.GetCartByCustomerID(CustomerId)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Get Cart",
		Data:    cart,
	})
}

// Controller Add Menu to Cart
func (ca *cartControllers) AddMenuToCartControllers(c echo.Context) error {
	req := payload.AddMenuToCartRequest{}

	CustomerId, err := middleware.IsCustomer(c)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(400, "Field cannot be empty")
	}

	err = ca.cartUsecase.AddMenuToCart(CustomerId, &req)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, "Success Add Menu to Cart")
}
