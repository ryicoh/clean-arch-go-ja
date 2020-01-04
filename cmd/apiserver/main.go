package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ryicoh/clean-arch/internal/infrastructure/conf/env"
	"github.com/ryicoh/clean-arch/internal/infrastructure/datastore"
	"github.com/ryicoh/clean-arch/internal/infrastructure/datastore/gorm"
	"github.com/ryicoh/clean-arch/internal/infrastructure/web/echo"
	"github.com/ryicoh/clean-arch/internal/interface/controller"
	"github.com/ryicoh/clean-arch/internal/interface/repository"
	"github.com/ryicoh/clean-arch/internal/interface/web/route"
	"github.com/ryicoh/clean-arch/internal/usecase"
)

func main() {
	cnf := env.New()
	s := echo.NewServer(cnf)
	db, err := gorm.New(datastore.NewDBConfigFromENV())
	if err != nil {
		fmt.Println(err)
		return
	}

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	appController := controller.NewAppController(userController)

	route.Register(s, appController)

	err = s.Start(cnf.GetPort())
	if err != nil {
		fmt.Println(err)
	}
}
