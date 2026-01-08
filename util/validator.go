package util

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.Validator.Struct(i)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			// Jika terjadi kesalahan validasi internal
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}

		validationErrors := err.(validator.ValidationErrors)
		errorMessage := ""
		for _, fieldError := range validationErrors {
			// Menyesuaikan pesan error menjadi pesan yang lebih informatif
			switch fieldError.Tag() {
			case "required":
				errorMessage += fmt.Sprintf(`Kolom %s wajib diisi`+"\n", fieldError.Field())
			case "email":
				errorMessage += fmt.Sprintf("Kolom %s harus berupa alamat email yang valid"+"\n", fieldError.Field())
			case "min":
				errorMessage += fmt.Sprintf("Kolom %s minimal %s karakter"+"\n", fieldError.Field(), fieldError.Param())
			case "max":
				errorMessage += fmt.Sprintf("Kolom %s maksimal %s karakter"+"\n", fieldError.Field(), fieldError.Param())
			case "number":
				errorMessage += fmt.Sprintf("Kolom %s harus berupa angka"+"\n", fieldError.Field())
			// case "len":
			// 	errorMessage += fmt.Sprintf("Field %s must be %s length"+"\n", fieldError.Field(), fieldError.Param())
			default:
				errorMessage += fmt.Sprintf("Kolom %s tidak valid"+"\n", fieldError.Field())
			}
		}
		return errors.New(errorMessage)
	}
	return nil
}
