package routes

import (
	"github.com/badrunp/media-sosial-web/be/controllers"
	"github.com/badrunp/media-sosial-web/be/repositories"
	service "github.com/badrunp/media-sosial-web/be/services"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Auth interface{
	Init()*echo.Group
}

type auth struct {
	e *echo.Group
	db *gorm.DB
}

func AuthRoute(e *echo.Group, db *gorm.DB)*auth{
	return &auth{e: e, db: db}
}

func (a *auth) Init()*echo.Group {
	
	repository := repositories.UserRepository(a.db)
	service := service.UserService(repository)
	authController := controllers.AuthController(service)

	g := a.e.Group("/auth")
	g.POST("/register", authController.Register)
	g.POST("/login", authController.Login)

	return g
}