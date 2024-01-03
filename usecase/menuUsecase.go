package usecase

import (
	"Chiprek/models"
	"Chiprek/models/payload"
	"Chiprek/repository/database"

	"github.com/labstack/echo"
)

type MenuUsecase interface {
	GetAllMenu() (menu []models.Menu, err error)
	GetMenuByID(id int) (menu *models.Menu, err error)
	CreateMenu(req *payload.CreateMenuRequest) (menu *models.Menu, err error)
	UpdateMenu(id int, req *payload.UpdateMenuRequest) (*models.Menu, error)
	DeleteMenu(menu *models.Menu) error
}

type menuUsecase struct {
	menuRespository database.MenuRepository
}

func NewMenuUsecase(menuRespository database.MenuRepository) *menuUsecase {
	return &menuUsecase{menuRespository}
}

// Get All Menu
func (m *menuUsecase) GetAllMenu() (menu []models.Menu, err error) {
	menu, err = m.menuRespository.GetAllMenu()
	if err != nil {
		return nil, err
	}

	return menu, nil
}

// Get Menu By ID
func (m *menuUsecase) GetMenuByID(id int) (menu *models.Menu, err error) {
	menu, err = m.menuRespository.GetMenuByID(id)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

// Create Menu
func (m *menuUsecase) CreateMenu(req *payload.CreateMenuRequest) (menu *models.Menu, err error) {
	menuReq := &models.Menu{
		Nama:  req.Nama,
		Harga: req.Harga,
	}

	if err := m.menuRespository.CreateMenu(menuReq); err != nil {
		return nil, echo.NewHTTPError(400, err.Error())
	}

	return menuReq, nil
}

// Update Menu
func (m *menuUsecase) UpdateMenu(id int, req *payload.UpdateMenuRequest) (*models.Menu, error) {
	menu, err := m.menuRespository.GetMenuByID(id)
	if err != nil {
		return nil, echo.NewHTTPError(400, err.Error())
	}

	menu.Nama = req.Nama
	menu.Harga = req.Harga

	err = m.menuRespository.UpdateMenu(menu)
	if err != nil {
		return nil, echo.NewHTTPError(400, err.Error())
	}

	return menu, nil
}

// Delete Menu
func (m *menuUsecase) DeleteMenu(menu *models.Menu) error {
	if err := m.menuRespository.DeleteMenu(menu); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return nil
}
