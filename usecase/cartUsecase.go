package usecase

import "Chiprek/repository/database"

type CartUsecase interface{}

type cartUsecase struct {
	cartRepository database.CartRepository
}

func NewCartUsecase(cartRepository database.CartRepository) *cartUsecase {
	return &cartUsecase{cartRepository}
}

// Logic Add Menu to Cart
