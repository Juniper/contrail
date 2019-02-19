package analytics

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/collector"
)

const (
	typeRESTAPITrace = "RestApiTrace"
	typeVNCAPIDebug  = "VncApiDebug"
	typeVNCAPIError  = "VncApiError"
	typeVNCAPIInfo   = "VncApiInfo"
	typeVNCAPINotice = "VncApiNotice"
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

type payloadVNCAPIMessage struct {
	messageType string
	Message     string `json:"api_msg"`
}

// RESTAPITrace sends message with type RestApiTrace
func RESTAPITrace(ctx echo.Context, reqBody, resBody []byte) collector.MessageBuilder {
	/* TODO: Should be reverted as introspect service for Intent API will be introduced.
	if ctx.Request().Method == "GET" {
		return collector.NewEmptyMessageBuilder()
	}
	return &payloadRESTAPITrace{
		RequestID:    "req-" + uuid.NewV4().String(),
		URL:          ctx.Request().URL.String(),
		Method:       ctx.Request().Method,
		RequestData:  string(reqBody),
		Status:       strconv.Itoa(ctx.Response().Status) + " " + http.StatusText(ctx.Response().Status),
		ResponseBody: string(resBody),
		RequestError: "",
	}
	*/
	return collector.NewEmptyMessageBuilder()
}

func (p *payloadRESTAPITrace) Build() *collector.Message {
	return &collector.Message{
		SandeshType: typeRESTAPITrace,
		Payload:     p,
	}
}

// VNCAPIMessage sends message with type VncApiDebug, VncApiInfo, VncApiNotice or
// VncApiError depends on level
func VNCAPIMessage(entry *logrus.Entry) collector.MessageBuilder {
	var messageType string
	switch entry.Level {
	case logrus.DebugLevel:
		messageType = typeVNCAPIDebug
	case logrus.InfoLevel:
		messageType = typeVNCAPIInfo
	case logrus.WarnLevel:
		messageType = typeVNCAPINotice
	default:
		messageType = typeVNCAPIError
	}
	return &payloadVNCAPIMessage{
		messageType: messageType,
		Message:     entry.Message,
	}
}

func (p *payloadVNCAPIMessage) Build() *collector.Message {
	return &collector.Message{
		SandeshType: p.messageType,
		Payload:     p,
	}
}
