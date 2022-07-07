package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Home interface {
	Index(c echo.Context) error
}

type home struct {}

func HomeController()*home{
	return &home{}
}

func (h *home) Index(c echo.Context) error {
	return c.String(http.StatusOK,"Hello World!")
}