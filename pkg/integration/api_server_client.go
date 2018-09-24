package integration

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/apisrv/client"
	pkglog "github.com/Juniper/contrail/pkg/log"
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
	l := pkglog.NewLogger("http-api-client")
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
	var responseData apisrv.FQNameToIDResponse
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

// CreateResource creates resource.
func (c *HTTPAPIClient) CreateResource(t *testing.T, path string, requestData interface{}) {
	var responseData interface{}
	r, err := c.Create(context.Background(), path, requestData, &responseData)
	require.NoError(
		t,
		err,
		fmt.Sprintf("creating resource failed\n requestData: %+v\n "+
			"response: %+v\n responseData: %+v", requestData, r, responseData),
	)
}

// GetResource gets resource.
func (c *HTTPAPIClient) GetResource(t *testing.T, path string, responseData interface{}) {
	r, err := c.Read(context.Background(), path, &responseData)
	assert.NoError(
		t,
		err,
		fmt.Sprintf("getting resource failed\n response: %+v\n responseData: %+v", r, responseData),
	)
}

// UpdateResource updates resource.
func (c *HTTPAPIClient) UpdateResource(t *testing.T, path string, requestData interface{}) {
	var responseData interface{}
	r, err := c.Update(context.Background(), path, requestData, &responseData)
	assert.NoError(
		t,
		err,
		fmt.Sprintf("updating resource failed\n requestData: %+v\n "+
			"response: %+v\n responseData: %+v", requestData, r, responseData),
	)
}

// DeleteResource deletes resource.
func (c *HTTPAPIClient) DeleteResource(t *testing.T, path string) {
	c.log.Debug("deleting resource %v", path)
	var responseData interface{}
	r, err := c.Delete(context.Background(), path, &responseData)
	require.NoError(t, err, "deleting resource failed\n response: %+v\n responseData: %+v", r, responseData)
}

// CheckResourceDoesNotExist checks that there is no resource with given path.
func (c *HTTPAPIClient) CheckResourceDoesNotExist(t *testing.T, path string) {
	var responseData interface{}
	r, err := c.Do(context.Background(), echo.GET, path, nil, nil, &responseData, []int{http.StatusNotFound})
	assert.NoError(t, err, "getting resource failed\n response: %+v\n responseData: %+v", r, responseData)
}
