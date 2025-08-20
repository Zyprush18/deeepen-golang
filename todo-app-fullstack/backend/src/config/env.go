package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type takeEnv struct {
	DBHost	string
	DBPort 	string
	DBName string
	DBUser string
	DBPass string
	JwtKey string
}

func newEnv() *takeEnv  {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalln(err.Error())
	}

	return &takeEnv{
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
		DBUser: os.Getenv("DB_USERNAME"),
		DBPass: os.Getenv("DB_PASSWORD"),
		JwtKey: os.Getenv("JWT_SECRET_KEY"),
	}
}