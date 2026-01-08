package controllers

import (
	"Chiprek/models/payload"
	"Chiprek/util"
	"net/http"

	"github.com/labstack/echo"
)

func UploadImageCloudBase64Controller(c echo.Context) error {

	payloadImage := payload.UploadImageCloudinaryBase64{}
	c.Bind(&payloadImage)

	resp, err := util.UploadImageCloudBase64(&payloadImage)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error To upload Image")
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "image upload success",
		Data:    resp,
	})
}
