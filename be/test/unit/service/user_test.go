package test

import (
	"testing"

	"github.com/badrunp/media-sosial-web/be/configs"
	migrations "github.com/badrunp/media-sosial-web/be/database/migrations"
	"github.com/badrunp/media-sosial-web/be/entity"
	"github.com/badrunp/media-sosial-web/be/helpers"
	"github.com/badrunp/media-sosial-web/be/repositories"
	"github.com/badrunp/media-sosial-web/be/requests"
	"github.com/badrunp/media-sosial-web/be/services"
	"github.com/badrunp/media-sosial-web/be/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMain(t *testing.M){
	helpers.LoadEnv("../../../.env")
	
	db, _ := utils.DBConnectTest(configs.TestDatabaseConfig().Host)
	migrations.AddMigrations(db, "")
	userRepository := repositories.UserRepository(db)
	userRepository.DeleteAll()
	t.Run()
	userRepository.DeleteAll()
}

func TestCreateUser(t *testing.T) {
	var (
		db *gorm.DB
		err error
		user entity.User
		data requests.User
	)

	db, err = utils.DBConnectTest(configs.DatabaseConfig().Host)
	if assert.NoError(t, err) {

		data = requests.User{
			Username: "badrun_service",
			Email: "bbadrunn@gmail.com",
			Password: "123456",
		}

		userRepository := repositories.UserRepository(db)
		services := services.UserService(userRepository)
		user, err = services.CreateUser(data)
		assert.NoError(t, err)
		assert.Equal(t, data.Email, user.Email)

		userRepository.DeleteAll()
	}	
	
}