package keystone

import (
	"crypto/tls"
	"encoding/json"
	"net"
	"net/http"
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

// NewAuth creates new keystone auth
func (k *Client) NewAuth() *keystone.Auth {
	auth := keystone.New(k.AuthURL)
	auth.Client = k.httpClient
	return auth
}

func (k *Client) tokenRequest(method string, c echo.Context) (*http.Response, error) {
	tokenURL := k.AuthURL + "/auth/tokens"
	request, err := http.NewRequest(method, tokenURL, c.Request().Body)
	if err != nil {
		return nil, err
	}
	request.Header = c.Request().Header
	request.ContentLength = c.Request().ContentLength
	resp, err := k.httpClient.Do(request)

	return resp, err
}

// CreateToken sends token create request to keystone endpoint.
func (k *Client) CreateToken(c echo.Context) error {
	resp, err := k.tokenRequest(echo.POST, c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer resp.Body.Close() // nolint: errcheck
	c.Response().Header().Set("X-Subject-Token",
		resp.Header.Get("X-Subject-Token"))
	authResponse := &kscommon.AuthResponse{}
	_ = json.NewDecoder(resp.Body).Decode(authResponse) // nolint: errcheck

	return c.JSON(resp.StatusCode, authResponse)
}

// ValidateToken sends validate token request to keystone endpoint.
func (k *Client) ValidateToken(c echo.Context) error {
	resp, err := k.tokenRequest(echo.GET, c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer resp.Body.Close() // nolint: errcheck
	validateTokenResponse := &kscommon.ValidateTokenResponse{}
	_ = json.NewDecoder(resp.Body).Decode(validateTokenResponse) // nolint: errcheck

	return c.JSON(resp.StatusCode, validateTokenResponse)
}

// GetProjects sends project get request to keystone endpoint.
func (k *Client) GetProjects(c echo.Context) error {
	projectURL := k.AuthURL + strings.TrimPrefix(c.Request().URL.Path, authPrefix)
	request, err := http.NewRequest(echo.GET, projectURL, c.Request().Body)
	if err != nil {
		return err
	}
	request.Header = c.Request().Header
	resp, err := k.httpClient.Do(request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer resp.Body.Close() // nolint: errcheck
	projectsResponse := &ProjectListResponse{}
	_ = json.NewDecoder(resp.Body).Decode(projectsResponse) // nolint: errcheck

	return c.JSON(resp.StatusCode, projectsResponse)
}

// GetProject sends project get request to keystone endpoint.
func (k *Client) GetProject(c echo.Context, id string) error {
	urlParts := []string{
		k.AuthURL, strings.TrimPrefix(c.Request().URL.Path, authPrefix), id}
	projectURL := strings.Join(urlParts, pathSep)
	request, err := http.NewRequest(echo.GET, projectURL, c.Request().Body)
	if err != nil {
		return err
	}
	request.Header = c.Request().Header
	resp, err := k.httpClient.Do(request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer resp.Body.Close() // nolint: errcheck
	projectResponse := &ProjectResponse{}
	_ = json.NewDecoder(resp.Body).Decode(projectResponse) // nolint: errcheck

	return c.JSON(resp.StatusCode, projectResponse)
}
