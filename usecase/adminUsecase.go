package usecase

import (
	"Chiprek/middleware"
	"Chiprek/models"
	"Chiprek/models/payload"
	"Chiprek/repository/database"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type AdminUsecase interface {
	LoginAdmin(req *payload.LoginAdminRequest) (res payload.LoginResponse, err error)
	DashboardAdmin() (res payload.DashboardAdminResponse, err error)
	GetMenuByID(id int) (menu *models.Menu, err error)
}

type adminUsecase struct {
	adminRepository       database.AdminRepository
	transactionRepository database.TransactionRepository
	menuRepository        database.MenuRepository
	customerRepository    database.CustomerRepository
}

func NewAdminUsecase(adminRepository database.AdminRepository, transactionRepository database.TransactionRepository, menuRepository database.MenuRepository, customerRepository database.CustomerRepository) *adminUsecase {
	return &adminUsecase{adminRepository, transactionRepository, menuRepository, customerRepository}
}

// Logic Login Admin
func (a *adminUsecase) LoginAdmin(req *payload.LoginAdminRequest) (res payload.LoginResponse, err error) {
	admin, err := a.adminRepository.GetAdminUsername(req.Username)
	if err != nil {
		echo.NewHTTPError(400, "Invalid Login")
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		echo.NewHTTPError(400, err.Error())
		return
	}

	token, err := middleware.CreateToken(int(admin.ID))
	if err != nil {
		echo.NewHTTPError(400, "Failed to generate token")
		return
	}

	admin.Token = token

	res = payload.LoginResponse{
		Username: admin.Username,
		Token:    token,
	}

	return
}

// Logic Dashboard Admin
func (a *adminUsecase) DashboardAdmin() (res payload.DashboardAdminResponse, err error) {
	totalMenus, err := a.menuRepository.TotalMenu()
	if err != nil {
		echo.NewHTTPError(400, "Failed to get total menu")
		return
	}

	totalOrders, err := a.transactionRepository.TotalOrder()
	if err != nil {
		echo.NewHTTPError(400, "Failed to get total transaction")
		return
	}

	totalCustomers, err := a.customerRepository.TotalCustomer()
	if err != nil {
		echo.NewHTTPError(400, "Failed to get total customer")
		return
	}

	totalIncome, err := a.transactionRepository.SumTransactionsAmount()
	if err != nil {
		echo.NewHTTPError(400, "Failed to get total income")
		return
	}

	res = payload.DashboardAdminResponse{
		TotalMenus:  uint(totalMenus),
		TotalOrders: uint(totalOrders),
		TotalUsers:  uint(totalCustomers),
		TotalIncome: uint(totalIncome),
	}

	return
}

// Get Menu By ID
func (a *adminUsecase) GetMenuByID(id int) (menu *models.Menu, err error) {
	menu, err = a.menuRepository.GetMenuByID(id)
	if err != nil {
		return nil, err
	}

	return menu, nil
}
