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
	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// Resource constants
const (
	AccessControlListSchemaID    = "access-control-list"
	ApplicationPolicySetSchemaID = "application-policy-set"
	DomainType                   = "domain"
	DefaultDomainUUID            = "beefbeef-beef-beef-beef-beefbeef0002"
	NetworkIPAMSchemaID          = "network_ipam"
	NetworkIpamSingularPath      = "/network-ipam"
	NetworkIpamPluralPath        = "/network-ipams"
	ProjectType                  = "project"
	ProjectSchemaID              = "project"
	ProjectSingularPath          = "/project"
	ProjectPluralPath            = "/projects"
	SecurityGroupSchemaID        = "security_group"
	SecurityGroupSingularPath    = "/security-group"
	SecurityGroupPluralPath      = "/security-groups"
	VirtualNetworkSchemaID       = "virtual_network"
	VirtualNetworkSingularPath   = "/virtual-network"
	VirtualNetworkPluralPath     = "/virtual-networks"
)

const (
	fqNameToIDPath = "/fqname-to-id"
)

// HTTPAPIClient is API Server client for tests purposes.
type HTTPAPIClient struct {
	*client.HTTP
}

// NewHTTPAPIClient creates HTTP client of API Server.
func NewHTTPAPIClient(t *testing.T, apiServerURL string) *HTTPAPIClient {
	l := pkglog.NewLogger("http-api-client")
	l.WithFields(logrus.Fields{"endpoint": apiServerURL}).Debug("Connecting to API Server")
	c := client.NewHTTP(
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
		HTTP: c,
	}
}

type fqNameToIDLegacyRequest struct {
	FQName []string `json:"fq_name"`
	Type   string   `json:"type"`
}

// FQNameToID performs FQName to ID request.
func (c *HTTPAPIClient) FQNameToID(t *testing.T, fqName []string, resourceType string) (uuid string) {
	var responseData apisrv.FqNameToIDResponse
	r, err := c.Do(
		echo.POST,
		fqNameToIDPath,
		&fqNameToIDLegacyRequest{
			FQName: fqName,
			Type:   resourceType,
		},
		&responseData,
		[]int{http.StatusOK},
	)
	assert.NoError(t, err, "getting resource failed\n response: %+v\n responseData: %+v", r, responseData)

	return responseData.UUID
}

// CreateNetworkIPAM creates NetworkIPAM resource.
func (c *HTTPAPIClient) CreateNetworkIPAM(t *testing.T, ni *models.NetworkIpam) {
	c.CreateResource(t, NetworkIpamPluralPath, &services.CreateNetworkIpamRequest{NetworkIpam: ni})
}

// DeleteNetworkIPAM deletes NetworkIPAM resource.
func (c *HTTPAPIClient) DeleteNetworkIPAM(t *testing.T, uuid string) {
	c.DeleteResource(t, path.Join(NetworkIpamSingularPath, uuid))
}

// CreateProject creates Project resource.
func (c *HTTPAPIClient) CreateProject(t *testing.T, p *models.Project) {
	c.CreateResource(t, ProjectPluralPath, &services.CreateProjectRequest{Project: p})
}

// DeleteProject deletes Project resource.
func (c *HTTPAPIClient) DeleteProject(t *testing.T, uuid string) {
	c.DeleteResource(t, path.Join(ProjectSingularPath, uuid))
}

// CreateSecurityGroup creates SecurityGroup resource.
func (c *HTTPAPIClient) CreateSecurityGroup(t *testing.T, p *models.SecurityGroup) {
	c.CreateResource(t, SecurityGroupPluralPath, &services.CreateSecurityGroupRequest{SecurityGroup: p})
}

// DeleteSecurityGroup deletes SecurityGroup resource.
func (c *HTTPAPIClient) DeleteSecurityGroup(t *testing.T, uuid string) {
	c.DeleteResource(t, path.Join(SecurityGroupSingularPath, uuid))
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
	assert.NoError(
		t,
		err,
		fmt.Sprintf("creating resource failed\n requestData: %+v\n "+
			"response: %+v\n responseData: %+v", requestData, r, responseData),
	)
}

// GetResource gets resource.
func (c *HTTPAPIClient) GetResource(t *testing.T, path string, responseData interface{}) {
	r, err := c.Read(path, &responseData)
	assert.NoError(
		t,
		err,
		fmt.Sprintf("getting resource failed\n response: %+v\n responseData: %+v", r, responseData),
	)
}

// DeleteResource deletes resource.
func (c *HTTPAPIClient) DeleteResource(t *testing.T, path string) {
	var responseData interface{}
	r, err := c.Delete(path, &responseData)
	assert.NoError(t, err, "deleting resource failed\n response: %+v\n responseData: %+v", r, responseData)
}

// CheckResourceDoesNotExist checks that there is no resource with given path.
func (c *HTTPAPIClient) CheckResourceDoesNotExist(t *testing.T, path string) {
	var responseData interface{}
	r, err := c.Do(
		echo.GET,
		path,
		nil,
		&responseData,
		[]int{http.StatusNotFound},
	)
	assert.NoError(t, err, "getting resource failed\n response: %+v\n responseData: %+v", r, responseData)
}
