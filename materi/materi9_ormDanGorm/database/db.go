package database

// menginstall driver postgres dengan go get gorm.io/driver/postgres

import (
	"fmt"
	"log"
	"materi9_ormDanGorm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "admin"
	dbPort   = "5432"
	dbname   = "learning-gorm"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	// db akan mengadung referensi dari database dengan tipe data *gorm.DB
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}
	// debugging atau loggier yang kemudian di chaining dengan method AutoMigrate yang digunakan untuk migrassi secara otomatis
	// dari struct yang telah dibuat
	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

// CRUD dengan GORM
func GetDB() *gorm.DB {
	return db
}
