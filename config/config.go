package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT         string
	DATABASE_URI string
)

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file")
	}

	PORT = os.Getenv("PORT")
	DATABASE_URI = os.Getenv("DATABASE_URI")
	if PORT == "" {
		PORT = "8000"
	}
}
