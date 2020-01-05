package gorm

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/ryicoh/clean-arch/internal/infrastructure/datastore"
)

type database struct {
	db *gorm.DB
}

// New defines ...
func NewDatabase(c *datastore.DBConfig) (datastore.DB, error) {
	conn, err := gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@%s(%s)/%s?charset=utf8mb4&multiStatements=True&parseTime=True&loc=Asia%%2FTokyo",
			c.User,
			c.Password,
			c.Protocol,
			c.Host,
			c.Name,
		))

	if err != nil {
		return nil, err
	}

	return &database{db: conn}, nil
}

func (d *database) AutoMigrate(i ...interface{}) datastore.DB {
	return &database{d.db.AutoMigrate(i...)}
}

func (d *database) Find(i interface{}) datastore.DB {
	return &database{d.db.Find(i)}
}

func (d *database) Error() error {
	return d.db.Error
}
