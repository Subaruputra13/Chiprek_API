package database

import (
	"Chiprek/config"
	"Chiprek/models"

	"gorm.io/gorm"
)

type AdminRepository interface {
	LoginAdmin(admin *models.Admin) error
	GetAdminUsername(username string) (*models.Admin, error)
}

type adminRespository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRespository {
	return &adminRespository{db}
}

// Login Admin
func (a *adminRespository) LoginAdmin(admin *models.Admin) error {
	if err := config.DB.Where("password = ?", admin.Password).First(&admin).Error; err != nil {
		return err
	}

	return nil
}

// Get Admin
func (a *adminRespository) GetAdminUsername(username string) (*models.Admin, error) {
	var admin models.Admin

	if err := config.DB.Where("username = ?", username).First(&admin).Error; err != nil {
		return nil, err
	}

	return &admin, nil
}

// Update Transaction
func (a *adminRespository) UpdateTransaction(transaction *models.Transaction) error {
	if err := config.DB.Updates(&transaction).Error; err != nil {
		return err
	}

	return nil
}
