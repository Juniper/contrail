package apisrv

import (
	"net/http"
	"strings"

	"github.com/Juniper/contrail/pkg/services"
	"github.com/labstack/echo"
)

// Handler which serves a set of registered links.
type Handler struct {
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

// NewHandler creates a new Handler.
func NewHandler() *Handler {
	return &Handler{}
}

// Register adds a new link to the Handler.
func (h *Handler) Register(path string, method string, name string, rel string) {
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
func (h *Handler) Handle(c echo.Context) error {
	r := c.Request()
	addr := services.GetRequestSchema(r) + r.Host

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
