package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDatabase() (err error) {
	_, ok := os.LookupEnv("ENVIROMENT")

	var loadError error
	if ok {
		loadError = godotenv.Load("../.env")
	} else {
		loadError = godotenv.Load(".env")
	}

	if loadError != nil {
		fmt.Println("could not found .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s TimeZone=%s",
		os.Getenv("HOST_NAME"),
		os.Getenv("USER_NAME"),
		os.Getenv("DB_NAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("TIME_ZONE"))

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to database")
	}

	exist := database.Migrator().HasTable(&User{})
	if !exist {
		database.Migrator().CreateTable(&User{})
	}
	Db = database

	return
}
