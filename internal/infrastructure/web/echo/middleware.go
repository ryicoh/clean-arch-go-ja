package echo

import (
	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"github.com/ryicoh/clean-arch/internal/adapter/interface/conf"
)

func newSessionMiddleware(cnf conf.Config) (echo.MiddlewareFunc, error) {
	redisConfing := cnf.GetRedisConfig()
	store, err := session.NewRedisStore(
		32, "tcp",
		redisConfing.Host+":"+redisConfing.Port,
		redisConfing.Password,
		[]byte("secret"),
	)
	if err != nil {
		return nil, err
	}

	store.MaxAge(31104000)
	return session.Sessions("session", store), nil
}
