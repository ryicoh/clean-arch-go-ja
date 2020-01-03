package usecase

import (
	"github.com/ryicoh/clean-arch/internal/domain/model"
	"github.com/ryicoh/clean-arch/internal/interface/repository"
)

type (
	UserUsecase interface {
		GetUsers(offset int) ([]*model.User, error)
	}

	userUsecase struct {
		userRepository repository.UserRepository
	}
)

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{userRepository: userRepository}
}

func (u *userUsecase) GetUsers(offset int) ([]*model.User, error) {
	opts := &repository.UserOptions{Limit: 20, Offset: offset}
	users, err := u.userRepository.Fetch(opts)
	if err != nil {
		return nil, err
	}

	return users, nil
}
