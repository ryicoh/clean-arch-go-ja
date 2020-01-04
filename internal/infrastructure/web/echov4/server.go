package echov4

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ryicoh/clean-arch/internal/infrastructure/conf"
	"github.com/ryicoh/clean-arch/internal/infrastructure/web"
)

type server struct {
	echo *echo.Echo
	cnf  conf.Config
}

func NewServer(cnf conf.Config) web.Server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	sessMid, err := newSessionMiddleware(cnf)
	if err != nil {
		panic(err)
	}

	e.Use(sessMid)

	return &server{echo: e, cnf: cnf}
}

func (s *server) newContext(c echo.Context) web.Context {
	return &context{ctx: c, cnf: s.cnf}
}

func (s *server) Start(port int) error {
	return s.echo.Start(fmt.Sprintf(":%d", port))
}

func (s *server) Group(path string) web.Group {
	return &group{server: s, group: s.echo.Group(path)}
}

func (s *server) GET(path string, handler web.Handler) {
	s.echo.GET(
		path,
		func(c echo.Context) error {
			code, res, err := handler(s.newContext(c))
			if err != nil {
				return echo.NewHTTPError(code, err.Error())
			}

			return c.JSON(http.StatusOK, res)
		},
	)
}

func (s *server) POST(path string, handler web.Handler) {
	s.echo.POST(
		path,
		func(c echo.Context) error {
			code, res, err := handler(s.newContext(c))
			if err != nil {
				return echo.NewHTTPError(code, err.Error())
			}

			return c.JSON(http.StatusOK, res)
		},
	)
}

func (c *server) GetConfig() conf.Config {
	return c.cnf
}
