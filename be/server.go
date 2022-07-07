package main

import (
	"fmt"
	"os"

	"github.com/badrunp/media-sosial-web/be/configs"
	migration "github.com/badrunp/media-sosial-web/be/database/migrations"
	seed "github.com/badrunp/media-sosial-web/be/database/seeds"
	"github.com/badrunp/media-sosial-web/be/helpers"
	"github.com/badrunp/media-sosial-web/be/routes"
	"github.com/badrunp/media-sosial-web/be/utils"
)

func main() {
	helpers.LoadEnv(".env")
	db, _ := utils.DBConnect(configs.DatabaseConfig().Host)

	args := os.Args
	migration.RegisterMigration(db, args)
	seed.RegisterSeeders(db, args)

	r := routes.NewRoutes(db)
	var port int = configs.AppConfig().Port 
	var start string = fmt.Sprintf(":%d", port)
	r.Init().Logger.Fatal(r.Init().Start(start))
}