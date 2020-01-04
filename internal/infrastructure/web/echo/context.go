package echo

import (
	"fmt"
	"mime/multipart"
	"net/http"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"github.com/ryicoh/clean-arch/internal/infrastructure/conf"
	"github.com/spf13/cast"
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

func (c *context) GetSessionValue(key string) (string, error) {
	sess := session.Default(c.ctx)
	v := sess.Get(key)

	if v == nil {
		return "", fmt.Errorf("%q session not found", key)
	}

	return cast.ToString(v), nil
}

func (c *context) SetSessionValue(key, value string) error {
	sess := session.Default(c.ctx)
	sess.Set(key, value)

	if err := sess.Save(); err != nil {
		return err
	}

	return nil
}
