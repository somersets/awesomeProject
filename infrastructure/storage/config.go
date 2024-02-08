package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"os"
)

type Config struct {
	Host       string
	Port       string
	DbName     string
	DbUser     string
	DbPassword string
}

func NewConfig() *postgres.Config {
	dbApiConfig := &Config{
		Host:       os.Getenv("HOST"),
		Port:       os.Getenv("PORT"),
		DbName:     os.Getenv("DB_NAME"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
	}
	return &postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			dbApiConfig.Host,
			dbApiConfig.DbUser,
			dbApiConfig.DbPassword,
			dbApiConfig.DbName,
			dbApiConfig.Port,
		),
	}

}
