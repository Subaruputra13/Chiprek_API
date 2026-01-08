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
	e.Use(mid.CORS())

	//Validator
	e.Validator = &util.CustomValidator{Validator: validator.New()}

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

	//Transaction
	transactionRepository := database.NewTransactionRepository(db)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepository, cartRepository)
	transactionController := controllers.NewTransactionController(transactionUsecase)

	//Admin
	adminRepository := database.NewAdminRepository(db)
	adminUsecase := usecase.NewAdminUsecase(adminRepository, transactionRepository, menuRepository, customerRepository)
	adminController := controllers.NewAdminController(adminUsecase)

	// Auth Route
	e.POST("/admin", adminController.LoginAdminController)
	e.POST("/notification", transactionController.GetNotificationController)

	//Admin Route
	me := e.Group("/dashboard", m.IsLoggedIn)
	me.GET("", adminController.DashboardAdminController)
	me.GET("/menu", menuController.GetAllMenuController)
	me.GET("/menu/:id", menuController.GetMenuByIDController)
	me.POST("/menu", menuController.CreateMenuController)
	me.PUT("/menu/:id", menuController.UpdateMenuController)
	me.DELETE("/menu/:id", menuController.DeleteMenuController)
	me.GET("/transaction", transactionController.GetAllTransactionController)
	me.PUT("/transaction/:id", transactionController.UpdateTransactionByIdController)
	me.GET("/transaction/:id", transactionController.GetTransactionByIdController)

	//Customer Route
	cu := e.Group("/customer")
	cu.POST("", customerController.CreateCustomerControllers)
	cu.GET("/cart", cartController.GetCartByCustomerIDControllers, m.IsLoggedIn)
	cu.POST("/cart", cartController.AddMenuToCartControllers, m.IsLoggedIn)
	cu.POST("/transaction", transactionController.CreateTransactionController, m.IsLoggedIn)
	cu.GET("/transaction", transactionController.GetTransactionByCustomerIdController, m.IsLoggedIn)
	cu.DELETE("/cart", cartController.DeleteCartItemControllers, m.IsLoggedIn)

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
