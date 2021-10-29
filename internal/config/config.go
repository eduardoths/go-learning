package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	JWT_SECRET  string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading environment variables, Err= %s", err.Error())
		panic("Couldn't load environment variables")
	}
}

func GetConfig() Config {
	return Config{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_PORT:     os.Getenv("DB_PORT"),
		JWT_SECRET:  os.Getenv("JWT_SECRET"),
	}
}

func (c Config) GetDBConnector() *gorm.DB {
	dialector := postgres.Open(fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.DB_HOST,
		c.DB_USER,
		c.DB_PASSWORD,
		c.DB_NAME,
		c.DB_PORT,
	))

	db, err := gorm.Open(dialector)
	if err != nil {
		log.Fatalf("Error connecting to database, Err= %s", err.Error())
	}
	return db
}
