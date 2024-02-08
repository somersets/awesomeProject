package usecase

import (
	"awesomeProject/domain"
	"awesomeProject/usecase/repository"
)

type User interface {
	Disable(id int) (int, error)
	Create(u *domain.User) (*domain.UserDTO, error)
	GetOneById(userId int) (*domain.UserDTO, error)
	Update(userFormUpdate *domain.UserProfileUpdateForm, userId int) (*domain.UserDTO, error)
	GetSexOrientations() ([]*domain.UserSexualOrientationType, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
	dbRepository   repository.DBRepository
}

func NewUserUseCase(uR repository.UserRepository, dbR repository.DBRepository) User {
	return &userUseCase{userRepository: uR, dbRepository: dbR}
}

func (usc *userUseCase) GetSexOrientations() ([]*domain.UserSexualOrientationType, error) {
	orientations, err := usc.userRepository.GetSexOrientations()
	if err != nil {
		return nil, err
	}

	return orientations, nil
}

func (usc *userUseCase) Update(userFormUpdate *domain.UserProfileUpdateForm, userId int) (*domain.UserDTO, error) {
	user, err := usc.userRepository.UpdateUser(userFormUpdate, userId)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (usc *userUseCase) GetOneById(userId int) (*domain.UserDTO, error) {
	user, err := usc.userRepository.GetOneById(userId)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (usc *userUseCase) Create(user *domain.User) (*domain.UserDTO, error) {
	createdUser, err := usc.userRepository.CreateUser(user)
	/*data, err := usc.dbRepository.Transaction(func(i interface{}) (interface{}, error) {

		// do mailing
		// do logging
		// do another process
		return u, err
	})*/
	//userRegistry, ok := data.(*userRegistry.User)

	//if !ok {
	//	return nil, errors.New("cast error")
	//}

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (usc *userUseCase) Disable(id int) (int, error) {
	id, err := usc.userRepository.UpdateDisable(id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
