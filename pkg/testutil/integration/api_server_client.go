package integration

import (
	"context"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/services"
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
func NewTestingHTTPClient(t *testing.T, apiServerURL string) *HTTPAPIClient {
	l := log.NewLogger("http-api-client")
	l.WithFields(logrus.Fields{"endpoint": apiServerURL}).Debug("Connecting to API Server")

	c, err := NewHTTPClient(apiServerURL)
	require.NoError(t, err, "connecting to API Server failed")

	return &HTTPAPIClient{
		HTTP: c,
		log:  l,
	}
}

// NewHTTPClient creates HTTP client of API Server using default testing configuration.
func NewHTTPClient(apiServerURL string) (*client.HTTP, error) {
	c := client.NewHTTP(
		apiServerURL,
		apiServerURL+authEndpointSuffix,
		AdminUserID,
		AdminUserPassword,
		true,
		client.GetKeystoneScope(DefaultDomainID, "",
			AdminProjectID, AdminProjectName),
	)
	c.Debug = true

	return c, c.Login(context.Background())
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
