package usecase

import (
	"Chiprek/models"
	"Chiprek/models/payload"
	"Chiprek/repository/database"

	"github.com/labstack/echo"
)

type CartUsecase interface{}

type cartUsecase struct {
	cartRepository database.CartRepository
	menuRepository database.MenuRepository
}

func NewCartUsecase(
	cartRepository database.CartRepository,
	menuRepository database.MenuRepository,
) *cartUsecase {
	return &cartUsecase{cartRepository, menuRepository}
}

// Logic Add Menu to Cart
func (c *cartUsecase) AddMenuToCart(req *payload.AddMenuToCartRequest) error {
	menu, err := c.menuRepository.GetMenuByID(req.MenuID)
	if err != nil {
		return echo.NewHTTPError(400, "Menu not found")
	}

	// create cart
	cartReq := &models.Cart{
		TotalPrice: menu.Price * req.Quantity,
		Status:     true,
	}

	err = c.cartRepository.CreateCart(cartReq)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	// create cart item
	cartItemReq := &models.CartItem{
		Quantity: req.Quantity,
		TakeAway: req.TakeAway,
		Note:     req.Note,
		MenuID:   int(menu.ID),
		CartID:   int(cartReq.ID),
	}

	err = c.cartRepository.CreateCartItem(cartItemReq)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return nil

}
