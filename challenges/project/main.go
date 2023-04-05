package main

import (
	"fmt"
	"log"
	"project/entity"
	"project/routers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "admin"
	dbPort   = "5432"
	dbName   = "project-book"
	db       *gorm.DB
	err      error
)

func init() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	db.Debug().AutoMigrate(entity.Book{})
}

func main() {
	routers.StartServer(db).Run(":4000")
}
