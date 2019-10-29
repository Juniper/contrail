package keystone

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/Juniper/contrail/pkg/proxy"
	"github.com/labstack/echo"
)

// Client represents a Keystone client. It proxies requests to given auth URLs.
type Client struct{}

// CreateToken sends token create request to remote Keystone.
func (k *Client) CreateToken(c echo.Context, authURLs []string) error {
	return k.proxyRequestToRemoteKeystone(c, authURLs)
}

// ValidateToken sends validate token request to remote Keystone.
func (k *Client) ValidateToken(c echo.Context, authURLs []string) error {
	return k.proxyRequestToRemoteKeystone(c, authURLs)
}

// GetDomains sends domain get request to remote Keystone.
func (k *Client) GetDomains(c echo.Context, authURLs []string) error {
	return k.proxyRequestToRemoteKeystone(c, authURLs)
}

// GetProjects sends project get request to remote Keystone.
func (k *Client) GetProjects(c echo.Context, authURLs []string) error {
	return k.proxyRequestToRemoteKeystone(c, authURLs)
}

// GetProject sends project get request to remote Keystone.
func (k *Client) GetProject(c echo.Context, authURLs []string, id string) error {
	c.Request().URL.Path = path.Join(c.Request().URL.Path, id)
	return k.proxyRequestToRemoteKeystone(c, authURLs)
}

func (k *Client) proxyRequestToRemoteKeystone(c echo.Context, authURLs []string) error {
	rp, err := proxy.NewReverseProxy(authURLs)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("new reverse proxy: %v", err))
	}

	r := c.Request()
	r.URL.Path = strings.TrimPrefix(r.URL.Path, LocalAuthPath)

	rp.ServeHTTP(c.Response(), r)
	return nil
}
