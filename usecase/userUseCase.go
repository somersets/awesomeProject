package usecase

import (
	"awesomeProject/domain"
	"awesomeProject/usecase/repository"
)

type User interface {
	Disable(id int) (int, error)
	Create(u *domain.User) (*domain.User, error)
	GetOneById(userId int) (*domain.UserDTO, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
	dbRepository   repository.DBRepository
}

func NewUserUseCase(uR repository.UserRepository, dbR repository.DBRepository) User {
	return &userUseCase{userRepository: uR, dbRepository: dbR}
}

func (usc *userUseCase) GetOneById(userId int) (*domain.UserDTO, error) {
	user, err := usc.userRepository.GetOneById(userId)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (usc *userUseCase) Create(user *domain.User) (*domain.User, error) {
	user, err := usc.userRepository.CreateUser(user)
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

	return user, nil
}

func (usc *userUseCase) Disable(id int) (int, error) {
	id, err := usc.userRepository.UpdateDisable(id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
