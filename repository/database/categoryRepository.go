package database

import (
	"Chiprek/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategory() (category []models.Category, err error)
	GetCategoryByID(id int) (category *models.Category, err error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

// Get All Category
func (c *categoryRepository) GetAllCategory() (category []models.Category, err error) {

	if err := c.db.Preload("Menu").Find(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

// Get Category By ID
func (c *categoryRepository) GetCategoryByID(id int) (category *models.Category, err error) {
	if err := c.db.Preload("Menu").Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}
