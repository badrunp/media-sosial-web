package repositories

import (
	"fmt"

	"github.com/badrunp/media-sosial-web/be/entity"
	"github.com/badrunp/media-sosial-web/be/models"
	"gorm.io/gorm"
)

type User interface {
	Create(data entity.User)(entity.User, error)
	FindByEmail(email string) (entity.User, error)
	DeleteAll()(bool, error)
	FindByID(id int) (entity.User, error)
	DeleteById(id int)(bool, error)
	GetAll()([]entity.User, error)
}

type user struct {
	db *gorm.DB
}

func UserRepository(db *gorm.DB)*user {
	return &user{db: db}
}

func (u *user) GetAll()([]entity.User, error){
	var (
		err error
		users []entity.User
	)

	err = u.db.Select("id", "username", "email", "password").Find(&users).Error
	if err != nil {
		return []entity.User{}, err
	}

	fmt.Println(users)
	
	return users, nil
}

func (u *user) Create(data entity.User)(entity.User, error){
	user := models.User{
		Username: data.Username,
		Email: data.Email,
		Password: data.Password,
	}

	err := u.db.Create(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	
	result, err := u.FindByID(int(user.ID))
	if err != nil {
		return entity.User{}, err
	}

	return result, nil
}

func (u *user) DeleteById(id int)(bool, error){
	_, err := u.FindByID(id)
	if err != nil {
		return false, err
	}

	u.db.Unscoped().Delete(&models.User{}, id)
	
	return true, nil
}

func (u *user) DeleteAll()(bool, error){
	err := u.db.Exec("TRUNCATE TABLE users RESTART IDENTITY").Error
	if err != nil {
		return false, err
	}
	
	return true, nil
}

func (u *user) FindByEmail(email string) (entity.User, error){
	var user models.User
	err := u.db.First(&user, "email = ?", email).Error
	if err != nil {
		return entity.User{}, err
	}
	
	return entity.User{
		ID: int(user.ID),
		Username: user.Username,
		Email: user.Email,
		Password: user.Password,
	}, nil
}

func (u *user) FindByID(id int) (entity.User, error){
	var user models.User
	err := u.db.First(&user, "id = ?", id).Error
	if err != nil {
		return entity.User{}, err
	}
	
	return entity.User{
		ID: int(user.ID),
		Username: user.Username,
		Email: user.Email,
		Password: user.Password,
	}, nil
}