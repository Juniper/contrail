package collector

import (
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestVNCAPIStatsLog(t *testing.T) {
	tests := []struct {
		name       string
		objectType string
		method     string
		domain     string
		project    string
		user       string
		agent      string
		requestID  string
		status     int
		response   string
	}{
		{
			name:       "post VNCAPIStatsLog message",
			objectType: "project",
			method:     "POST",
			domain:     "domain_uuid",
			project:    "project_uuid",
			user:       "alice",
			agent:      "go-test-client",
			requestID:  "req-001",
			status:     http.StatusOK,
			response:   "POST Response",
		},
		{
			name:       "update VNCAPIStatsLog message",
			objectType: "project",
			method:     "PUT",
			domain:     "red_domain_uuid",
			project:    "red_project_uuid",
			user:       "alice",
			agent:      "go-test-client",
			requestID:  "req-002",
			status:     http.StatusOK,
			response:   "DELETE Response",
		},
		{
			name:       "get VNCAPIStatsLog message",
			objectType: "project",
			method:     "GET",
			domain:     "red_domain_uuid",
			project:    "red_project_uuid",
			user:       "alice",
			agent:      "go-test-client",
			requestID:  "req-003",
			status:     http.StatusOK,
			response:   "DELETE Response",
		},
		{
			name:       "delete VNCAPIStatsLog message",
			objectType: "virtual-network",
			method:     "DELETE",
			domain:     "blue_domain_uuid",
			project:    "blue_project_uuid",
			user:       "bob",
			agent:      "go-test-client",
			requestID:  "req-004",
			status:     http.StatusNotFound,
			response:   "DELETE Response",
		},
	}

	e := echo.New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, "http://localhost/proxy_url", nil)
			assert.NoError(t, err)
			ctx := e.NewContext(req, nil)
			ctx.Set(keyVNCAPIStatsLogTimeStamp, time.Now())
			ctx.Set(keyVNCAPIStatsLogObjectType, tt.objectType)
			ctx.Request().Header.Set("X-Domain-Name", tt.domain)
			ctx.Request().Header.Set("X-Project-Name", tt.project)
			ctx.Request().Header.Set("X-User-Name", tt.user)
			ctx.Request().Header.Set("X-Contrail-Useragent", tt.agent)
			ctx.Request().Header.Set("X-Request-Id", tt.requestID)
			resp := echo.NewResponse(nil, e)
			resp.Status = tt.status
			mockEchoContent := &mockEcho{
				Context:  ctx,
				response: resp,
			}

			c := &mockCollector{}

			c.Send(VNCAPIStatsLog(mockEchoContent, nil, []byte(tt.response)))

			assert.Equal(t, c.message.SandeshType, typeVNCAPIStatsLog)
			m, ok := c.message.Payload.(*payloadVNCAPIStatsLog)
			assert.True(t, ok)
			assert.Equal(t, m.OperationType, tt.method)
			assert.Equal(t, m.User, tt.user)
			assert.Equal(t, m.UserAgent, tt.agent)
			assert.Equal(t, m.DomainName, tt.domain)
			assert.Equal(t, m.ProjectName, tt.project)
			assert.Equal(t, m.ObjectType, tt.objectType)
			assert.True(t, m.ResponseTimeInUSec > 0)
			assert.Equal(t, m.ResponseSize, len(tt.response))
			assert.Equal(t, m.ResponseCode, strconv.Itoa(tt.status))
		})
	}
}
