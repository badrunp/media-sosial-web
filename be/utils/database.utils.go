package utils

import (
	"fmt"
	"log"

	"github.com/badrunp/media-sosial-web/be/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnect(host string)(*gorm.DB, error){
	var err error
	var db *gorm.DB
	user := configs.DatabaseConfig().User
	password := configs.DatabaseConfig().Password
	dbname := configs.DatabaseConfig().Name
	port := configs.DatabaseConfig().Port 
	sslMode := configs.DatabaseConfig().SslMode

	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Shanghai", host, user, password, dbname, port, sslMode)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt: true,
	})
	if err != nil && host == configs.DatabaseConfig().Host {
		log.Fatal("error ", err.Error())
	}

	fmt.Println("Database connected on port", port)
	fmt.Println("ssl mode =", sslMode)
	return db, err
}

func DBConnectTest(host string)(*gorm.DB, error){
	var err error
	var db *gorm.DB
	user := configs.TestDatabaseConfig().User
	password := configs.TestDatabaseConfig().Password
	dbname := configs.TestDatabaseConfig().Name
	port := configs.TestDatabaseConfig().Port 
	sslMode := configs.TestDatabaseConfig().SslMode

	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Shanghai", host, user, password, dbname, port, sslMode)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil && host == configs.TestDatabaseConfig().Host {
		log.Fatal("error ", err.Error())
	}

	fmt.Println("Database connected on port", port)
	fmt.Println("ssl mode =", sslMode)
	return db, err
}



