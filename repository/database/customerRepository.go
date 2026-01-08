package database

import (
	"Chiprek/models"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	CreateCustomer(customer *models.Customer) (models.Customer, error)
	UpdateCustomer(customer *models.Customer) (*models.Customer, error)
	DeleteCustomer(customer *models.Customer) error
	TotalCustomer() (total int64, err error)
}

type customerRespository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *customerRespository {
	return &customerRespository{db}
}

// Create Customer
func (c *customerRespository) CreateCustomer(customer *models.Customer) (models.Customer, error) {
	if err := c.db.Create(&customer).Error; err != nil {
		return models.Customer{}, err
	}

	return *customer, nil
}

// Update Customer
func (c *customerRespository) UpdateCustomer(customer *models.Customer) (*models.Customer, error) {
	if err := c.db.Updates(&customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

// Delete Customer
func (c *customerRespository) DeleteCustomer(customer *models.Customer) error {
	if err := c.db.Delete(&customer).Error; err != nil {
		return err
	}

	return nil
}

// Total Customer
func (c *customerRespository) TotalCustomer() (total int64, err error) {
	Customers := []models.Customer{}
	if err := c.db.Model(&Customers).Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}
