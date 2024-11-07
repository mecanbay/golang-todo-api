package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	FIBER_APPNAME       string
	FIBER_HOST          string
	FIBER_PORT          string
	FIBER_PREFORK       string
	FIBER_SERVER_HEADER string
	DB_HOST             string
	DB_USER             string
	DB_PASSWORD         string
	DB_NAME             string
	DB_PORT             string
	DB_SSL_MODE         string
	DB_TIMEZONE         string
}

func LoadEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Err loading .env file...")
	}

	return &Env{
		FIBER_APPNAME:       os.Getenv("FIBER_APPNAME"),
		FIBER_HOST:          os.Getenv("FIBER_HOST"),
		FIBER_PORT:          os.Getenv("FIBER_PORT"),
		FIBER_PREFORK:       os.Getenv("FIBER_PREFORK"),
		FIBER_SERVER_HEADER: os.Getenv("FIBER_SERVER_HEADER"),
		DB_HOST:             os.Getenv("DB_HOST"),
		DB_USER:             os.Getenv("DB_USER"),
		DB_PASSWORD:         os.Getenv("DB_PASSWORD"),
		DB_NAME:             os.Getenv("DB_NAME"),
		DB_PORT:             os.Getenv("DB_PORT"),
		DB_SSL_MODE:         os.Getenv("DB_SSL_MODE"),
		DB_TIMEZONE:         os.Getenv("DB_TIMEZONE"),
	}
}
