package integration

import (
	"fmt"
	"net/http"
	"path"
	"testing"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// Resource constants
const (
	DomainType                 = "domain"
	DefaultDomainUUID          = "beefbeef-beef-beef-beef-beefbeef0002"
	ProjectType                = "project"
	ProjectSchemaID            = "project"
	ProjectSingularPath        = "/project"
	ProjectPluralPath          = "/projects"
	NetworkIPAMSchemaID        = "network_ipam"
	NetworkIpamSingularPath    = "/network-ipam"
	NetworkIpamPluralPath      = "/network-ipams"
	VirtualNetworkSchemaID     = "virtual_network"
	VirtualNetworkSingularPath = "/virtual-network"
	VirtualNetworkPluralPath   = "/virtual-networks"
)

// HTTPAPIClient is API Server client for tests purposes.
type HTTPAPIClient struct {
	*apisrv.Client
	log *logrus.Entry
}

// NewHTTPAPIClient creates HTTP client of API Server.
func NewHTTPAPIClient(t *testing.T, apiServerURL string) *HTTPAPIClient {
	l := pkglog.NewLogger("http-api-client")
	l.WithFields(logrus.Fields{"endpoint": apiServerURL}).Debug("Connecting to API Server")
	c := apisrv.NewClient(
		apiServerURL,
		apiServerURL+authEndpointSuffix,
		AdminUserID,
		AdminUserPassword,
		DefaultDomainID,
		true,
		&keystone.Scope{
			Project: &keystone.Project{
				ID:   AdminProjectID,
				Name: AdminProjectName,
				Domain: &keystone.Domain{
					ID: DefaultDomainID,
				},
			},
		},
	)
	c.Debug = true

	err := c.Login()
	require.NoError(t, err, "connecting API Server failed")

	return &HTTPAPIClient{
		Client: c,
		log:    l,
	}
}

// CreateProject creates Project resource.
func (c *HTTPAPIClient) CreateProject(t *testing.T, p *models.Project) {
	c.CreateResource(t, ProjectPluralPath, &services.CreateProjectRequest{Project: p})
}

// DeleteProject deletes Project resource.
func (c *HTTPAPIClient) DeleteProject(t *testing.T, uuid string) {
	c.DeleteResource(t, path.Join(ProjectSingularPath, uuid))
}

// CreateNetworkIPAM creates NetworkIPAM resource.
func (c *HTTPAPIClient) CreateNetworkIPAM(t *testing.T, ni *models.NetworkIpam) {
	c.CreateResource(t, NetworkIpamPluralPath, &services.CreateNetworkIpamRequest{NetworkIpam: ni})
}

// DeleteNetworkIPAM deletes NetworkIPAM resource.
func (c *HTTPAPIClient) DeleteNetworkIPAM(t *testing.T, uuid string) {
	c.DeleteResource(t, path.Join(NetworkIpamSingularPath, uuid))
}

// CreateVirtualNetwork creates VirtualNetwork resource.
func (c *HTTPAPIClient) CreateVirtualNetwork(t *testing.T, vn *models.VirtualNetwork) {
	c.CreateResource(t, VirtualNetworkPluralPath, &services.CreateVirtualNetworkRequest{VirtualNetwork: vn})
}

// GetVirtualNetwork gets VirtualNetwork resource.
func (c *HTTPAPIClient) GetVirtualNetwork(t *testing.T, uuid string) *models.VirtualNetwork {
	var responseData services.GetVirtualNetworkResponse
	c.GetResource(t, path.Join(VirtualNetworkSingularPath, uuid), &responseData)
	return responseData.VirtualNetwork
}

// DeleteVirtualNetwork deletes VirtualNetwork resource.
func (c *HTTPAPIClient) DeleteVirtualNetwork(t *testing.T, uuid string) {
	c.DeleteResource(t, path.Join(VirtualNetworkSingularPath, uuid))
}

// CreateResource creates resource.
func (c *HTTPAPIClient) CreateResource(t *testing.T, path string, requestData interface{}) {
	var responseData interface{}
	r, err := c.Create(
		path,
		requestData,
		&responseData,
	)
	c.log.WithFields(logrus.Fields{
		"requestData":  requestData,
		"response":     r,
		"responseData": responseData,
	}).Debug("Got Create response")
	assert.NoError(t, err, fmt.Sprintf("creating resource failed\n requestData: %v\n response: %+v", requestData, r))
}

// GetResource gets resource.
func (c *HTTPAPIClient) GetResource(t *testing.T, path string, responseData interface{}) {
	r, err := c.Read(path, &responseData)
	c.log.WithFields(logrus.Fields{
		"response":     r,
		"responseData": responseData,
	}).Debug("Got Get response")
	assert.NoError(t, err, fmt.Sprintf("getting resource failed\n response: %+v", r))
}

// DeleteResource deletes resource.
func (c *HTTPAPIClient) DeleteResource(t *testing.T, path string) {
	r, err := c.EnsureDeleted(path, nil)
	c.log.WithField("response", r).Debug("Got Delete response")
	assert.NoError(t, err, "deleting resource failed\n response: %+v", r)
}

// CheckResourceDoesNotExist checks that there is no resource with given path.
func (c *HTTPAPIClient) CheckResourceDoesNotExist(t *testing.T, path string) {
	r, err := c.Do(
		echo.GET,
		path,
		nil,
		nil,
		[]int{http.StatusNotFound},
	)
	assert.NoError(t, err, "getting resource failed\n response: %+v", r)
}
