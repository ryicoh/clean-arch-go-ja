package datastore

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
