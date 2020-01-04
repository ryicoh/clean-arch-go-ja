package appcontext

import (
	"errors"
	"strconv"
)

func (c *context) GetQueryOffset() (int, error) {
	str := c.ctx.GetQueryParam("Offset")
	if str == "" {
		return 0, errors.New("query parameter 'Offset' is empty")
	}

	offset, err := strconv.Atoi(str)
	if err != nil {
		return 0, errors.New("query parameter 'Offset' is invalid")
	}

	return offset, nil
}
