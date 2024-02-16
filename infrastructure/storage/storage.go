package storage

import (
	"awesomeProject/domain"
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type Storage struct {
	config *postgres.Config
}

func New(config *postgres.Config) *Storage {
	return &Storage{config: config}
}

func (storage *Storage) Open() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(*storage.config), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Error occured while connecting to database: %s", err.Error())
		return nil, err
	}

	return db, nil
}

func (storage *Storage) AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&domain.User{},
		&domain.Token{},
		&domain.UserSexualOrientationType{},
		&domain.UserPhoto{},
		&domain.UserActivationLink{},
		&domain.Chat{},
		&domain.ChatMember{},
		&domain.ChatMessageHistory{},
	)

	if err != nil {
		return err
	}
	return nil
}

func (storage *Storage) AutoSeeders(db *gorm.DB) error {
	if err := db.AutoMigrate(&domain.UserSexualOrientationType{}); err == nil && db.Migrator().HasTable(&domain.UserSexualOrientationType{}) {
		if err := db.First(&domain.UserSexualOrientationType{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			db.Model(&domain.UserSexualOrientationType{}).Create(&domain.UserSexualOrientationType{Orientation: "Heterosexual"})
			db.Model(&domain.UserSexualOrientationType{}).Create(&domain.UserSexualOrientationType{Orientation: "Homosexual"})
			db.Model(&domain.UserSexualOrientationType{}).Create(&domain.UserSexualOrientationType{Orientation: "Bisexual"})
			db.Model(&domain.UserSexualOrientationType{}).Create(&domain.UserSexualOrientationType{Orientation: "Asexual"})
		}
	}
	return nil
}

func (storage *Storage) Close() error {
	return nil
}
