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

	//Customer
	customerRepository := database.NewCustomerRepository(db)
	customerUsecase := usecase.NewCustomerUsecase(customerRepository)
	customerController := controllers.NewCustomerControllers(customerUsecase)

	//Menu
	menuRepository := database.NewMenuRepository(db)
	menuUsecase := usecase.NewMenuUsecase(menuRepository)
	menuController := controllers.NewMenuControllers(menuUsecase)

	//Category
	categoryRepository := database.NewCategoryRepository(db)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository)
	categoryController := controllers.NewCategoryControllers(categoryUsecase)

	//Cart
	cartRepository := database.NewCartRepository(db)
	cartUsecase := usecase.NewCartUsecase(cartRepository, menuRepository)
	cartController := controllers.NewCartControllers(cartUsecase)

	// Auth Route
	e.POST("/admin", adminController.LoginAdminController)

	//Admin Route
	me := e.Group("/dashboard/menu", m.IsLoggedIn)
	me.GET("", menuController.GetAllMenuController)
	me.POST("", menuController.CreateMenuController)
	me.PUT("/:id", menuController.UpdateMenuController)
	me.DELETE("/:id", menuController.DeleteMenuController)

	//Customer Route
	cu := e.Group("/customer", m.IsLoggedIn)
	cu.POST("", customerController.CreateCustomerControllers)
	cu.GET("/cart", cartController.GetCartByCustomerIDControllers)
	cu.POST("/cart", cartController.AddMenuToCartControllers)

	//Menu Route
	e.GET("/menu", menuController.GetAllMenuController)
	e.GET("/menu/:id", menuController.GetMenuByIDController)

	//Image Upload Route
	e.POST("/upload/image", controllers.UploadImageCloudBase64Controller)

	//Category Route
	c := e.Group("/category")
	c.GET("", categoryController.GetAllCategoryController)
	c.GET("/:id", categoryController.GetCategoryByIDController)

}
