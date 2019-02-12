package collector

import (
	"net/http"
	"strconv"
	"testing"

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

type mockSender struct {
	message *message
}

func (s *mockSender) sendMessage(m *message) {
	s.message = m
}

func TestDBRequestTrace(t *testing.T) {
	tests := []struct {
		name      string
		operation string
		v         map[string]interface{}
	}{
		{
			name:      "Create DBRequestTrace message",
			operation: "create",
			v: map[string]interface{}{
				"type":   "project",
				"uuid":   "created_project_uuid",
				"fqname": []string{"default-domain", "default-project"},
			},
		},
		{
			name:      "Update DBRequestTrace message",
			operation: "update",
			v: map[string]interface{}{
				"type":   "project",
				"uuid":   "updated_project_uuid",
				"fqname": []string{"default-domain", "default-project"},
			},
		},
		{
			name:      "Delete DBRequestTrace message",
			operation: "delete",
			v: map[string]interface{}{
				"type":   "project",
				"uuid":   "deleted_project_uuid",
				"fqname": []string{"default-domain", "default-project"},
			},
		},
	}

	c, err := NewCollector(&Config{})
	assert.NoError(t, err)
	s := &mockSender{}
	c.sender = s

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.DBRequestTrace(tt.operation, tt.v)

			assert.Equal(t, s.message.SandeshType, typeDBRequestTrace)
			m, ok := s.message.Payload.(*payloadDBRequestTrace)
			assert.True(t, ok)
			assert.Equal(t, m.Operation, tt.operation)
			assert.Equal(t, m.Body, tt.v)
		})
	}
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

	c, err := NewCollector(&Config{})
	assert.NoError(t, err)
	s := &mockSender{}
	c.sender = s

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

			c.RESTAPITrace(mockEchoContent, []byte(tt.request), []byte(tt.response))

			assert.Equal(t, s.message.SandeshType, typeRESTAPITrace)
			m, ok := s.message.Payload.(*payloadRESTAPITrace)
			assert.True(t, ok)
			assert.Equal(t, m.URL, tt.url)
			assert.Equal(t, m.Method, tt.method)
			assert.Equal(t, m.RequestData, tt.request)
			assert.Equal(t, m.Status, strconv.Itoa(tt.status)+" "+http.StatusText(tt.status))
			assert.Equal(t, m.ResponseBody, tt.response)
		})
	}
}
