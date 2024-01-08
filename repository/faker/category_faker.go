package faker

import (
	"Chiprek/models"

	"gorm.io/gorm"
)

func CategoryFaker(db *gorm.DB) *models.Category {
	category := []models.Category{
		{Name: "Makanan"},
		{Name: "Minuman"},
		{Name: "Tambahan"},
	}

	for _, v := range category {
		var exist models.Category
		err := db.Where("nama = ?", v.Name).First(&exist).Error
		if err != nil {
			db.Create(&v)
		}
	}

	return nil
}
