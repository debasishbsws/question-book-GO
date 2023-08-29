package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT          string
	DB_USER_NAME  string
	DB_PASSWORD   string
	DATABASE_NAME string
	DB_HOST       string
)

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file")
	}

	PORT = os.Getenv("PORT")
	DB_USER_NAME = os.Getenv("DB_USER_NAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DATABASE_NAME = os.Getenv("DATABASE_NAME")
	DB_HOST = os.Getenv("DB_HOST")
	if PORT == "" {
		PORT = "8000"
	}

	fmt.Println("PORT:", PORT)
	fmt.Println("DB_USER_NAME:", DB_USER_NAME)

}
