package echo

import (
	"errors"
	"strconv"

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

func (c *context) GetQueryOffset() (int, error) {
	str := c.ctx.QueryParam("Offset")
	if str == "" {
		return 0, errors.New("Query parameter 'Offset' is empty")
	}

	offset, err := strconv.Atoi(str)
	if err != nil {
		return 0, errors.New("Query parameter 'Offset' is invalid")
	}

	return offset, nil
}
