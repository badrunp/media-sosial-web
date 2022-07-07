package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/badrunp/media-sosial-web/be/controllers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHomeController(t *testing.T){

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")
	h := controllers.HomeController()

	if assert.NoError(t, h.Index(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello World!", rec.Body.String())
	}

}
