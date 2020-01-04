package echov4

import (
	"mime/multipart"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
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

func (c *context) GetSession(key string) (string, error) {
	sess, err := session.Get("session", c.ctx)
	if err != nil {
		return "", err
	}

	v := sess.Values[key]
	return cast.ToString(v), nil
}

func (c *context) SetSession(key, value string) error {
	sess, _ := session.Get("session", c.ctx)

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 365,
		HttpOnly: true,
	}
	sess.Values[key] = value

	err := sess.Save(c.GetRequest(), c.ctx.Response())
	if err != nil {
		return err
	}

	return nil
}
