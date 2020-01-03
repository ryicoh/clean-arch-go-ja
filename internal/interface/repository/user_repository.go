package repository

import (
	"github.com/ryicoh/clean-arch/internal/domain/model"
	"github.com/ryicoh/clean-arch/internal/infrastructure/datastore"
)

type (
	UserOptions struct {
		Limit  int
		Offset int
	}

	UserRepository interface {
		Fetch(opts *UserOptions) ([]*model.User, error)
	}

	userRepository struct {
		db datastore.DB
	}
)

func NewUserRepository(db datastore.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Fetch(opts *UserOptions) ([]*model.User, error) {
	var users []*model.User
	err := r.db.Find(&users).Error()
	if err != nil {
		return nil, err
	}

	return users, nil
}
