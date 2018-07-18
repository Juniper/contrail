package discovery

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestDiscovery(t *testing.T) {
	type register struct {
		path   string
		method string
		name   string
		rel    string
	}

	tests := []struct {
		addr      string
		registers []register
		expected  string
	}{
		{
			"addr",
			[]register{
				{"/path1", "GET", "test1", "rel1"},
				{"path2", "", "test2", "rel2"},
			},
			`{
				"href": "addr",
				"links": [
					{"link": { "href": "addr/path1", "method": "GET", "name": "test1", "rel": "rel1" }},
					{"link": { "href": "addr/path2", "method": null, "name": "test2", "rel": "rel2" }}
				]
			}`,
		},
		{
			"http://localhost:8082/",
			[]register{
				{"/path1", "GET", "test1", "rel1"},
				{"path2", "", "test2", "rel2"},
			},
			`{
				"href": "http://localhost:8082",
				"links": [
					{"link": { "href": "http://localhost:8082/path1", "method": "GET", "name": "test1", "rel": "rel1" }},
					{"link": { "href": "http://localhost:8082/path2", "method": null, "name": "test2", "rel": "rel2" }}
				]
			}`,
		},
	}

	for _, tt := range tests {
		dh := NewHandler(tt.addr)

		for _, l := range tt.registers {
			dh.Register(l.path, l.method, l.name, l.rel)
		}

		e := echo.New()
		rec := httptest.NewRecorder()
		c := e.NewContext(httptest.NewRequest(echo.GET, "/", nil), rec)
		err := dh.Handle(c)

		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, tt.expected, rec.Body.String())
		}
	}
}
