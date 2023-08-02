package infrastructure

import (
	"awesomeProject/domain"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func (api *API) newPostgre() (*gorm.DB, error) {
	dbApiConfig := api.config.DbConfig
	dbConfig := postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			dbApiConfig.Host,
			dbApiConfig.DbUser,
			dbApiConfig.DbPassword,
			dbApiConfig.DbName,
			dbApiConfig.Port,
		),
	}
	db, err := gorm.Open(postgres.New(dbConfig), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Error occured while connecting to database: %s", err.Error())
		return nil, err
	}
	err = db.AutoMigrate(&domain.User{}, &domain.Token{})
	if err != nil {
		logrus.Printf("Failed migrating!")
		return nil, err
	}
	return db, err
}
