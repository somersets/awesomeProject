package auth

import (
	"awesomeProject/domain"
	"awesomeProject/infrastructure/utils"
	"awesomeProject/usecase/repository/auth"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) auth.AuthRepository {
	return &authRepository{db: db}
}

func (aR *authRepository) LoginUser(loginForm *domain.LoginFormDTO) (*domain.UserDTO, error) {
	var alreadyExistUser *domain.User

	var (
		loginRes *domain.UserDTO
		err      error
	)
	existUserErr := aR.db.Model(&domain.User{}).Where("email = ?", loginForm.Email).Joins("SexOrientation").Preload("Photos").First(&alreadyExistUser).Error
	if existUserErr != nil || !utils.CheckPasswordHash(loginForm.Password, *alreadyExistUser.Password) {
		loginRes, err = nil, errors.New(fmt.Sprintf("Email or password is incorrect"))
		return loginRes, err
	}

	loginRes, err = domain.NewUserDTO(alreadyExistUser), nil
	return loginRes, err
}

func (aR *authRepository) CreateUser(user *domain.User, activationLink string) (*domain.UserDTO, error) {
	var dbUser *domain.User
	if err := aR.db.Model(&domain.User{}).Where("email = ? OR phone = ?", user.Email, user.Phone).First(&dbUser).Error; err == nil {
		if dbUser.Email == user.Email {
			return nil, errors.New(fmt.Sprintf("A user with email %s already existing", user.Email))
		}
		if dbUser.Phone == user.Phone {
			return nil, errors.New(fmt.Sprintf("A user with phone %d already existing", user.Phone))
		}
	}

	newUser := &domain.User{
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Patronymic: user.Patronymic,
		Password:   user.Password,
		Email:      user.Email,
		Phone:      user.Phone,
		Birthday:   user.Birthday,
		Gender:     user.Gender,
		Activated:  false,
		DeletedAt:  gorm.DeletedAt{},
		Photos:     nil,
		ActivationLink: domain.UserActivationLink{
			Link:     activationLink,
			ExpireAt: time.Now().Add(10 * time.Minute),
		},
		SexOrientationID: nil,
	}

	if err := aR.db.Model(&domain.User{}).Create(&newUser).Error; err != nil {
		return nil, err
	}

	return domain.NewUserDTO(newUser), nil
}

func (aR *authRepository) ActivateUser(activationLink string) (*domain.UserDTO, error) {
	var activeLink *domain.UserActivationLink

	if err := aR.db.Model(&domain.UserActivationLink{}).Where("link = ? AND expire_at > ?", activationLink, time.Now()).First(&activeLink).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("The link was expired"))
	}

	var user *domain.User
	if err := aR.db.Model(&domain.User{}).Where("id = ?", activeLink.UserId).First(&user).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("Unexpected error during account activation"))
	}

	if user.Activated {
		return nil, errors.New(fmt.Sprintf("A user was already activated"))
	}

	aR.db.Model(&user).Update("activated", true)

	return domain.NewUserDTO(user), nil
}
