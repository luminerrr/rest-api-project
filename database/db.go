package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rest-api-project/models"
	"github.com/joho/godotenv"
	_"github.com/lib/pq"
)



func GetEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("cannot load the .env data")
	}

	return os.Getenv(key)
}

var (
	db *gorm.DB
	err error
)

func StartDB() {
	var (
		host = GetEnv("HOST")
		username = GetEnv("DB_USERNAME")
		password = GetEnv("DB_PASSWORD")
		port = GetEnv("PORT")
		dbname = GetEnv("DB_NAME")
		
	)
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbname, port)

	db,err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database", err)
	}

	db.Debug().AutoMigrate(models.Items{}, models.Orders{})
}

func GetDB() *gorm.DB{
	return db
}