package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ryicoh/clean-arch/internal/adapter/repository"
	"github.com/ryicoh/clean-arch/internal/adapter/web/controller"
	"github.com/ryicoh/clean-arch/internal/adapter/web/route"
	"github.com/ryicoh/clean-arch/internal/infrastructure/conf"
	"github.com/ryicoh/clean-arch/internal/infrastructure/datastore"
	"github.com/ryicoh/clean-arch/internal/infrastructure/web/echo"
	"github.com/ryicoh/clean-arch/internal/usecase"
)

func main() {
	cnf := conf.New()
	s := echo.NewServer(cnf)
	dbCnf := datastore.NewDBConfigFromConfig(cnf.GetDatabaseConfig())
	db, err := datastore.New(dbCnf)
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
