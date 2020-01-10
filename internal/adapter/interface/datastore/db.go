package datastore

import "github.com/ryicoh/clean-arch/internal/adapter/interface/conf"

type (
	DBConfig struct {
		Host     string
		User     string
		Password string
		Name     string
		Protocol string
	}

	DB interface {
		AutoMigrate(...interface{}) DB
		Limit(int) DB
		Offset(int) DB
		Find(interface{}) DB
		Error() error
	}
)

func NewDBConfigFromConfig(cnf conf.DatabaseConfig) *DBConfig {
	return &DBConfig{
		Host:     cnf.Host,
		User:     cnf.User,
		Password: cnf.Password,
		Name:     cnf.Name,
		Protocol: cnf.Protocol,
	}
}
