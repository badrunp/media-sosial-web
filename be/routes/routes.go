package routes

import (
	"net/http"

	"github.com/badrunp/media-sosial-web/be/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type Routes interface {
	Init() *echo.Echo
}

type routes struct{
	db *gorm.DB
}

func NewRoutes(db *gorm.DB)*routes{
	return &routes{db: db}
}

func (r *routes) Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	
	e.Validator = &utils.Validator{Validation: validator.New()}
	g := e.Group("/api")
	HomeRoute(g, r.db).Init()
	AuthRoute(g, r.db).Init()
	UserRoute(g, r.db).Init()

	return e
}