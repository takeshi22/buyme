package models

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var Client *redis.Client

func InitDatabase() (err error) {
	_, ok := os.LookupEnv("ENVIROMENT")

	var loadError error
	if ok {
		loadError = godotenv.Load("../.test.env")
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

	Db = database

	return
}

func InitRedis() (err error) {
	dsn := "redis:6379"
	Client = redis.NewClient((&redis.Options{
		Addr: dsn,
	}))
	_, err = Client.Ping().Result()
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}
