package config

import (
	"os"

	"github.com/joho/godotenv"
)

var cfg *appConfig

type appConfig struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func Load() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	cfg = &appConfig{
		API: APIConfig{
			Port: os.Getenv("PORT"),
		},
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_NAME"),
		},
	}

	return nil
}

func Get() *appConfig {
	return cfg
}
