package database

import (
	"Chiprek/models"

	"gorm.io/gorm"
)

type MenuRepository interface {
	GetAllMenu() (menu []models.Menu, err error)
	GetMenuByID(id int) (menu *models.Menu, err error)
	CreateMenu(menu *models.Menu) error
	UpdateMenu(menu *models.Menu) (*models.Menu, error)
	DeleteMenu(menu *models.Menu) error
	TotalMenu() (total int64, err error)
}

type menuRespository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *menuRespository {
	return &menuRespository{db}
}

// Get All Menu
func (m *menuRespository) GetAllMenu() (menu []models.Menu, err error) {

	if err := m.db.Find(&menu).Error; err != nil {
		return nil, err
	}

	return menu, nil
}

// Get Menu By ID
func (m *menuRespository) GetMenuByID(id int) (menu *models.Menu, err error) {
	if err := m.db.Where("id = ?", id).First(&menu).Error; err != nil {
		return nil, err
	}

	return menu, nil
}

// Create Menu
func (m *menuRespository) CreateMenu(menu *models.Menu) error {
	if err := m.db.Create(&menu).Error; err != nil {
		return err
	}

	return nil
}

// Update Menu
func (m *menuRespository) UpdateMenu(menu *models.Menu) (*models.Menu, error) {
	if err := m.db.Updates(&menu).Error; err != nil {
		return nil, err
	}

	return menu, nil
}

// Delete Menu
func (m *menuRespository) DeleteMenu(menu *models.Menu) error {
	if err := m.db.Delete(&menu).Error; err != nil {
		return err
	}

	return nil
}

// Total Menus
func (m *menuRespository) TotalMenu() (total int64, err error) {
	Menus := []models.Menu{}
	if err := m.db.Model(&Menus).Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}
