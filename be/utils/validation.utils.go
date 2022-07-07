package utils

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Validator struct {
	Validation *validator.Validate
}

func (cv *Validator) Validate(i interface{}) error {
  if err := cv.Validation.Struct(i); err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return nil
}