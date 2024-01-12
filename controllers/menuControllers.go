package controllers

import (
	"Chiprek/models/payload"
	"Chiprek/usecase"
	"strconv"

	"github.com/labstack/echo"
)

type MenuControllers interface {
	GetAllMenuController(c echo.Context) error
	CreateMenuController(c echo.Context) error
	UpdateMenuController(c echo.Context) error
	DeleteMenuController(c echo.Context) error
}

type menuControllers struct {
	menuUsecase usecase.MenuUsecase
}

func NewMenuControllers(menuUsecase usecase.MenuUsecase) *menuControllers {
	return &menuControllers{menuUsecase}
}

// Controller Get All Menu
func (m *menuControllers) GetAllMenuController(c echo.Context) error {
	menu, err := m.menuUsecase.GetAllMenu()
	if err != nil {
		return echo.NewHTTPError(400, err.Error())

	}

	return c.JSON(200, payload.Response{
		Message: "Success Get All Menu",
		Data:    menu,
	})
}

// Controller Get All Menu
// func (m *menuControllers) GetAllMenuControllerByTableNumber(c echo.Context) error {
// 	menu, err := m.menuUsecase.GetAllMenu()
// 	if err != nil {
// 		return echo.NewHTTPError(400, err.Error())

// 	}

// 	return c.JSON(200, payload.Response{
// 		Message: "Success Get All Menu",
// 		Data:    menu,
// 	})
// }

// Controller Get Menu By ID
func (m *menuControllers) GetMenuByIDController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	menu, err := m.menuUsecase.GetMenuByID(int(id))
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Get Menu By ID",
		Data:    menu,
	})
}

// Controller Create Menu
func (m *menuControllers) CreateMenuController(c echo.Context) error {
	req := payload.CreateMenuRequest{}

	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(400, "Field tidak boleh kosong")
	}

	menu, err := m.menuUsecase.CreateMenu(&req)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Create Menu",
		Data:    menu,
	})

}

// Controller Update Menu
func (m *menuControllers) UpdateMenuController(c echo.Context) error {
	req := payload.UpdateMenuRequest{}

	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(400, "Field tidak boleh kosong")
	}

	id, _ := strconv.Atoi(c.Param("id"))

	menu, err := m.menuUsecase.UpdateMenu(id, &req)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Update Menu",
		Data:    menu,
	})
}

// Controller Delete Menu
func (m *menuControllers) DeleteMenuController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	menu, err := m.menuUsecase.GetMenuByID(int(id))
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	err = m.menuUsecase.DeleteMenu(menu)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, "Delete Menu Success")
}
