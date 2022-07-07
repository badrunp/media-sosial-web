package controllers

import (
	"net/http"

	"github.com/badrunp/media-sosial-web/be/helpers"
	"github.com/badrunp/media-sosial-web/be/services"
	"github.com/labstack/echo/v4"
)

type User interface{
	Register(c echo.Context) error
}

type user struct{
	userService services.User
}

func UserController(userService services.User) *user {
	return &user{userService: userService}
}

func (a *user) Index(c echo.Context) error {
	users, err := a.userService.GetAllUser()
	if err != nil {
		return helpers.ResJSON(c, http.StatusBadRequest, "Failed get all users!", nil, err)
	}
	return helpers.ResJSON(c, http.StatusOK, "Get all users success", users, nil)
}
