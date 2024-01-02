package database

import (
	"Chiprek/config"
	"Chiprek/models"

	"gorm.io/gorm"
)

type AdminRepository interface {
	LoginAdmin(admin *models.Admin) error
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
