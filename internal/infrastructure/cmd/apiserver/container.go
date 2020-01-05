package apiserver

import (
	"github.com/ryicoh/clean-arch/internal/infrastructure/conf"
	"github.com/ryicoh/clean-arch/internal/infrastructure/conf/yaml"
	"github.com/ryicoh/clean-arch/internal/infrastructure/datastore"
	"github.com/ryicoh/clean-arch/internal/infrastructure/datastore/gorm"
	"github.com/ryicoh/clean-arch/internal/infrastructure/web/echo"
	"github.com/ryicoh/clean-arch/internal/interface/controller"
	"github.com/ryicoh/clean-arch/internal/interface/repository"
	"github.com/ryicoh/clean-arch/internal/interface/web/route"
	"github.com/ryicoh/clean-arch/internal/usecase"
	"go.uber.org/dig"
)

type (
	Container interface {
		Build() error
		Run() error
	}
	container struct {
		*dig.Container
	}
)

func NewContainer() Container {
	return &container{dig.New()}
}

func (c *container) Build() error {
	if err := c.Provide(yaml.New); err != nil {
		return err
	}

	if err := c.Provide(datastore.NewDBConfigFromENV); err != nil {
		return err
	}

	if err := c.Provide(gorm.NewDatabase); err != nil {
		return err
	}

	if err := c.Provide(repository.NewUserRepository); err != nil {
		return err
	}

	if err := c.Provide(usecase.NewUserUsecase); err != nil {
		return err
	}

	if err := c.Provide(controller.NewUserController); err != nil {
		return err
	}

	if err := c.Provide(controller.NewAppController); err != nil {
		return err
	}

	if err := c.Provide(echo.NewServer); err != nil {
		return err
	}

	if err := c.Provide(route.NewRouter); err != nil {
		return err
	}

	return nil
}

func (c *container) Run() error {
	err := c.Invoke(func(router route.Router, cnf conf.Config) error {
		return router.Start(cnf.GetPort())
	})

	if err != nil {
		return err
	}
	return nil
}
