package usecase

import (
	"Chiprek/middleware"
	"Chiprek/models/payload"
	"Chiprek/repository/database"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type AdminUsecase interface {
	LoginAdmin(req *payload.LoginAdminRequest) (res payload.LoginResponse, err error)
}

type adminUsecase struct {
	AdminRepository database.AdminRepository
}

func NewAdminUsecase(adminRepository database.AdminRepository) *adminUsecase {
	return &adminUsecase{adminRepository}
}

// Logic Login Admin
func (a *adminUsecase) LoginAdmin(req *payload.LoginAdminRequest) (res payload.LoginResponse, err error) {
	admin, err := a.AdminRepository.GetAdminUsername(req.Username)
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
