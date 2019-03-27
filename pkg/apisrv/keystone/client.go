package keystone

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"github.com/databus23/keystone"
	"github.com/labstack/echo"
	//"github.com/sirupsen/logrus"

	//auth2 "github.com/Juniper/contrail/pkg/auth"
	//"github.com/Juniper/contrail/pkg/errutil"
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

// Authenticate
func (k *Client) Authenticate(ctx context.Context, c echo.Context, tokenString string) (context.Context, error) {
	resp, err := k.tokenRequest(echo.GET, c)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	return nil, nil
	//validatedToken := resp
	/*
		roles := []string{}
		for _, r := range validatedToken.Roles {
			roles = append(roles, r.Name)
		}
		project := validatedToken.Project
		if project == nil {
			logrus.Debug("No project in a token")
			return nil, errutil.ErrorUnauthenticated
		}
		domain := validatedToken.Project.Domain.ID
		user := validatedToken.User

		objPerms := auth2.NewObjPerms(validatedToken)
		authContext := auth2.NewContext(domain, project.ID, user.ID, roles, tokenString, objPerms)

		var authKey interface{} = "auth"
		newCtx := context.WithValue(ctx, authKey, authContext)
		return newCtx, nil
	*/
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
