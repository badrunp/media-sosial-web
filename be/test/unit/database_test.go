package test

import (
	"testing"

	"github.com/badrunp/media-sosial-web/be/configs"
	migration "github.com/badrunp/media-sosial-web/be/database/migrations"
	seed "github.com/badrunp/media-sosial-web/be/database/seeds"
	"github.com/badrunp/media-sosial-web/be/helpers"
	"github.com/badrunp/media-sosial-web/be/utils"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.M){
	helpers.LoadEnv("../../.env")
	t.Run()
}

func TestDatabaseConection(t *testing.T){
	_ , err := utils.DBConnectTest(configs.DatabaseConfig().Host)

	assert.NoError(t, err)
}

func TestDatabaseConnectionError(t *testing.T){
	_ , err := utils.DBConnectTest("error")

	assert.Error(t, err)
	assert.EqualError(t, err, "failed to connect to `host=error user=gorm database=sosial_media_test`: hostname resolving error (lookup error: no such host)")
}

func TestDatabaseMigration(t *testing.T){
	db, _ := utils.DBConnectTest(configs.DatabaseConfig().Host)

	result := migration.RegisterMigration(db, []string{"testing","migrate:up"})
	assert.Equal(t, "migration success", result)
}

func TestDatabaseMigrationFailed(t *testing.T){
	db, _ := utils.DBConnectTest(configs.DatabaseConfig().Host)

	result := migration.RegisterMigration(db, []string{"testing","migrates:up"})
	assert.Equal(t, "migration failed", result)
}

func TestDatabaseSeed(t *testing.T){
	db, _ := utils.DBConnectTest(configs.DatabaseConfig().Host)

	result := seed.RegisterSeeders(db, []string{"testing","seed:run"})
	assert.Equal(t, "seed success", result)
}

func TestDatabaseSeedFailed(t *testing.T){
	db, _ := utils.DBConnectTest(configs.DatabaseConfig().Host)

	result := seed.RegisterSeeders(db, []string{"testing","seeds:up"})
	assert.Equal(t, "seed failed", result)
}