package database

import (
	"github.com/badrunp/media-sosial-web/be/models"
	"gorm.io/gorm"
)

type userMigration struct {
	db *gorm.DB
	user *models.User 
}

func NewUserMigration(db *gorm.DB, user *models.User)*userMigration{
	return &userMigration{db: db, user: user}
}

func (u *userMigration) Up(){
	u.db.AutoMigrate(u.user)
}