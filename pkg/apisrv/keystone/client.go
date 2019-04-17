package keystone

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/databus23/keystone"
	"github.com/labstack/echo"

	kscommon "github.com/Juniper/contrail/pkg/keystone"
)

const (
	keystoneVersion = "v3"
	authPrefix      = "/keystone/" + keystoneVersion
	pathSep         = "/"
)

// Client represents a client.
type Client struct {
	AuthURL      string `yaml:"authurl"`
	LocalAuthURL string `yaml:"local_authurl"`
	httpClient   *http.Client
	InSecure     bool `yaml:"insecure"`
}

// NewKeystoneClient makes keystone client.
func NewKeystoneClient(authURL string, insecure bool) *Client {
	c := &Client{
		AuthURL:      authURL,
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
	k.AuthURL = authURL + "/" + keystoneVersion
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
	auth := keystone.New(k.AuthURL)
	auth.Client = k.httpClient
	return auth
}

func (k *Client) tokenRequest(c echo.Context) error {
	r := c.Request()
	r.URL.Path = "/auth/tokens"
	tokenURL, err := url.Parse(k.AuthURL)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	server := httputil.NewSingleHostReverseProxy(tokenURL)
	server.ServeHTTP(c.Response(), r)
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

// GetProjects sends project get request to keystone endpoint.
func (k *Client) GetProjects(c echo.Context) error {
	r := c.Request()
	r.URL.Path = strings.TrimPrefix(r.URL.Path, authPrefix)
	projectURL, err := url.Parse(k.AuthURL)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	server := httputil.NewSingleHostReverseProxy(projectURL)
	server.ServeHTTP(c.Response(), r)
	return nil
}

// GetProject sends project get request to keystone endpoint.
func (k *Client) GetProject(c echo.Context, id string) error {
	r := c.Request()
	urlParts := []string{
		strings.TrimPrefix(r.URL.Path, authPrefix), id}
	r.URL.Path = strings.Join(urlParts, pathSep)
	projectURL, err := url.Parse(k.AuthURL)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	server := httputil.NewSingleHostReverseProxy(projectURL)
	server.ServeHTTP(c.Response(), r)
	return nil
}
