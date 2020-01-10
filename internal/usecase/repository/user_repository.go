package repository

import "github.com/ryicoh/clean-arch/internal/domain/model"

type (
	UserOptions interface {
		GetLimit() int
		SetLimit(limit int) *userOptions
		GetOffset() int
		SetOffset(offset int) *userOptions
	}

	userOptions struct {
		limit  int
		offset int
	}

	UserRepository interface {
		Fetch(opts UserOptions) ([]*model.User, error)
	}
)

func NewUserOptions() UserOptions {
	return &userOptions{limit: 20, offset: 0}
}

func (o *userOptions) GetLimit() int {
	return o.limit
}

func (o *userOptions) SetLimit(limit int) *userOptions {
	o.limit = limit
	return o
}

func (o *userOptions) GetOffset() int {
	return o.offset
}

func (o *userOptions) SetOffset(offset int) *userOptions {
	o.offset = offset
	return o
}
