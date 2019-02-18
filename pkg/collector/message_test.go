package collector

import (
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type mockEcho struct {
	echo.Context
	response *echo.Response
}

func (m *mockEcho) Response() *echo.Response {
	return m.response
}

type mockCollector struct {
	message *Message
}

func (c *mockCollector) Send(b MessageBuilder) {
	c.message = b.Build()
}

func TestRESTAPITrace(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		url      string
		status   int
		request  string
		response string
	}{
		{
			name:     "post RESTAPITrace message",
			method:   "POST",
			url:      "http://localhost/proxy_url",
			status:   http.StatusOK,
			request:  "POST Request",
			response: "POST Response",
		},
		{
			name:     "delete RESTAPITrace message",
			method:   "DELETE",
			url:      "http://localhost/proxy_url",
			status:   http.StatusOK,
			request:  "DELETE Request",
			response: "DELETE Response",
		},
	}

	e := echo.New()

	for _, tt := range tests {
		t.Run("RESTAPITrace", func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.url, nil)
			assert.NoError(t, err)
			ctx := e.NewContext(req, nil)
			resp := echo.NewResponse(nil, e)
			resp.Status = tt.status
			mockEchoContent := &mockEcho{
				Context:  ctx,
				response: resp,
			}

			messageBuilder := RESTAPITrace(mockEchoContent, []byte(tt.request), []byte(tt.response))
			assert.NotNil(t, messageBuilder)
			message := messageBuilder.Build()
			assert.Nil(t, message)

			/* TODO: Should be reverted as introspect service for Intent API will be introduced.
			assert.Equal(t, message.SandeshType, typeRESTAPITrace)
			messageBuilder := RESTAPITrace(mockEchoContent, []byte(tt.request), []byte(tt.response))
			assert.NotNil(t, messageBuilder)
			message := messageBuilder.Build()
			assert.NotNil(t, message)
			m, ok := message.Payload.(*payloadRESTAPITrace)
			assert.True(t, ok)
			assert.Equal(t, m.URL, tt.url)
			assert.Equal(t, m.Method, tt.method)
			assert.Equal(t, m.RequestData, tt.request)
			assert.Equal(t, m.Status, strconv.Itoa(tt.status)+" "+http.StatusText(tt.status))
			assert.Equal(t, m.ResponseBody, tt.response)
			*/
		})
	}
}

func TestVNCAPIMessage(t *testing.T) {
	tests := []struct {
		call    func()
		level   string
		message string
	}{
		{
			call: func() {
				logrus.Debugf("debug message")
			},
			level:   typeVNCAPIDebug,
			message: "debug message",
		},
		{
			call: func() {
				logrus.Infof("message")
			},
			level:   typeVNCAPIInfo,
			message: "message",
		},
		{
			call: func() {
				logrus.Warnf("warning message")
			},
			level:   typeVNCAPINotice,
			message: "warning message",
		},
		{
			call: func() {
				logrus.Errorf("error message")
			},
			level:   typeVNCAPIError,
			message: "error message",
		},
	}

	c := &mockCollector{}
	AddLoggerHook(c)
	logrus.SetLevel(logrus.DebugLevel)

	for _, tt := range tests {
		t.Run("APIMessage", func(t *testing.T) {
			tt.call()
			assert.NotNil(t, c.message)
			assert.Equal(t, c.message.SandeshType, tt.level)
			m, ok := c.message.Payload.(*payloadVNCAPIMessage)
			assert.True(t, ok)
			assert.Equal(t, m.Message, tt.message)
		})
	}

	c.message = nil
	ignoreAPIMessage().Errorf("message")
	assert.Nil(t, c.message)
	ignoreAPIMessage().Warn("warning message")
	assert.Nil(t, c.message)
}
