package discovery

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/fileutil"
)

type linkDetails struct {
	Path   string  `json:"href"`
	Method *string `json:"method"`
	Name   string  `json:"name"`
	Rel    string  `json:"rel"`
}

type link struct {
	Link linkDetails `json:"link"`
}

// Handler which serves a set of registered links.
type Handler interface {
	// Register adds a new link to the Handler.
	Register(path string, method string, name string, rel string)
	// Handle requests to return the links.
	Handle(c echo.Context) error
}

type handler struct {
	links []*link
}

// NewHandler creates a new Handler.
func NewHandler() Handler {
	return &handler{}
}

// Register adds a new link to the Handler.
func (h *handler) Register(path string, method string, name string, rel string) {
	// path is assumed to be relative with respect to addr
	path = strings.TrimPrefix(path, "/")

	var m *string
	if method != "" {
		m = &method
	}

	var l link
	l.Link.Path = path
	l.Link.Method = m
	l.Link.Name = name
	l.Link.Rel = rel
	h.links = append(h.links, &l)
}

// Handle requests to return the links.
func (h *handler) Handle(c echo.Context) error {
	r := c.Request()
	addr := fileutil.GetRequestSchema(r) + r.Host

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
