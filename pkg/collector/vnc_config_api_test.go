package collector

import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

func TestVNCAPIConfigLog(t *testing.T) {
	tests := []struct {
		name      string
		metadata  *basemodels.Metadata
		operation string
		error     string
		url       string
		method    string
		agent     string
		remoteIP  string
		body      string
		domain    string
		project   string
		user      string
	}{
		{
			name: "update project",
			metadata: &basemodels.Metadata{
				UUID:   "project_uuid",
				FQName: []string{"main", "second"},
				Type:   "project",
			},
			operation: "http_put",
			url:       "http://localhost/project/project_uuid",
			method:    "UPDATE",
			agent:     "Go-http-client/1.1",
			remoteIP:  "127.0.0.1",
			body:      "{\"test\":\"messsage\"}",
			domain:    "default-domain",
			project:   "default-project",
			user:      "username",
		},
		{
			name: "delete project with error",
			metadata: &basemodels.Metadata{
				UUID:   "project_uuid",
				FQName: []string{"main", "second"},
				Type:   "project",
			},
			operation: "delete",
			error:     "Error Message",
			url:       "http://localhost/project/project_uuid",
			method:    "DELETE",
			agent:     "Go-http-client/1.1",
			remoteIP:  "127.0.0.1",
			body:      "error message",
			domain:    "default-domain",
			project:   "default-project",
			user:      "username",
		},
	}

	e := echo.New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.url, nil)
			assert.NoError(t, err)
			ctx := e.NewContext(req, nil)
			ctx.Set(keyVNCAPIConfigLogMetadata, tt.metadata)
			ctx.Set(keyVNCAPIConfigLogOperation, tt.operation)
			ctx.Set(keyVNCAPIConfigLogError, tt.error)
			ctx.Request().Header.Set("X-Domain-Name", tt.domain)
			ctx.Request().Header.Set("X-Project-Name", tt.project)
			ctx.Request().Header.Set("X-User-Name", tt.user)
			ctx.Request().Header.Set("X-Contrail-Useragent", tt.agent)
			ctx.Request().Header.Set("Host", tt.remoteIP)
			resp := echo.NewResponse(nil, e)
			resp.Status = http.StatusOK
			mockEchoContent := &mockEcho{
				Context:  ctx,
				response: resp,
			}

			c := &mockCollector{}

			c.Send(VNCAPIConfigLog(mockEchoContent, nil, []byte(tt.body)))

			assert.Equal(t, c.message.SandeshType, typeVNCAPIConfigLog)
			m, ok := c.message.Payload.(*payloadVNCAPIConfigLog)
			assert.True(t, ok)
			assert.Equal(t, m.UUID, tt.metadata.UUID)
			assert.Equal(t, m.ObjectType, tt.metadata.Type)
			assert.Equal(t, m.FQName, basemodels.FQNameToString(tt.metadata.FQName))
			assert.Equal(t, m.URL, tt.url)
			assert.Equal(t, m.Operation, tt.operation)
			assert.Equal(t, m.UserAgent, tt.agent)
			assert.Equal(t, m.RemoteIP, tt.remoteIP)
			assert.Equal(t, m.Body, tt.body)
			assert.Equal(t, m.Domain, tt.domain)
			assert.Equal(t, m.Project, tt.project)
			assert.Equal(t, m.User, tt.user)
			assert.Equal(t, m.Error, tt.error)
		})
	}
}
