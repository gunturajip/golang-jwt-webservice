package database

import (
	"fmt"
	"golang-jwt-auth/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}
	// host     := "localhost"
	host := os.Getenv("PGHOST")
	// user     := "postgres"
	user := os.Getenv("PGUSER")
	// password := "postgres"
	password := os.Getenv("PGPASSWORD")
	// port     := "5432"
	port := os.Getenv("PGPORT")
	// dbname   := "db-go-sql"
	dbname := os.Getenv("PGDATABASE")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	fmt.Println("successfully connected to database🔥")
	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
