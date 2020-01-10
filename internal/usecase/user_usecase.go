package usecase

import (
	"github.com/ryicoh/clean-arch/internal/domain/model"
	"github.com/ryicoh/clean-arch/internal/usecase/repository"
)

type (
	UserUsecase struct {
		userRepository repository.UserRepository
	}
)

func NewUserUsecase(userRepository repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepository: userRepository}
}

func (u *UserUsecase) GetUsers(offset int) ([]*model.User, error) {
	opts := repository.NewUserOptions()
	users, err := u.userRepository.Fetch(opts)
	if err != nil {
		return nil, err
	}

	return users, nil
}
