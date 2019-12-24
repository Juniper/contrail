package baseapisrv

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

// HomepageHandler which serves a set of registered links.
type HomepageHandler struct {
	links []*link
}

type link struct {
	Link linkDetails `json:"link"`
}

type linkDetails struct {
	Path   string  `json:"href"`
	Method *string `json:"method"`
	Name   string  `json:"name"`
	Rel    string  `json:"rel"`
}

// NewHomepageHandler creates a new HomepageHandler.
func NewHomepageHandler() *HomepageHandler {
	return &HomepageHandler{}
}

// Register adds a new link to the HomepageHandler.
func (h *HomepageHandler) Register(path string, method string, name string, rel string) {
	// path is assumed to be relative with respect to addr
	path = strings.TrimPrefix(path, "/")

	var m *string
	if method != "" {
		m = &method
	}

	h.links = append(h.links, &link{
		Link: linkDetails{
			Path:   path,
			Method: m,
			Name:   name,
			Rel:    rel,
		},
	})
}

// Handle requests to return the links.
func (h *HomepageHandler) Handle(c echo.Context) error {
	r := c.Request()
	addr := GetRequestSchema(r) + r.Host

	var reply struct {
		Addr  string  `json:"href"`
		Links []*link `json:"links"`
	}

	reply.Addr = addr
	for _, l := range h.links {
		reply.Links = append(reply.Links, &link{
			Link: linkDetails{
				Path:   strings.Join([]string{addr, l.Link.Path}, "/"),
				Method: l.Link.Method,
				Name:   l.Link.Name,
				Rel:    l.Link.Rel,
			},
		})
	}

	return c.JSON(http.StatusOK, reply)
}

// GetRequestSchema returns 'https://' for TLS based request or 'http://' otherwise
func GetRequestSchema(r *http.Request) string {
	if r.TLS != nil {
		return "https://"
	}
	return "http://"
}
