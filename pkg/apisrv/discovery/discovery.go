package discovery

import (
	"net/http"
	"path"

	"github.com/labstack/echo"
)

type link struct {
	Path   string  `json:"href"`
	Method *string `json:"method"`
	Name   string  `json:"name"`
	Rel    string  `json:"rel"`
}

type discovery struct {
	Addr  string  `json:"href"`
	Links []*link `json:"links"`
}

// Creates a new set of links to be discovered.
func NewDiscovery(addr string) *discovery { // nolint
	return &discovery{
		Addr: addr,
	}
}

// Adds a new link to the discovery.
func (d *discovery) Register(href string, method string, name string, rel string) {
	// path is assumed to be relative with respect to addr
	var m *string
	if method != "" {
		m = &method
	}

	d.Links = append(d.Links, &link{
		Path:   path.Join(d.Addr, href),
		Method: m,
		Name:   name,
		Rel:    rel,
	})
}

// Returns a handler to GET the links.
func MakeHandler(d *discovery) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, d)
	}
}
