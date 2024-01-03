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

	//Menu
	menuRepository := database.NewMenuRepository(db)
	menuUsecase := usecase.NewMenuUsecase(menuRepository)
	menuController := controllers.NewMenuControllers(menuUsecase)

	//Admin Route
	e.POST("/admin", adminController.LoginAdminController)
	m := e.Group("/Dashboard/Menu", m.IsLoggedIn)
	m.GET("", menuController.GetAllMenuController)
	m.POST("", menuController.CreateMenuController)
	m.PUT("/:id", menuController.UpdateMenuController)
	m.DELETE("/:id", menuController.DeleteMenuController)
}
