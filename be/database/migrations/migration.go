package database

import (
	"fmt"
	"os"

	"github.com/badrunp/media-sosial-web/be/models"
	"gorm.io/gorm"
)

type Migrations interface {
	Add(mg Migration)
}

type Migration interface {
	Up()
}

type migration struct {
	Migrations []Migration
}

func NewMigration() *migration {
	return &migration{}
}

func (m *migration) Add(mg Migration) {
	m.Migrations = append(m.Migrations, mg)
}

func RegisterMigration(db *gorm.DB, args []string) string {
	
	if len(args[1:]) > 0 {
		isMigrate := args[1:][0] == "migrate:up"
		if isMigrate {
			AddMigrations(db, args[0])

			fmt.Println("migration success")
			if args[0] != "testing" {
				os.Exit(1)
			} 
			return "migration success"
		}
	}else{
		if mode := os.Getenv("NODE_ENV"); mode == "production" {
			AddMigrations(db, "")
			return "migration success"
		}
	}

	return "migration failed"
}

func AddMigrations(db *gorm.DB, mode string){
	migrate := NewMigration()
	userMigrate := NewUserMigration(db, &models.User{})

	migrate.Add(userMigrate)
	
	if mode != "testing" {
		for _, mg := range migrate.Migrations {
			mg.Up()
		} 
	} 
}