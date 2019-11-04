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
func (k *Client) CreateToken(ctx echo.Context, authURLs []string) error {
	return k.proxyRequestToRemoteKeystone(ctx, authURLs)
}

// ValidateToken sends validate token request to remote Keystone.
func (k *Client) ValidateToken(ctx echo.Context, authURLs []string) error {
	return k.proxyRequestToRemoteKeystone(ctx, authURLs)
}

// GetDomains sends domain get request to remote Keystone.
func (k *Client) GetDomains(ctx echo.Context, authURLs []string) error {
	return k.proxyRequestToRemoteKeystone(ctx, authURLs)
}

// GetProjects sends project get request to remote Keystone.
func (k *Client) GetProjects(ctx echo.Context, authURLs []string) error {
	return k.proxyRequestToRemoteKeystone(ctx, authURLs)
}

// GetProject sends project get request to remote Keystone.
func (k *Client) GetProject(ctx echo.Context, authURLs []string, id string) error {
	ctx.Request().URL.Path = path.Join(ctx.Request().URL.Path, id)
	return k.proxyRequestToRemoteKeystone(ctx, authURLs)
}

func (k *Client) proxyRequestToRemoteKeystone(ctx echo.Context, authURLs []string) error {
	rp, err := proxy.NewReverseProxy(authURLs)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("new reverse proxy: %v", err))
	}

	r := ctx.Request()
	r.URL.Path = strings.TrimPrefix(r.URL.Path, LocalAuthPath)

	rp.ServeHTTP(ctx.Response(), r)
	return nil
}
