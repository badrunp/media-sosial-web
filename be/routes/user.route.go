package routes

import (
	"github.com/badrunp/media-sosial-web/be/controllers"
	"github.com/badrunp/media-sosial-web/be/repositories"
	"github.com/badrunp/media-sosial-web/be/services"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type User interface{
	Init()*echo.Group
}

type user struct {
	e *echo.Group
	db *gorm.DB
}

func UserRoute(e *echo.Group, db *gorm.DB)*user{
	return &user{e: e, db: db}
}

func (h *user) Init()*echo.Group{

	userRepository := repositories.UserRepository(h.db)
	userService := services.UserService(userRepository)
	homeController := controllers.UserController(userService)

	h.e.GET("/users", homeController.Index)

	return h.e
}