package controllers

import (
	"Chiprek/models/payload"
	"Chiprek/usecase"

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
