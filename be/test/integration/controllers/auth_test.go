package test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/badrunp/media-sosial-web/be/configs"
	"github.com/badrunp/media-sosial-web/be/controllers"
	migrations "github.com/badrunp/media-sosial-web/be/database/migrations"
	"github.com/badrunp/media-sosial-web/be/entity"
	"github.com/badrunp/media-sosial-web/be/helpers"
	"github.com/badrunp/media-sosial-web/be/repositories"
	service "github.com/badrunp/media-sosial-web/be/services"
	"github.com/badrunp/media-sosial-web/be/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.M){
	helpers.LoadEnv("../../../.env")

	db, _ := utils.DBConnectTest(configs.TestDatabaseConfig().Host)
	migrations.AddMigrations(db, "")
	repository := repositories.UserRepository(db)
	repository.DeleteAll()
	t.Run()
	repository.DeleteAll()
}

func TestRegisterController(t *testing.T){
	f := make(url.Values)
	f.Set("username", "badrun")
	f.Set("email", "bbadrunn@gmail.com")
	f.Set("password", "123456")

	e := echo.New()
	e.Validator = &utils.Validator{Validation: validator.New()}

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/auth/register")

	db, err := utils.DBConnectTest(configs.TestDatabaseConfig().Host)
	if assert.NoError(t, err) {
		repository := repositories.UserRepository(db)
		if assert.NoError(t, err) {
			service := service.UserService(repository)
			rc := controllers.AuthController(service)
	
			if assert.NoError(t, rc.Register(c)) {
				assert.Equal(t, http.StatusOK, rec.Code)
				log.Println(rec.Body.String())
			}
		}
	}

}

func TestLoginController(t *testing.T){
	f := make(url.Values)
	f.Set("email", "bbadrunn@gmail.com")
	f.Set("password", "123456")

	e := echo.New()
	e.Validator = &utils.Validator{Validation: validator.New()}

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/auth/login")

	db, err := utils.DBConnectTest(configs.TestDatabaseConfig().Host)
	if assert.NoError(t, err) {
		repository := repositories.UserRepository(db)
		repository.Create(entity.User{Username: "badrun", Email: "bbadrunn@gmail.com", Password: utils.HashAndSalt([]byte("123456"))})
		if assert.NoError(t, err) {
			service := service.UserService(repository)
			rc := controllers.AuthController(service)
	
			if assert.NoError(t, rc.Login(c)) {
				assert.Equal(t, http.StatusOK, rec.Code)
				log.Println(rec.Body.String())
			}
		}
	}

}