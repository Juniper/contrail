package baseapisrv

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestDiscovery(t *testing.T) {
	type register struct {
		method      string
		path        string
		optionFuncs []RouteOptionFunc
	}

	tests := []struct {
		host      string
		registers []register
		expected  string
	}{
		{
			"addr",
			[]register{
				{
					method:      "GET",
					path:        "/path1",
					optionFuncs: []RouteOptionFunc{WithHomepageName("test1"), WithHomepageType("rel1")},
				},
				{
					method:      "",
					path:        "path2",
					optionFuncs: []RouteOptionFunc{WithHomepageName("test2"), WithHomepageType("rel2")},
				},
			},
			`{
				"href": "http://addr",
				"links": [
					{"link": { "href": "http://addr/path1", "method": "GET", "name": "test1", "rel": "rel1" }},
					{"link": { "href": "http://addr/path2", "method": null, "name": "test2", "rel": "rel2" }}
				]
			}`,
		},
		{
			"localhost:8082",
			[]register{
				{
					method:      "GET",
					path:        "/path1",
					optionFuncs: []RouteOptionFunc{WithHomepageName("test1"), WithHomepageType("rel1")},
				},
				{
					method:      "",
					path:        "path2",
					optionFuncs: []RouteOptionFunc{WithHomepageName("test2"), WithHomepageType("rel2")},
				},
			},
			`{
				"href": "http://localhost:8082",
				"links": [
					{"link": { "href": "http://localhost:8082/path1", "method": "GET", "name": "test1", "rel": "rel1" }},
					{"link": { "href": "http://localhost:8082/path2", "method": null, "name": "test2", "rel": "rel2" }}
				]
			}`,
		},

		{
			"addr",
			[]register{
				{
					method: "POST",
					path:   "/resources",
					optionFuncs: []RouteOptionFunc{
						WithNoHomepageMethod(),
						WithHomepageName("resource"),
						WithHomepageType(CollectionEndpoint),
					},
				},
				{
					method: "GET",
					path:   "/resources",
					optionFuncs: []RouteOptionFunc{
						WithNoHomepageMethod(),
						WithHomepageName("resource"),
						WithHomepageType(CollectionEndpoint),
					},
				},
			},
			`{
				"href": "http://addr",
				"links": [
					{"link": { "href": "http://addr/resources", "method": null, "name": "resource", "rel": "collection" }}
				]
			}`,
		},

		{
			"addr",
			[]register{
				{
					method: "PUT",
					path:   "/resource/:id",
					optionFuncs: []RouteOptionFunc{
						WithNoHomepageMethod(),
					},
				},
				{
					method: "GET",
					path:   "/resource/:id",
					optionFuncs: []RouteOptionFunc{
						WithNoHomepageMethod(),
					},
				},
				{
					method: "DELETE",
					path:   "/resource/:id",
					optionFuncs: []RouteOptionFunc{
						WithNoHomepageMethod(),
					},
				},
			},
			`{
				"href": "http://addr",
				"links": [
					{"link": { "href": "http://addr/resource", "method": null, "name": "resource", "rel": "resource-base" }}
				]
			}`,
		},

		{
			"addr",
			[]register{
				{
					method: "POST",
					path:   "/some-action",
				},
			},
			`{
				"href": "http://addr",
				"links": [
					{"link": { "href": "http://addr/some-action", "method": "POST", "name": "some-action", "rel": "action" }}
				]
			}`,
		},

		{
			"addr",
			[]register{
				{
					method:      "GET",
					path:        "/keystone/v3/projects",
					optionFuncs: []RouteOptionFunc{WithHomepageType(CollectionEndpoint)},
				},
				{
					method: "GET",
					path:   "/keystone/v3/projects/:id",
				},
			},
			`{
				"href": "http://addr",
				"links": [
					{"link": { "href": "http://addr/keystone/v3/projects", "method": "GET", "name": "keystone/v3/projects", "rel": "collection" }},
					{"link": { "href": "http://addr/keystone/v3/projects", "method": "GET", "name": "keystone/v3/projects", "rel": "resource-base" }}
				]
			}`,
		},
	}

	for _, tt := range tests {
		dh := NewHomepageHandler()

		for _, r := range tt.registers {
			options := makeRouteOptions(r.optionFuncs, r.method, r.path)
			dh.Register(options.homepageEntry)
		}

		e := echo.New()
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(echo.GET, "/", nil)
		req.Host = tt.host
		c := e.NewContext(req, rec)
		err := dh.Handle(c)

		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, tt.expected, rec.Body.String())
		}
	}
}
