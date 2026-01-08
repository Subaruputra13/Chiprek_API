package database

import (
	"Chiprek/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	// GetCartByID(id int) (cart *models.Cart, err error)
	// GetCartItemByID(id int) (cartItem *models.CartItem, err error)
	GetCartByCustomerID(customerID int) (cart *models.Cart, err error)
	GetCartItemByID(item int) (cartItem *models.CartItem, err error)
	CreateCart(cart *models.Cart) error
	CreateCartItem(cartItem *models.CartItem) error
	UpdateCart(cart *models.Cart) error
	UpdateCartItem(cartItem *models.CartItem) error
	DeleteCart(cart *models.Cart) error
	DeleteCartItem(cartItem *models.CartItem) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *cartRepository {
	return &cartRepository{db}
}

// Get cart by ID
// func (c *cartRepository) GetCartByID(id int) (cart *models.Cart, err error) {
// 	err = c.db.Where("id = ?", id).First(&cart).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return cart, nil
// }

// Get cart item by ID
// func (c *cartRepository) GetCartItemByID(id int) (cartItem *models.CartItem, err error) {
// 	err = c.db.Where("id = ?", id).First(&cartItem).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return cartItem, nil
// }

// Get Cart By Customer ID
func (c *cartRepository) GetCartByCustomerID(customerID int) (cart *models.Cart, err error) {
	err = c.db.Preload("CartItem.Menu").Preload("Customer").Where("customer_id = ? AND status = ?", customerID, true).First(&cart).Error
	if err != nil {
		return nil, err
	}

	return cart, nil
}

// Get Cart Item By Cart ID
func (c *cartRepository) GetCartItemByID(item int) (cartItem *models.CartItem, err error) {
	err = c.db.Preload("Menu").Where("id = ?", item).First(&cartItem).Error
	if err != nil {
		return nil, err
	}

	return cartItem, nil
}

// Create cart
func (c *cartRepository) CreateCart(cart *models.Cart) error {
	err := c.db.Create(&cart).Error
	if err != nil {
		return err
	}

	return nil
}

// Create cart item
func (c *cartRepository) CreateCartItem(cartItem *models.CartItem) error {
	err := c.db.Create(&cartItem).Error
	if err != nil {
		return err
	}

	return nil
}

// // Get Cart by Status Active
// func (c *cartRepository) GetCartByStatusActive() (*models.Cart, error) {
// 	var cart models.Cart
// 	err := c.db.Where("status = ?", true).First(&cart).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &cart, nil
// }

// Update cart
func (c *cartRepository) UpdateCart(cart *models.Cart) error {
	err := c.db.Updates(&cart).Error
	if err != nil {
		return err
	}

	return nil
}

// Update cart item
func (c *cartRepository) UpdateCartItem(cartItem *models.CartItem) error {
	err := c.db.Updates(&cartItem).Error
	if err != nil {
		return err
	}

	return nil
}

// Delete cart
func (c *cartRepository) DeleteCart(cart *models.Cart) error {
	err := c.db.Delete(&cart).Error
	if err != nil {
		return err
	}

	return nil
}

// Delete cart item
func (c *cartRepository) DeleteCartItem(cartItem *models.CartItem) error {
	err := c.db.Delete(&cartItem).Error
	if err != nil {
		return err
	}

	return nil
}
