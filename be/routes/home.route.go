package routes

import (
	"github.com/badrunp/media-sosial-web/be/controllers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Home interface{
	Init()*echo.Group
}

type home struct {
	e *echo.Group
	db *gorm.DB
}

func HomeRoute(e *echo.Group, db *gorm.DB)*home{
	return &home{e: e, db: db}
}

func (h *home) Init()*echo.Group{

	homeController := controllers.HomeController()

	h.e.GET("", homeController.Index)

	return h.e
}