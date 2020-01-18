package datastore

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/ryicoh/clean-arch/internal/adapter/interface/conf"
	"github.com/ryicoh/clean-arch/internal/adapter/interface/datastore"
)

type database struct {
	db *gorm.DB
}

func NewDBConfigFromConfig(cnf conf.DatabaseConfig) *datastore.DBConfig {
	return &datastore.DBConfig{
		Host:     cnf.Host,
		User:     cnf.User,
		Password: cnf.Password,
		Name:     cnf.Name,
		Protocol: cnf.Protocol,
	}
}

// New defines ...
func New(c *datastore.DBConfig) (datastore.DB, error) {
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

func (d *database) Limit(limit int) datastore.DB {
	return &database{d.db.Limit(limit)}
}

func (d *database) Offset(offset int) datastore.DB {
	return &database{d.db.Limit(offset)}
}

func (d *database) Find(i interface{}) datastore.DB {
	return &database{d.db.Find(i)}
}

func (d *database) Error() error {
	return d.db.Error
}
