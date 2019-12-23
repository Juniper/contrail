package baseapisrv

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/Juniper/asf/pkg/client"
	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func gRPCMiddleware(grpcServer http.Handler) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
			w := c.Response()
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				grpcServer.ServeHTTP(w, r)
				return nil
			}
			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}
	}
}

type customBinder struct{}

func (*customBinder) Bind(i interface{}, c echo.Context) (err error) {
	rq := c.Request()
	ct := rq.Header.Get(echo.HeaderContentType)
	err = echo.ErrUnsupportedMediaType
	if !strings.HasPrefix(ct, echo.MIMEApplicationJSON) {
		db := new(echo.DefaultBinder)
		return db.Bind(i, c)
	}

	dec := json.NewDecoder(rq.Body)
	dec.UseNumber()
	err = dec.Decode(i)
	if err == io.EOF {
		return nil
	}
	return err
}

func recorderMiddleware(log *logrus.Entry) echo.MiddlewareFunc {
	file := viper.GetString("recorder.file")
	scenario := &struct {
		Workflow []*recorderTask `yaml:"workflow,omitempty"`
	}{}
	var mutex sync.Mutex
	return middleware.BodyDump(func(c echo.Context, requestBody, responseBody []byte) {
		var data interface{}
		err := json.Unmarshal(requestBody, &data)
		if err != nil {
			log.WithError(err).Error("Malformed JSON input")
		}
		var expected interface{}
		err = json.Unmarshal(responseBody, &expected)
		if err != nil {
			log.WithError(err).Error("Malformed JSON response")
		}
		task := &recorderTask{
			Request: &client.Request{
				Method:   c.Request().Method,
				Path:     c.Request().URL.Path,
				Expected: []int{c.Response().Status},
				Data:     data,
			},
			Expect: expected,
		}
		mutex.Lock()
		defer mutex.Unlock()
		scenario.Workflow = append(scenario.Workflow, task)
		err = fileutil.SaveFile(file, scenario)
		if err != nil {
			log.WithError(err).Error("Failed to save scenario to file")
		}
	})
}
