package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/badrunp/media-sosial-web/be/configs"
	"github.com/badrunp/media-sosial-web/be/controllers"
	migrations "github.com/badrunp/media-sosial-web/be/database/migrations"
	"github.com/badrunp/media-sosial-web/be/repositories"
	"github.com/badrunp/media-sosial-web/be/services"
	"github.com/badrunp/media-sosial-web/be/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/users")

	db, err := utils.DBConnectTest(configs.TestDatabaseConfig().Host)
	migrations.AddMigrations(db, "")
	if assert.NoError(t, err) {

		userRepository := repositories.UserRepository(db)
		userService := services.UserService(userRepository)
		userController := controllers.UserController(userService)

		if assert.NoError(t, userController.Index(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			fmt.Println(rec.Body.String())
		}

	}

}