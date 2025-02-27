package usecase

import (
	"Chiprek/models"
	"Chiprek/models/payload"
	"Chiprek/repository/database"

	"github.com/labstack/echo"
)

type CartUsecase interface {
	AddMenuToCart(id int, req *payload.AddMenuToCartRequest) error
	GetCartByCustomerID(id int) (*models.Cart, error)
	DeleteCartItem(id int, req *payload.DeleteCartItemRequest) error
	// UpdateCartItem(id int, req *payload.UpdateCartItemRequest) error
	// UpdateCartItem(req *payload.UpdateCartItemRequest) error
	// DeleteCartItem(id int) error
	// DeleteCart(id int) error
}

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
func (c *cartUsecase) AddMenuToCart(id int, req *payload.AddMenuToCartRequest) error {
	menu, err := c.menuRepository.GetMenuByID(req.MenuID)
	if err != nil {
		return echo.NewHTTPError(400, "Menu not found")
	}

	// Check if cart exist
	_, err = c.cartRepository.GetCartByCustomerID(id)
	if err != nil {
		cartReq := models.Cart{
			CustomerID: id,
			Status:     true,
		}

		err = c.cartRepository.CreateCart(&cartReq)
		if err != nil {
			return echo.NewHTTPError(400, "Failed to create cart")
		}
	}

	cart, err := c.cartRepository.GetCartByCustomerID(id)
	if err != nil {
		return echo.NewHTTPError(400, "Cart not found")
	}

	cartItemReq := models.CartItem{
		CartID:   int(cart.ID),
		MenuID:   int(menu.ID),
		Quantity: req.Quantity,
		TakeAway: req.TakeAway,
		Note:     req.Note,
	}

	err = c.cartRepository.CreateCartItem(&cartItemReq)
	if err != nil {
		return echo.NewHTTPError(400, "Failed to create cart item")
	}

	// update cart total price
	cart.TotalPrice += (menu.Price * req.Quantity)
	cart.TotalItem += req.Quantity
	err = c.cartRepository.UpdateCart(cart)
	if err != nil {
		return echo.NewHTTPError(400, "Failed to update cart")
	}

	return nil
}

// Logic Get Cart By Customer ID
func (c *cartUsecase) GetCartByCustomerID(id int) (*models.Cart, error) {
	cart, err := c.cartRepository.GetCartByCustomerID(id)
	if err != nil {
		return nil, echo.NewHTTPError(400, err.Error())
	}

	return cart, nil
}

// // Logic Update Cart Item
// func (c *cartUsecase) UpdateCartItem(id int, req *payload.UpdateCartItemRequest) error {
// 	cart, err := c.cartRepository.GetCartByCustomerID(id)
// 	if err != nil {
// 		return echo.NewHTTPError(400, err.Error())
// 	}

// 	cartItem, err := c.cartRepository.GetCartItemByID(req.CartItemID)
// 	if err != nil {
// 		return echo.NewHTTPError(400, err.Error())
// 	}

// 	cartItem.Quantity = req.Quantity
// 	cartItem.TakeAway = req.TakeAway
// 	cartItem.Note = req.Note

// 	err = c.cartRepository.UpdateCartItem(cartItem)
// 	if err != nil {
// 		return echo.NewHTTPError(400, err.Error())
// 	}

// 	// update cart total price
// 	_, err = c.menuRepository.GetMenuByID(int(cartItem.MenuID))
// 	if err != nil {
// 		return echo.NewHTTPError(400, err.Error())
// 	}

// 	err = c.cartRepository.UpdateCart(cart)
// 	if err != nil {
// 		return echo.NewHTTPError(400, err.Error())
// 	}

// 	return nil
// }

// Logic Delete Cart Item
func (c *cartUsecase) DeleteCartItem(id int, req *payload.DeleteCartItemRequest) error {
	cart, err := c.cartRepository.GetCartByCustomerID(id)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	cartItem, err := c.cartRepository.GetCartItemByID(req.CartItemID)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	err = c.cartRepository.DeleteCartItem(cartItem)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	// update cart total price
	_, err = c.menuRepository.GetMenuByID(int(cartItem.MenuID))
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	cart.TotalPrice -= (cartItem.Menu.Price * cartItem.Quantity)
	cart.TotalItem -= cartItem.Quantity

	err = c.cartRepository.UpdateCart(cart)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	// delete cart if total item is 0
	if cart.TotalItem == 0 {
		err = c.cartRepository.DeleteCart(cart)
		if err != nil {
			return echo.NewHTTPError(400, err.Error())
		}
	}

	return nil
}
