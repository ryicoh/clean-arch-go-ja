package route

import (
	"github.com/ryicoh/clean-arch/internal/infrastructure/web"
	"github.com/ryicoh/clean-arch/internal/interface/controller"
	"github.com/ryicoh/clean-arch/internal/interface/web/appcontext"
)

type (
	handler func(appcontext.Context) (code int, i interface{}, err error)
)

func Register(s web.Server, ctrler controller.AppController) {
	apiv1 := s.Group("/api/v1")

	{
		user := apiv1.Group("/users")
		user.GET("", castWebHandler(ctrler.GetUsers))
	}
}

func castWebHandler(h handler) web.Handler {
	return func(c web.Context) (code int, data interface{}, err error) {
		return h(appcontext.New(c))
	}
}
