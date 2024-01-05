package faker

import (
	"Chiprek/models"

	"gorm.io/gorm"
)

func CategoryFaker(db *gorm.DB) *models.Category {
	category := []models.Category{
		{Nama: "Makanan"},
		{Nama: "Minuman"},
		{Nama: "Tambahan"},
	}

	for _, v := range category {
		var exist models.Category
		err := db.Where("nama = ?", v.Nama).First(&exist).Error
		if err != nil {
			db.Create(&v)
		}
	}

	return nil
}
