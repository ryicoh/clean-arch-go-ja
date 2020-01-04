package echo

import (
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo"
	"github.com/ryicoh/clean-arch/internal/infrastructure/conf"
)

type context struct {
	ctx echo.Context
	cnf conf.Config
}

func (c *context) GetConfig() conf.Config {
	return c.cnf
}

func (c *context) GetQueryParam(query string) string {
	return c.ctx.QueryParam(query)
}

func (c *context) Bind(i interface{}) error {
	return c.Bind(i)
}

func (c *context) GetRequest() *http.Request {
	return c.ctx.Request()
}

func (c *context) GetMultipartForm() (*multipart.Form, error) {
	return c.ctx.MultipartForm()
}
