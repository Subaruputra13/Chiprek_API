package faker

import (
	"Chiprek/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AdminFaker(db *gorm.DB) *models.Admin {
	passwordhash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)

	return &models.Admin{
		Password: string(passwordhash),
	}
}
