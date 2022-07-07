package database

import (
	"github.com/badrunp/media-sosial-web/be/models"
	"github.com/badrunp/media-sosial-web/be/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userSeed struct {
	db *gorm.DB
}

func UserSeed(db *gorm.DB)*userSeed{
	return &userSeed{db: db}
}

func (u *userSeed) Run(){
	user := models.User{Username: "admin", Email: "admin@gmail.com", Password: utils.HashAndSalt([]byte("123456"))}
	u.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)
}