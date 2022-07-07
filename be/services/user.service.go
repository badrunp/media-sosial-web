package services

import (
	"github.com/badrunp/media-sosial-web/be/entity"
	"github.com/badrunp/media-sosial-web/be/repositories"
	"github.com/badrunp/media-sosial-web/be/requests"
)

type User interface {
	CreateUser(user requests.User)(entity.User, error)
	GetAllUser()([]entity.User, error)
	GetUserByEmail(data requests.UserLogin)(entity.User, error)
}

type user struct {
	repository repositories.User
}

func UserService(repository repositories.User)*user{
	return &user{repository: repository}
}

func (u *user) CreateUser(data requests.User)(entity.User, error){
	var user entity.User
	var err error
	user = entity.User{
		Username: data.Username,
		Email: data.Email,
		Password: data.Password,
	}

	user, err = u.repository.Create(user)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u *user) GetAllUser()([]entity.User, error) {
	users, err := u.repository.GetAll()
	if err != nil {
		return []entity.User{}, nil
	}
	return users, nil
}

func (u *user) GetUserByEmail(data requests.UserLogin)(entity.User, error){
	user, err := u.repository.FindByEmail(data.Email)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}