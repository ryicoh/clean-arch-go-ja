package appcontext

import (
	"errors"
	"mime/multipart"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/form"
	"github.com/ryicoh/clean-arch/internal/infrastructure/datastore"
	"github.com/ryicoh/clean-arch/pkg"
)

const (
	mimeApplicationJSON = "application/json"
	mimeMultipartForm   = "multipart/form-data"
)

type binder struct{}

// newBinder はbinderを返します。
func newBinder() *binder {
	return &binder{}
}

// Bind 関数は、リクエストをモデルにBindingし、Validateもする関数です。
// 対応してるContext-Typeは、application/jsonとmultipart/form-dataです。
// ファイル（バイナリー）を含む場合は、フィールド名を${File}にしフィールド名+FileのKeyでバイナリを追加してください。
// 追加しない場合は、FileFormInvalidErrorが出ます。
func (cb *binder) Bind(i interface{}, c Context) error {
	if err := cb.BindWithoutFileAndValidation(i, c); err != nil {
		return err
	}

	if err := cb.validate(i, c); err != nil {
		return err
	}

	if err := cb.BindFile(i, c); err != nil {
		return err
	}

	return nil
}

func (cb *binder) BindWithoutFileAndValidation(i interface{}, c Context) error {
	if err := c.MinimumBind(i); err != nil {
		return err
	}

	var err error
	switch cb.getContentType(c.GetRequest().Header) {
	case mimeApplicationJSON:
		return nil
	case mimeMultipartForm:
		err = cb.bindMultiPartFormWithoutFile(i, c)
	}

	if err != nil {
		return err
	}

	return nil
}

func (cb *binder) validate(i interface{}, c Context) error {
	errs, err := c.ValidateAndGetErrorFields(i)
	if err != nil {
		return err
	}

	if errs != nil {
		return err
	}

	return nil
}

func (cb *binder) getContentType(h http.Header) string {
	ct := h.Get("Content-Type")

	if i := strings.Index(ct, ";"); i > 0 {
		ct = ct[:i]
	}

	return ct
}

func (cb *binder) bindMultiPartFormWithoutFile(it interface{}, c Context) error {
	mf, err := c.GetMultipartForm()
	if err != nil {
		return err
	}

	decoder := form.NewDecoder()
	if err := decoder.Decode(it, mf.Value); err != nil {
		return err
	}

	return nil
}

func (cb *binder) BindFile(it interface{}, c Context) error {
	if cb.getContentType(c.GetRequest().Header) != mimeMultipartForm {
		return nil
	}

	mf, err := c.GetMultipartForm()
	if err != nil {
		return err
	}

	for name, headers := range mf.File {
		err = cb.bindFile(name, headers, it)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cb *binder) bindFile(name string, headers []*multipart.FileHeader, it interface{}) error {
	if len(headers) == 0 {
		return nil
	}

	if len(headers) != 1 {
		return errors.New("unsupported multiple files")
	}

	matched, err := regexp.MatchString(`.*File$`, name)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("file key is invalid")
	}

	field, err := pkg.FindField(reflect.ValueOf(it), strings.TrimSuffix(name, "File"))
	if err != nil {
		return err
	}

	if field.Kind() != reflect.String {
		return errors.New("field is not a string")
	}

	if field.String() != "${File}" {
		return errors.New("the field must be set ${File}")
	}

	file, err := headers[0].Open()
	if err != nil {
		return err
	}

	filepath, err := datastore.Upload(headers[0].Filename, file, "")
	if err != nil {
		return err
	}

	field.SetString(filepath)

	return nil
}
