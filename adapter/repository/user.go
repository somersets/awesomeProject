package repository

import (
	"awesomeProject/domain"
	"awesomeProject/usecase/repository"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type useRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &useRepository{db: db}
}

func (uR *useRepository) UpdateUser(userFormUpdate *domain.UserProfileUpdateForm, userId int) (*domain.UserDTO, error) {
	var user *domain.User
	if err := uR.db.Model(&domain.User{}).Where("id = ?", userId).Updates(userFormUpdate).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	if err := uR.db.Model(&domain.User{}).Joins("SexOrientation").Preload("Photos").First(&user, userId).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	return domain.NewUserDTO(user), nil
}

func (uR *useRepository) CreateUser(user *domain.User) (*domain.UserDTO, error) {
	var alreadyExistUser *domain.User
	if err := uR.db.Model(&domain.User{}).Where("email = ?", user.Email).First(&alreadyExistUser).Error; err == nil {
		if alreadyExistUser.Email == user.Email {
			return nil, errors.New(fmt.Sprintf("A user with email %s already exists", user.Email))
		}
	}

	if err := uR.db.Model(&domain.User{}).Create(&user).Error; err != nil {
		return nil, err
	}

	return domain.NewUserDTO(user), nil
}

func (uR *useRepository) GetOneById(userId int) (*domain.UserDTO, error) {
	var user *domain.User
	if err := uR.db.First(&user, userId).Preload("Photos", func(db *gorm.DB) *gorm.DB {
		return db.Order(clause.OrderByColumn{Column: clause.Column{Name: "order"}, Desc: false})
	}).Joins("SexOrientation").Find(&user, userId).Error; err != nil {
		return nil, err
	}

	return domain.NewUserDTO(user), nil
}

func (uR *useRepository) GetSexOrientations() ([]*domain.UserSexualOrientationType, error) {
	var orientations []*domain.UserSexualOrientationType
	result := uR.db.Find(&orientations)

	if result.Error != nil {
		return nil, result.Error
	}

	return orientations, nil

}

func (uR *useRepository) UpdateDisable(id int) (int, error) {
	var user *domain.User
	if err := uR.db.Model(&domain.User{}).First(&user, id).Error; err != nil {
		return 0, err
	}
	if err := uR.db.Model(&domain.User{}).Where("id = ?", id).Update("account_disabled", true).Error; err != nil {
		return 0, err
	}
	return id, nil
}
