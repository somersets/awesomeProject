package infrastructure

import (
	"awesomeProject/infrastructure/storage"
	"gorm.io/driver/postgres"
	"os"
)

type Config struct {
	ServerPort    string
	LoggerLevel   string
	StorageConfig *postgres.Config
}

func NewConfig() *Config {
	return &Config{
		ServerPort:    os.Getenv("SERVER_PORT"),
		LoggerLevel:   "debug",
		StorageConfig: storage.NewConfig(),
	}
}
