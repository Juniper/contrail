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

// Handler which serves a set of registered links.
type Handler struct {
	addr  string
	links []*link
}

// NewHandler creates a new Handler.
func NewHandler(addr string) *Handler {
	return &Handler{
		addr: addr,
	}
}

// Register adds a new link to the Handler.
func (h *Handler) Register(href string, method string, name string, rel string) {
	// path is assumed to be relative with respect to addr
	var m *string
	if method != "" {
		m = &method
	}

	h.links = append(h.links, &link{
		Path:   path.Join(h.addr, href),
		Method: m,
		Name:   name,
		Rel:    rel,
	})
}

// Handle requests to return the links.
func (h *Handler) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, struct {
		Addr  string  `json:"href"`
		Links []*link `json:"links"`
	}{h.addr, h.links})
}
