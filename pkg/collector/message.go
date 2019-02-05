package collector

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

const (
	typeRESTAPITrace = "RestApiTrace"
)

type payloadRESTAPITrace struct {
	RequestID    string `json:"request_id"`
	URL          string `json:"url"`
	Method       string `json:"method"`
	RequestData  string `json:"request_data"`
	Status       string `json:"status"`
	ResponseBody string `json:"response_body"`
	RequestError string `json:"request_error"`
}

// RESTAPITrace sends message with type RestApiTrace
func (collector *Collector) RESTAPITrace(c echo.Context, reqBody, resBody []byte) {
	if c.Request().Method == "GET" {
		return
	}
	collector.sendMessage(&message{
		SandeshType: typeRESTAPITrace,
		Payload: &payloadRESTAPITrace{
			RequestID:    "req-" + uuid.NewV4().String(),
			URL:          c.Request().URL.String(),
			Method:       c.Request().Method,
			RequestData:  string(reqBody),
			Status:       strconv.Itoa(c.Response().Status) + " " + http.StatusText(c.Response().Status),
			ResponseBody: string(resBody),
			RequestError: "",
		},
	})
}
