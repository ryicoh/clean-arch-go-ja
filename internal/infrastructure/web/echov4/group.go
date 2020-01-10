package echov4

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ryicoh/clean-arch/internal/adapter/interface/web"
)

type group struct {
	server web.Server
	group  *echo.Group
}

func (g *group) newContext(c echo.Context) web.Context {
	return &context{ctx: c, cnf: g.server.GetConfig()}
}

func (g *group) Group(path string) web.Group {
	return &group{server: g.server, group: g.group.Group(path)}
}

func (g *group) GET(path string, handler web.Handler) {
	g.group.GET(path,
		func(c echo.Context) error {
			code, res, err := handler(g.newContext(c))
			if err != nil {
				return echo.NewHTTPError(code, err.Error())
			}

			return c.JSON(http.StatusOK, res)
		},
	)
}

func (g *group) POST(path string, handler web.Handler) {
	g.group.POST(path,
		func(c echo.Context) error {
			code, res, err := handler(g.newContext(c))
			if err != nil {
				return echo.NewHTTPError(code, err.Error())
			}

			return c.JSON(http.StatusOK, res)
		},
	)
}
