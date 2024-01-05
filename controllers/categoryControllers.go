package controllers

import (
	"Chiprek/models/payload"
	"Chiprek/usecase"
	"strconv"

	"github.com/labstack/echo"
)

type CategoryControllers interface{}

type categoryControllers struct {
	categoryUsecase usecase.CategoryUsecase
}

func NewCategoryControllers(categoryUsecase usecase.CategoryUsecase) *categoryControllers {
	return &categoryControllers{categoryUsecase}
}

// Controller Get All Category
func (cc *categoryControllers) GetAllCategoryController(c echo.Context) error {
	category, err := cc.categoryUsecase.GetAllCategory()
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Get All Category",
		Data:    category,
	})
}

// Controller Get Category By ID
func (cc *categoryControllers) GetCategoryByIDController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	category, err := cc.categoryUsecase.GetCategoryByID(id)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Get Category By ID",
		Data:    category,
	})
}
