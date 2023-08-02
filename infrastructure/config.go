package infrastructure

import (
	"os"
)

type Config struct {
	ServerPort  string
	LoggerLevel string
	DbConfig    *DbConfig
}

type DbConfig struct {
	Host       string
	Port       string
	DbName     string
	DbUser     string
	DbPassword string
}

func NewConfig() *Config {
	LoadEnv()

	return &Config{
		ServerPort:  os.Getenv("SERVER_PORT"),
		LoggerLevel: "debug",
		DbConfig: &DbConfig{
			Host:       os.Getenv("HOST"),
			Port:       os.Getenv("PORT"),
			DbName:     os.Getenv("DB_NAME"),
			DbUser:     os.Getenv("DB_USER"),
			DbPassword: os.Getenv("DB_PASSWORD"),
		}}
}
