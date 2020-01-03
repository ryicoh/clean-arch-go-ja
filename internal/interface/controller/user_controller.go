package controller

import (
	"net/http"

	"github.com/ryicoh/clean-arch/internal/infrastructure/web"
	"github.com/ryicoh/clean-arch/internal/usecase"
)

type (
	UserController interface {
		GetUsers(web.Context) (code int, i interface{}, err error)
	}

	userController struct {
		userUsecase usecase.UserUsecase
	}
)

func NewUserController(userUsecase usecase.UserUsecase) UserController {
	return &userController{userUsecase: userUsecase}
}

func (c *userController) GetUsers(ctx web.Context) (code int, i interface{}, err error) {
	offset, err := ctx.GetQueryOffset()
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	users, err := c.userUsecase.GetUsers(offset)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, users, nil
}
