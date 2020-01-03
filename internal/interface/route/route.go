package route

import (
	"github.com/ryicoh/clean-arch/internal/infrastructure/web"
	"github.com/ryicoh/clean-arch/internal/interface/controller"
)

func Register(s web.Server, ctrler controller.AppController) {
	apiv1 := s.Group("/api/v1")

	{
		user := apiv1.Group("/users")
		user.GET("", ctrler.GetUsers)
	}
}
