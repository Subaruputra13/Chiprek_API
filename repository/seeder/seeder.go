package seeder

import (
	"Chiprek/repository/faker"

	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

func RegisterSeeder(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: faker.AdminFaker(db)},
		{Seeder: faker.CategoryFaker(db)},
	}
}

func DBSeed(db *gorm.DB) error {
	for _, seeder := range RegisterSeeder(db) {
		err := db.Debug().Create(seeder.Seeder).Error
		if err != nil {
			return err
		}
	}
	return nil
}
