package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ryicoh/clean-arch/internal/domain/model"
	"github.com/ryicoh/clean-arch/internal/infrastructure/datastore"
	"github.com/ryicoh/clean-arch/internal/infrastructure/datastore/gorm"
)

func main() {
	db, err := gorm.NewDatabase(datastore.NewDBConfigFromENV())
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.AutoMigrate(&model.User{}).Error()
	if err != nil {
		fmt.Println(err)
	}
}
