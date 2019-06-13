package integration

import (
	"context"
	"net/http"
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	apisrvkeystone "github.com/Juniper/contrail/pkg/apisrv/keystone"
)

// Resource constants
const (
	DomainType                 = "domain"
	DefaultDomainUUID          = "beefbeef-beef-beef-beef-beefbeef0002"
	NetworkIpamSingularPath    = "/network-ipam"
	NetworkIpamPluralPath      = "/network-ipams"
	ProjectType                = "project"
	ProjectSingularPath        = "/project"
	ProjectPluralPath          = "/projects"
	SecurityGroupSingularPath  = "/security-group"
	SecurityGroupPluralPath    = "/security-groups"
	VirtualNetworkSingularPath = "/virtual-network"
	VirtualNetworkPluralPath   = "/virtual-networks"
)

const (
	chownPath      = "/chown"
	fqNameToIDPath = "/fqname-to-id"
)

// HTTPAPIClient is API Server client for testing purposes.
type HTTPAPIClient struct {
	*client.HTTP
	log *logrus.Entry
}

// NewTestingHTTPClient creates HTTP client of API Server with testing capabilities.
// It logs in with given userID, such as "alice" and "bob".
func NewTestingHTTPClient(t *testing.T, apiServerURL string, userID string) *HTTPAPIClient {
	l := logutil.NewLogger("http-api-client")
	l.WithFields(logrus.Fields{"endpoint": apiServerURL}).Debug("Connecting to API Server")

	var c *client.HTTP
	var err error
	switch userID {
	case AdminUserID:
		c, err = NewAdminHTTPClient(apiServerURL)
	case BobUserID:
		c, err = NewHTTPClient(apiServerURL)
	default:
		require.FailNowf(t, "Invalid user ID: %v, only %v and %v are supported now", userID, AdminUserID, BobUserID)
	}
	require.NoError(t, err, "connecting to API Server failed")

	return &HTTPAPIClient{
		HTTP: c,
		log:  l,
	}
}

// NewAdminHTTPClient creates HTTP client of API Server using Alice user (admin) credentials.
func NewAdminHTTPClient(apiServerURL string) (*client.HTTP, error) {
	c := client.NewHTTP(&client.HTTPConfig{
		ID:       AdminUserID,
		Password: AdminUserPassword,
		Endpoint: apiServerURL,
		AuthURL:  apiServerURL + apisrvkeystone.AuthEndpointSuffix,
		Scope: keystone.NewScope(
			DefaultDomainID,
			DefaultDomainName,
			AdminProjectID,
			AdminProjectName,
		),
		Insecure: true,
	})

	_, err := c.Login(context.Background())
	return c, err
}

// NewHTTPClient creates HTTP client of API Server using Bob user credentials.
func NewHTTPClient(apiServerURL string) (*client.HTTP, error) {
	c := client.NewHTTP(&client.HTTPConfig{
		ID:       BobUserID,
		Password: BobUserPassword,
		Endpoint: apiServerURL,
		AuthURL:  apiServerURL + apisrvkeystone.AuthEndpointSuffix,
		Scope: keystone.NewScope(
			DefaultDomainID,
			DefaultDomainName,
			DemoProjectID,
			DemoProjectName,
		),
		Insecure: true,
	})

	_, err := c.Login(context.Background())
	return c, err
}

type fqNameToIDLegacyRequest struct {
	FQName []string `json:"fq_name"`
	Type   string   `json:"type"`
}

// FQNameToID performs FQName to ID request.
func (c *HTTPAPIClient) FQNameToID(t *testing.T, fqName []string, resourceType string) (uuid string) {
	var responseData services.FQNameToIDResponse
	r, err := c.Do(
		context.Background(),
		echo.POST,
		fqNameToIDPath,
		nil,
		&fqNameToIDLegacyRequest{
			FQName: fqName,
			Type:   resourceType,
		},
		&responseData,
		[]int{http.StatusOK},
	)
	assert.NoError(t, err, "FQName to ID failed\n response: %+v\n responseData: %+v", r, responseData)

	return responseData.UUID
}

type chownRequest struct {
	Owner string `json:"owner"`
	UUID  string `json:"uuid"`
}

// Chown performs "chown" request.
func (c *HTTPAPIClient) Chown(t *testing.T, owner, uuid string) {
	var responseData interface{}
	r, err := c.Do(
		context.Background(),
		echo.POST,
		chownPath,
		nil,
		&chownRequest{
			Owner: owner,
			UUID:  uuid,
		},
		&responseData,
		[]int{http.StatusOK},
	)
	assert.NoError(t, err, "Chown failed\n response: %+v\n responseData: %+v", r, responseData)
}

// CheckResourceDoesNotExist checks that there is no resource with given path.
func (c *HTTPAPIClient) CheckResourceDoesNotExist(t *testing.T, path string) {
	var responseData interface{}
	r, err := c.Do(context.Background(), echo.GET, path, nil, nil, &responseData, []int{http.StatusNotFound})
	assert.NoError(t, err, "getting resource failed\n response: %+v\n responseData: %+v", r, responseData)
}
