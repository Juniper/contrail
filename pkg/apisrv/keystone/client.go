package keystone

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"github.com/databus23/keystone"
	"github.com/labstack/echo"

	apicommon "github.com/Juniper/contrail/pkg/apisrv/common"
	kscommon "github.com/Juniper/contrail/pkg/keystone"
)

const (
	keystoneVersion = "v3"
	authPrefix      = "/keystone/" + keystoneVersion
	pathSep         = "/"
)

// Client represents a client.
type Client struct {
	AuthEndpoints []*apicommon.Endpoint `yaml:"auth_endpoints"`
	LocalAuthURL  string                `yaml:"local_authurl"`
	httpClient    *http.Client
	InSecure      bool `yaml:"insecure"`
}

// NewKeystoneClient makes keystone client.
func NewKeystoneClient(authURL string, insecure bool) *Client {
	c := &Client{
		LocalAuthURL: authURL,
		InSecure:     insecure,
	}
	c.Init()
	return c
}

// Init is used to initialize a keystone client.
func (k *Client) Init() {
	tr := &http.Transport{
		Dial:            (&net.Dialer{}).Dial,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: k.InSecure},
	}
	client := &http.Client{
		Transport: tr,
	}
	k.httpClient = client
}

// SetAuthEndpoint uses specified auth url in the keystone auth.
func (k *Client) SetAuthEndpoint(authEndpoints []*apicommon.Endpoint) {
	for _, endpoint := range authEndpoints {
		endpoint.URL = endpoint.URL + "/" + keystoneVersion
	}
	k.AuthEndpoints = authEndpoints
}

// SetAuthIdentity uses specified auth creds in the keystone auth request.
func (k *Client) SetAuthIdentity(
	c echo.Context, authRequest kscommon.AuthRequest) echo.Context {
	b, _ := json.Marshal(authRequest) // nolint: errcheck
	c.Request().Body = ioutil.NopCloser(bytes.NewReader(b))
	c.Request().ContentLength = int64(len(b))
	return c
}

// NewAuth creates new keystone auth
func (k *Client) NewAuth() *keystone.Auth {
	auth := keystone.New(k.LocalAuthURL)
	auth.Client = k.httpClient
	return auth
}

func (k *Client) tokenRequest(c echo.Context) error {
	r := c.Request()
	r.URL.Path = "/auth/tokens"
	servers := apicommon.NewReverseProxy(k.AuthEndpoints)
	servers.ServeHTTP(c.Response(), r)
	return nil
}

// CreateToken sends token create request to keystone endpoint.
func (k *Client) CreateToken(c echo.Context) error {
	return k.tokenRequest(c)
}

// ValidateToken sends validate token request to keystone endpoint.
func (k *Client) ValidateToken(c echo.Context) error {
	return k.tokenRequest(c)
}

// GetDomains sends domain get request to keystone endpoint.
func (k *Client) GetDomains(c echo.Context) error {
	r := c.Request()
	r.URL.Path = strings.TrimPrefix(r.URL.Path, authPrefix)
	servers := apicommon.NewReverseProxy(k.AuthEndpoints)
	servers.ServeHTTP(c.Response(), r)
	return nil
}

// GetProjects sends project get request to keystone endpoint.
func (k *Client) GetProjects(c echo.Context) error {
	r := c.Request()
	r.URL.Path = strings.TrimPrefix(r.URL.Path, authPrefix)
	servers := apicommon.NewReverseProxy(k.AuthEndpoints)
	servers.ServeHTTP(c.Response(), r)
	return nil
}

// GetProject sends project get request to keystone endpoint.
func (k *Client) GetProject(c echo.Context, id string) error {
	r := c.Request()
	urlParts := []string{
		strings.TrimPrefix(r.URL.Path, authPrefix), id}
	r.URL.Path = strings.Join(urlParts, pathSep)
	servers := apicommon.NewReverseProxy(k.AuthEndpoints)
	servers.ServeHTTP(c.Response(), r)
	return nil
}
