package collector

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

// message types in alphabetical order
const (
	typeDBRequestTrace        = "DBRequestTrace"
	typeMessageBusNotifyTrace = "MessageBusNotifyTrace"
	typeRESTAPITrace          = "RestApiTrace"
)

type payloadDBRequestTrace struct {
	RequestID string      `json:"request_id"`
	Operation string      `json:"operation"`
	Body      interface{} `json:"body"`
	Error     string      `json:"error"`
}

type payloadMessageBusNotifyTrace struct {
	RequestID string                     `json:"request_id"`
	Operation string                     `json:"operation"`
	Body      *bodyMessageBusNotifyTrace `json:"body"`
	Error     string                     `json:"error"`
}

type bodyMessageBusNotifyTrace struct {
	Operation string   `json:"oper"`
	RequestID string   `json:"request-id"`
	Type      string   `json:"type"`
	UUID      string   `json:"uuid"`
	FQName    []string `json:"fq_name"`
}

type payloadRESTAPITrace struct {
	RequestID    string `json:"request_id"`
	URL          string `json:"url"`
	Method       string `json:"method"`
	RequestData  string `json:"request_data"`
	Status       string `json:"status"`
	ResponseBody string `json:"response_body"`
	RequestError string `json:"request_error"`
}

// DBRequestTrace sends message with type DBRequestTrace
func (c *Collector) DBRequestTrace(operation string, v interface{}) {
	c.sendMessage(&message{
		SandeshType: typeDBRequestTrace,
		Payload: &payloadDBRequestTrace{
			RequestID: "req-" + uuid.NewV4().String(),
			Operation: operation,
			Body:      v,
			Error:     "",
		},
	})
}

// MessageBusNotifyTrace sends message with type MessageBusNotifyTrace
func (c *Collector) MessageBusNotifyTrace(operation string, typeName string, objUUID string, objFQName []string) {
	requestID := "req-" + uuid.NewV4().String()
	c.sendMessage(&message{
		SandeshType: typeMessageBusNotifyTrace,
		Payload: &payloadMessageBusNotifyTrace{
			RequestID: requestID,
			Operation: operation,
			Body: &bodyMessageBusNotifyTrace{
				Operation: operation,
				RequestID: requestID,
				Type:      typeName,
				UUID:      objUUID,
				FQName:    objFQName,
			},
			Error: "",
		},
	})
}

// RESTAPITrace sends message with type RestApiTrace
func (c *Collector) RESTAPITrace(ctx echo.Context, reqBody, resBody []byte) {
	if ctx.Request().Method == "GET" {
		return
	}
	c.sendMessage(&message{
		SandeshType: typeRESTAPITrace,
		Payload: &payloadRESTAPITrace{
			RequestID:    "req-" + uuid.NewV4().String(),
			URL:          ctx.Request().URL.String(),
			Method:       ctx.Request().Method,
			RequestData:  string(reqBody),
			Status:       strconv.Itoa(ctx.Response().Status) + " " + http.StatusText(ctx.Response().Status),
			ResponseBody: string(resBody),
			RequestError: "",
		},
	})
}
