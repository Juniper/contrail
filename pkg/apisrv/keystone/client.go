package keystone

import (
	"path"
	"strings"

	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/proxy"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// Client represents a Keystone client. It proxies requests to given auth URLs.
type Client struct {
	log *logrus.Entry
}

// NewClient returns a new Client.
func NewClient() *Client {
	return &Client{
		log: logutil.NewLogger("keystone-client"),
	}
}

// CreateToken sends token create request to remote Keystone.
func (c *Client) CreateToken(ctx echo.Context, authURLs []string) error {
	return c.proxyRequestToRemoteKeystone(ctx, authURLs)
}

// ValidateToken sends validate token request to remote Keystone.
func (c *Client) ValidateToken(ctx echo.Context, authURLs []string) error {
	return c.proxyRequestToRemoteKeystone(ctx, authURLs)
}

// GetDomains sends domain get request to remote Keystone.
func (c *Client) GetDomains(ctx echo.Context, authURLs []string) error {
	return c.proxyRequestToRemoteKeystone(ctx, authURLs)
}

// GetProjects sends project get request to remote Keystone.
func (c *Client) GetProjects(ctx echo.Context, authURLs []string) error {
	return c.proxyRequestToRemoteKeystone(ctx, authURLs)
}

// GetProject sends project get request to remote Keystone.
func (c *Client) GetProject(ctx echo.Context, authURLs []string, id string) error {
	ctx.Request().URL.Path = path.Join(ctx.Request().URL.Path, id)
	return c.proxyRequestToRemoteKeystone(ctx, authURLs)
}

func (c *Client) proxyRequestToRemoteKeystone(ctx echo.Context, authURLs []string) error {
	r := ctx.Request()
	r.URL.Path = strings.TrimPrefix(r.URL.Path, LocalAuthPath)
	return proxy.HandleRequest(ctx, authURLs, c.log)
}
