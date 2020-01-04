package datastore

import "github.com/ryicoh/clean-arch/internal/infrastructure/conf/yaml"

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
		Find(interface{}) DB
		Error() error
	}
)

func NewDBConfigFromENV() *DBConfig {
	cnf := yaml.New().GetDatabaseConfig()

	return &DBConfig{
		Host:     cnf.Host,
		User:     cnf.User,
		Password: cnf.Password,
		Name:     cnf.Name,
		Protocol: cnf.Protocol,
	}
}
