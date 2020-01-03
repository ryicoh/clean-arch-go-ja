package web

import "github.com/ryicoh/clean-arch/internal/infrastructure/conf"

type (
	Server interface {
		Start(port int) error
		GetConfig() conf.Config
		Group
	}

	Group interface {
		Group(path string) Group
		GET(path string, handler Handler)
		POST(path string, handler Handler)
	}

	Handler func(Context) (code int, data interface{}, errr error)

	Context interface {
		GetConfig() conf.Config
		GetQueryOffset() (int, error)
	}
)