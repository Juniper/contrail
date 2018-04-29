package keystone

import (
	"crypto/tls"
	"net"
	"net/http"

	"github.com/databus23/keystone"
	"github.com/labstack/echo"
)

const (
	keystoneVersion = "v3"
)

// KeystoneClient represents a client.
// nolint
type KeystoneClient struct {
	AuthURL      string `yaml:"authurl"`
	LocalAuthURL string `yaml:"local_authurl"`
	httpClient   *http.Client
	InSecure     bool `yaml:"insecure"`
}

// NewKeystoneClient makes keystone client.
func NewKeystoneClient(authURL string, insecure bool) *KeystoneClient {
	c := &KeystoneClient{
		AuthURL:      authURL,
		LocalAuthURL: authURL,
		InSecure:     insecure,
	}
	c.Init()
	return c
}

// Init is used to initialize a keystone client.
func (k *KeystoneClient) Init() {
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
func (k *KeystoneClient) SetAuthURL(authURL string) {
	k.AuthURL = authURL + "/" + keystoneVersion
}

// NewAuth creates new keystone auth
func (k *KeystoneClient) NewAuth() *keystone.Auth {
	auth := keystone.New(k.AuthURL)
	auth.Client = k.httpClient
	return auth
}

func (k *KeystoneClient) tokenRequest(method string, c echo.Context) error {
	tokenURL := k.AuthURL + "/auth/tokens"
	request, err := http.NewRequest(method, tokenURL, c.Request().Body)
	if err != nil {
		return err
	}
	request.Header = c.Request().Header
	resp, err := k.httpClient.Do(request)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}
	if method == echo.POST {
		c.Response().Header().Set("X-Subject-Token",
			resp.Header.Get("X-Subject-Token"))
	}
	return c.JSON(resp.StatusCode, resp.Body)
}

// CreateToken sends token create request to keystone endpoint.
func (k *KeystoneClient) CreateToken(c echo.Context) error {
	return k.tokenRequest(echo.POST, c)
}

// ValidateToken sends validate token request to keystone endpoint.
func (k *KeystoneClient) ValidateToken(c echo.Context) error {
	return k.tokenRequest(echo.GET, c)
}

// GetProjects sends project get request to keystone endpoint.
func (k *KeystoneClient) GetProjects(c echo.Context) error {
	projectURL := k.AuthURL + "/auth/projects"
	request, err := http.NewRequest(echo.GET, projectURL, c.Request().Body)
	if err != nil {
		return err
	}
	request.Header = c.Request().Header
	resp, err := k.httpClient.Do(request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(resp.StatusCode, resp.Body)
}
