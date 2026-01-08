package controllers

import (
	"Chiprek/models/payload"
	"Chiprek/usecase"
	"strconv"

	"github.com/labstack/echo"
)

type AdminController interface {
	LoginAdminController(c echo.Context) error
}

type adminController struct {
	adminUsecase usecase.AdminUsecase
}

func NewAdminController(adminUsecase usecase.AdminUsecase) *adminController {
	return &adminController{adminUsecase}
}

// Controller Login Admin
func (a *adminController) LoginAdminController(c echo.Context) error {
	req := payload.LoginAdminRequest{}

	c.Bind(&req)
	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(400, "Field tidak boleh kosong")
	}

	res, err := a.adminUsecase.LoginAdmin(&req)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Login Berhasil",
		Data:    res,
	})
}

// Controller Dashboard Admin
func (a *adminController) DashboardAdminController(c echo.Context) error {
	res, err := a.adminUsecase.DashboardAdmin()
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Get Dashboard Admin",
		Data:    res,
	})
}

// Controller Get Menu By ID
func (a *adminController) GetMenuByIDController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	menu, err := a.adminUsecase.GetMenuByID(id)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Get Menu By ID",
		Data:    menu,
	})
}
