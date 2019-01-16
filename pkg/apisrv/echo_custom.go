package apisrv

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/labstack/echo"
)

type customBinder struct{}

func (*customBinder) Bind(i interface{}, c echo.Context) (err error) {
	rq := c.Request()
	ct := rq.Header.Get(echo.HeaderContentType)
	err = echo.ErrUnsupportedMediaType
	if !strings.HasPrefix(ct, echo.MIMEApplicationJSON) {
		// default binder
		db := new(echo.DefaultBinder)
		return db.Bind(i, c)
	}
	// json parsing
	dec := json.NewDecoder(rq.Body)
	dec.UseNumber()
	err = dec.Decode(i)
	if err == io.EOF {
		return nil
	}
	return err
}
