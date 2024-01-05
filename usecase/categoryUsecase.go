package usecase

import (
	"Chiprek/models"
	"Chiprek/repository/database"
)

type CategoryUsecase interface {
	GetAllCategory() (category []models.Category, err error)
	GetCategoryByID(id int) (category *models.Category, err error)
}

type categoryUsecase struct {
	categoryRepository database.CategoryRepository
}

func NewCategoryUsecase(categoryRepository database.CategoryRepository) *categoryUsecase {
	return &categoryUsecase{categoryRepository}
}

// Get All Category
func (c *categoryUsecase) GetAllCategory() (category []models.Category, err error) {
	category, err = c.categoryRepository.GetAllCategory()
	if err != nil {
		return nil, err
	}

	return category, nil
}

// Get Category By ID
func (c *categoryUsecase) GetCategoryByID(id int) (category *models.Category, err error) {
	category, err = c.categoryRepository.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}
