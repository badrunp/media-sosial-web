package controllers

import (
	"net/http"

	"github.com/badrunp/media-sosial-web/be/entity"
	"github.com/badrunp/media-sosial-web/be/helpers"
	"github.com/badrunp/media-sosial-web/be/requests"
	"github.com/badrunp/media-sosial-web/be/services"
	"github.com/badrunp/media-sosial-web/be/utils"
	"github.com/labstack/echo/v4"
)

type Auth interface{
	Register(c echo.Context) error
	Login(c echo.Context) error
}

type auth struct{
	userService services.User
}

func AuthController(userService services.User) *auth {
	return &auth{userService: userService}
}

func (a *auth) Register(c echo.Context) error {
	var (
		data requests.User
	 	user entity.User
	 	err error
	)

	if err := c.Bind(&data); err != nil {
		return helpers.ResJSON(c, http.StatusBadRequest, "Bind error", nil, err)
	}

	if err := c.Validate(data); err != nil {
		return helpers.ResJSON(c, http.StatusBadRequest, "Validation", nil, err)
	}

	data.Password = utils.HashAndSalt([]byte(data.Password))

	user, err = a.userService.CreateUser(data)
	if err != nil {
		return helpers.ResJSON(c, http.StatusBadRequest, "Create user", nil, err)
	}

	return helpers.ResJSON(c, http.StatusOK, "Register Success", user, nil)
}

func (a *auth) Login(c echo.Context) error {
	var (
		data requests.UserLogin
		err error
		user entity.User
	)

	if err = c.Bind(&data); err != nil {
		return helpers.ResJSON(c, http.StatusBadRequest, "Bind error!", nil, err)
	}

	if err := c.Validate(data); err != nil {
		return helpers.ResJSON(c, http.StatusBadRequest, "Validation", nil, err)
	}

	user, err = a.userService.GetUserByEmail(data)
	if err != nil {
		return helpers.ResJSON(c, http.StatusBadRequest, "Email error", nil, "Email not register in sytem!")
	}
	
	if !utils.ComparePasswords(user.Password, []byte(data.Password)){
		return helpers.ResJSON(c, http.StatusBadRequest, "Password error", nil, "Password is not match!")
	}

	return helpers.ResJSON(c, http.StatusOK, "Login success", user, nil)
}