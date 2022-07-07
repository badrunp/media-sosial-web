package test

import (
	"testing"

	"github.com/badrunp/media-sosial-web/be/configs"
	migrations "github.com/badrunp/media-sosial-web/be/database/migrations"
	"github.com/badrunp/media-sosial-web/be/entity"
	"github.com/badrunp/media-sosial-web/be/helpers"
	"github.com/badrunp/media-sosial-web/be/repositories"
	repository "github.com/badrunp/media-sosial-web/be/repositories"
	"github.com/badrunp/media-sosial-web/be/utils"
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

func TestGetAllUsers(t *testing.T){
	db, err := utils.DBConnectTest(configs.TestDatabaseConfig().Host)
	
	if assert.NoError(t, err) {
		userRepository := repository.UserRepository(db)
		_, err := userRepository.GetAll()
		assert.NoError(t, err)
	}
}

func TestGetUserByEmail(t *testing.T) {
	db, err := utils.DBConnectTest(configs.DatabaseConfig().Host)
	
	if assert.NoError(t, err) {
		data := entity.User{
			Username: "test",
			Email: "bbadrunn@gmail.com",
			Password: "123456",
		}
		var email string = "bbadrunn@gmail.com"
		var err error
		var user entity.User

		userRepository := repository.UserRepository(db)

		user, err = userRepository.Create(data)
		assert.NoError(t, err)
		assert.Equal(t, data.Email, user.Email)

		user, err = userRepository.FindByEmail(email)
		assert.NoError(t, err)
		assert.Equal(t, email, user.Email)

		userRepository.DeleteAll()
	}
}

func TestGetUserByEmailFail(t *testing.T){
	db, err := utils.DBConnectTest(configs.DatabaseConfig().Host)
	
	if assert.NoError(t, err) {
		data := entity.User{
			Username: "test",
			Email: "bbadrunn@gmail.com",
			Password: "123456",
		}
		var email string = "notfoud@gmail.com"
		var err error
		var user entity.User

		userRepository := repository.UserRepository(db)

		user, err = userRepository.Create(data)
		assert.NoError(t, err)
		assert.Equal(t, data.Email, user.Email)

		_, err = userRepository.FindByEmail(email)
		assert.Error(t, err)
		assert.EqualError(t, err, "record not found")

		userRepository.DeleteAll()
	}
}

func TestGetUserById(t *testing.T) {
	db, err := utils.DBConnectTest(configs.DatabaseConfig().Host)
	
	if assert.NoError(t, err) {
		data := entity.User{
			Username: "test",
			Email: "test@gmail.com",
			Password: "123456",
		}

		userRepository := repository.UserRepository(db)

		newUser, err := userRepository.Create(data)
		assert.NoError(t, err)
		assert.Equal(t, data.Email, newUser.Email)

		user, err := userRepository.FindByID(newUser.ID)
		assert.NoError(t, err)
		assert.Equal(t, newUser.ID, user.ID)

		userRepository.DeleteAll()
	}
}

func TestGetUserByIdFail(t *testing.T) {
	db, err := utils.DBConnectTest(configs.DatabaseConfig().Host)
	
	if assert.NoError(t, err) {
		data := entity.User{
			Username: "test",
			Email: "test@gmail.com",
			Password: "123456",
		}
		var id int = 999
		var err error
		var user entity.User

		userRepository := repository.UserRepository(db)

		user, err = userRepository.Create(data)
		assert.NoError(t, err)
		assert.Equal(t, data.Email, user.Email)

		_, err = userRepository.FindByID(id)
		assert.Error(t, err)
		assert.EqualError(t, err, "record not found")

		userRepository.DeleteAll()
	}
}

func TestCreateUser(t *testing.T) {
	db, err := utils.DBConnectTest(configs.DatabaseConfig().Host)
	
	if assert.NoError(t, err) {
		data := entity.User{
			Username: "test",
			Email: "test@gmail.com",
			Password: "123456",
		}

		userRepository := repository.UserRepository(db)
		user, err := userRepository.Create(data)

		assert.NoError(t, err)
		assert.Equal(t, data.Email, user.Email)

		userRepository.DeleteAll()
	}
}

func TestCreateUserIsUniqueEmail(t *testing.T) {
	db, err := utils.DBConnectTest(configs.DatabaseConfig().Host)
	
	if assert.NoError(t, err) {
		var err error
		var user entity.User
		data := entity.User{
			Username: "test",
			Email: "test8@gmail.com",
			Password: "123456",
		}

		userRepository := repository.UserRepository(db)

		user, err = userRepository.Create(data)
		assert.NoError(t, err)
		assert.Equal(t, data.Email, user.Email)
		
		_, err = userRepository.Create(data)
		assert.Error(t, err)
		assert.EqualError(t, err, "ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)")

		userRepository.DeleteAll()
	}
}

func TestDeleteUsers(t *testing.T){
	db, err := utils.DBConnectTest(configs.DatabaseConfig().Host)
	if assert.NoError(t, err) {
		userRepository := repository.UserRepository(db)

		isD, err := userRepository.DeleteAll()
		assert.Equal(t, true, isD)
		assert.NoError(t, err)
	}
}

func TestDeleteUserById(t *testing.T){
	db, err := utils.DBConnectTest(configs.DatabaseConfig().Host)

	if assert.NoError(t, err) {
		var user entity.User
		var err error
		data := entity.User{
			Username: "test",
			Email: "test8@gmail.com",
			Password: "123456",
		}

		userRepository := repository.UserRepository(db)

		user, err = userRepository.Create(data)
		assert.NoError(t, err)
		assert.Equal(t, data.Email, user.Email)

		var id int = user.ID

		isDelete, err := userRepository.DeleteById(id)
		assert.NoError(t, err)
		assert.Equal(t, true, isDelete)

		_, err = userRepository.FindByID(id)
		assert.Error(t, err)
		assert.EqualError(t, err, "record not found")

		userRepository.DeleteAll()
	}

}

func TestDeleteUserByIdFail(t *testing.T){
	db, err := utils.DBConnectTest(configs.DatabaseConfig().Host)

	if assert.NoError(t, err) {
		var user entity.User
		var err error
		data := entity.User{
			Username: "test",
			Email: "test8@gmail.com",
			Password: "123456",
		}

		userRepository := repository.UserRepository(db)

		user, err = userRepository.Create(data)
		assert.NoError(t, err)
		assert.Equal(t, data.Email, user.Email)

		var id int = 1000

		isDelete, err := userRepository.DeleteById(id)
		assert.Error(t, err)
		assert.Equal(t, false, isDelete)
		assert.EqualError(t, err, "record not found")

		user2, err := userRepository.FindByID(user.ID)
		assert.NoError(t, err)
		assert.Equal(t, user.Email, user2.Email)

		userRepository.DeleteAll()
	}

}