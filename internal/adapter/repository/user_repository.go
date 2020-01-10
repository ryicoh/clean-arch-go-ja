package repository

import (
	"github.com/ryicoh/clean-arch/internal/adapter/interface/datastore"
	"github.com/ryicoh/clean-arch/internal/domain/model"
	"github.com/ryicoh/clean-arch/internal/usecase/repository"
)

type userRepository struct {
	db datastore.DB
}

func NewUserRepository(db datastore.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Fetch(opts repository.UserOptions) ([]*model.User, error) {
	var users []*model.User
	err := r.db.Limit(opts.GetLimit()).Offset(opts.GetOffset()).Find(&users).Error()
	if err != nil {
		return nil, err
	}

	return users, nil
}
