package routes

import (
	"Chiprek/controllers"
	m "Chiprek/middleware"
	"Chiprek/repository/database"
	"Chiprek/usecase"
	"Chiprek/util"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {
	//Middleware
	m.LogMiddleware(e)
	e.Pre(mid.RemoveTrailingSlash())

	//Validator
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	//Admin
	adminRepository := database.NewAdminRepository(db)
	adminUsecase := usecase.NewAdminUsecase(adminRepository)
	adminController := controllers.NewAdminController(adminUsecase)

	//Admin Route
	e.POST("/admin", adminController.LoginAdminController)
}
