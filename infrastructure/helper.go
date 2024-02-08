package infrastructure

import (
	"awesomeProject/infrastructure/storage"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
)

func (api *API) configureLoggerField() error {
	logLevel, err := logrus.ParseLevel(api.config.LoggerLevel)
	if err != nil {
		return err
	}
	api.logger.SetLevel(logLevel)
	return nil
}

func (api *API) configureStorageField() (*gorm.DB, error) {
	logrus.Println(api.config.StorageConfig)
	store := storage.New(api.config.StorageConfig)
	dbStorage, err := store.Open()

	migrateErr := store.AutoMigrate(dbStorage)
	if migrateErr != nil {
		log.Fatalf("Migrate error: %s", err.Error())
		return nil, migrateErr
	}

	seedersErr := store.AutoSeeders(dbStorage)
	if seedersErr != nil {
		log.Fatalf("Seeders error: %s", err.Error())
		return nil, seedersErr
	}

	if err != nil {
		return nil, err
	}

	return dbStorage, nil
}
