package appcontext

import (
	"mime/multipart"
	"net/http"

	"github.com/ryicoh/clean-arch/internal/infrastructure/web"
)

type (
	ErrorField struct {
	}

	Handler func(Context) (code int, data interface{}, errr error)

	Context interface {
		GetQueryOffset() (int, error)
		BindAndValidate(i interface{}) error
		ValidateAndGetErrorFields(i interface{}) ([]*ErrorField, error)

		// MinimumBind は再帰、ファイルを無視してバインディングします。
		MinimumBind(i interface{}) error

		GetRequest() *http.Request
		GetMultipartForm() (*multipart.Form, error)
	}

	context struct {
		ctx    web.Context
		binder *binder
	}
)

func New(ctx web.Context) Context {
	return &context{ctx: ctx, binder: newBinder()}
}

func (c *context) BindAndValidate(i interface{}) error {
	return nil
}

func (c *context) MinimumBind(i interface{}) error {
	return c.ctx.Bind(i)
}

func (c *context) GetRequest() *http.Request {
	return c.GetRequest()
}

func (c *context) GetMultipartForm() (*multipart.Form, error) {
	return c.GetMultipartForm()
}

func (c *context) ValidateAndGetErrorFields(i interface{}) ([]*ErrorField, error) {
	return nil, nil
}
