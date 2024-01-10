package database

import (
	"Chiprek/models"

	"gorm.io/gorm"
)

type CartRepository interface {
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
