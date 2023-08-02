package repository

import (
	"awesomeProject/domain"
	"awesomeProject/usecase/repository"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type useRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &useRepository{db: db}
}

func (uR *useRepository) CreateUser(user *domain.User) (*domain.User, error) {
	var alreadyExistUser *domain.User
	if err := uR.db.Model(&domain.User{}).Where("email = ?", user.Email).First(&alreadyExistUser).Error; err == nil {
		if alreadyExistUser.Email == user.Email {
			return nil, errors.New(fmt.Sprintf("A userRegistry with email %s already existing", user.Email))
		}
	}

	if err := uR.db.Model(&domain.User{}).Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (uR *useRepository) GetOneById(userId int) (*domain.UserDTO, error) {
	var user *domain.User
	if err := uR.db.Model(&domain.User{}).Where("id = ?", userId).First(&user).Error; err != nil {

	}
	return domain.NewUserDTO(user), nil
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
