package configs

import (
	"github.com/badrunp/media-sosial-web/be/helpers"
)

type App interface{}

type app struct {
	Name string
	Port int
}

func AppConfig() *app {
	return &app{
		Name: "Backend",
		Port: helpers.GetEnvNum("PORT", 8000),
	}
}