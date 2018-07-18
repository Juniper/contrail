package discovery

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDiscovery(t *testing.T) {
	discoHandler := NewHandler("addr")

	discoHandler.Register("/path1", "GET", "test1", "rel1")
	discoHandler.Register("path2", "", "test2", "rel2")

	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(httptest.NewRequest(echo.GET, "/", nil), rec)
	require.NoError(t, discoHandler.Handle(c))
	require.Equal(t, http.StatusOK, rec.Code)

	expected := `
	{
		"href": "addr",
		"links": [
			{ "href": "addr/path1", "method": "GET", "name": "test1", "rel": "rel1" },
			{ "href": "addr/path2", "method": null, "name": "test2", "rel": "rel2" }
		]
	}`
	assert.JSONEq(t, expected, rec.Body.String())
}
