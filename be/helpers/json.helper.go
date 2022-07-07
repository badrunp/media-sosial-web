package helpers

import (
	"github.com/badrunp/media-sosial-web/be/response"
	"github.com/labstack/echo/v4"
)

func ResJSON(c echo.Context, status int, message string, data interface{}, err interface{}) error {
	res := response.User{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   err,
	}

	return c.JSON(status, res)
}