package echov4

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/ryicoh/clean-arch/internal/adapter/interface/conf"
	"gopkg.in/boj/redistore.v1"
)

func newSessionMiddleware(cnf conf.Config) (echo.MiddlewareFunc, error) {
	redisConfing := cnf.GetRedisConfig()
	store, err := redistore.NewRediStore(
		32, "tcp",
		redisConfing.Host+":"+redisConfing.Port,
		redisConfing.Password,
		[]byte("secret-key"),
	)

	if err != nil {
		return nil, err
	}

	store.SetMaxAge(60 * 60 * 24 * 365)

	return session.Middleware(store), nil
}
