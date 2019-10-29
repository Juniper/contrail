package keystone

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"github.com/Juniper/contrail/pkg/proxy"
	"github.com/databus23/keystone"
	"github.com/labstack/echo"

	kscommon "github.com/Juniper/contrail/pkg/keystone"
)

// Keystone client constants.
const (
	LocalKeystonePath = "/keystone/v3"

	keystoneVersion = "v3"
)

// Client represents a client.
type Client struct {
	AuthURLs     string `yaml:"authurl"`
	LocalAuthURL string `yaml:"local_authurl"`
	httpClient   *http.Client
	InSecure     bool `yaml:"insecure"`
}

// NewClient makes keystone client.
func NewClient(authURL string, insecure bool) *Client {
	c := &Client{
		AuthURLs:     authURL,
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

// SetAuthURL uses specified auth url in the keystone auth.
func (k *Client) SetAuthURL(authURL string) {
	k.AuthURLs = authURL + "/" + keystoneVersion
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
	auth := keystone.New(k.AuthURLs)
	auth.Client = k.httpClient
	return auth
}

// CreateToken sends token create request to keystone endpoint.
func (k *Client) CreateToken(c echo.Context) error {
	return k.tokenRequest(c)
}

// ValidateToken sends validate token request to keystone endpoint.
func (k *Client) ValidateToken(c echo.Context) error {
	return k.tokenRequest(c)
}

func (k *Client) tokenRequest(c echo.Context) error {
	rp, err := proxy.NewReverseProxy([]string{k.AuthURLs})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("new reverse proxy: %v", err))
	}

	r := c.Request()
	r.URL.Path = "/auth/tokens"

	rp.ServeHTTP(c.Response(), r)
	return nil
}

// GetDomains sends domain get request to keystone endpoint.
func (k *Client) GetDomains(c echo.Context) error {
	rp, err := proxy.NewReverseProxy([]string{k.AuthURLs})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("new reverse proxy: %v", err))
	}

	r := c.Request()
	r.URL.Path = strings.TrimPrefix(r.URL.Path, LocalKeystonePath)

	rp.ServeHTTP(c.Response(), r)
	return nil
}

// GetProjects sends project get request to keystone endpoint.
func (k *Client) GetProjects(c echo.Context) error {
	rp, err := proxy.NewReverseProxy([]string{k.AuthURLs})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("new reverse proxy: %v", err))
	}

	r := c.Request()
	r.URL.Path = strings.TrimPrefix(r.URL.Path, LocalKeystonePath)

	rp.ServeHTTP(c.Response(), r)
	return nil
}

// GetProject sends project get request to keystone endpoint.
func (k *Client) GetProject(c echo.Context, id string) error {
	rp, err := proxy.NewReverseProxy([]string{k.AuthURLs})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("new reverse proxy: %v", err))
	}

	r := c.Request()
	r.URL.Path = strings.Join(
		[]string{strings.TrimPrefix(r.URL.Path, LocalKeystonePath), id},
		"/",
	)

	rp.ServeHTTP(c.Response(), r)
	return nil
}
