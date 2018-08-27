package integration

import (
	"context"
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
	AccessControlListSchemaID    = "access_control_list"
	ApplicationPolicySetSchemaID = "application_policy_set"
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
	chownPath      = "/chown"
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

	err := c.Login(context.Background())
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

// UpdateProject updates Project resource.
func (c *HTTPAPIClient) UpdateProject(t *testing.T, uuid string, requestData interface{}) {
	c.UpdateResource(t, path.Join(ProjectSingularPath, uuid), requestData)
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
	r, err := c.Create(context.Background(), path, requestData, &responseData)
	assert.NoError(
		t,
		err,
		fmt.Sprintf("creating resource failed\n requestData: %+v\n "+
			"response: %+v\n responseData: %+v", requestData, r, responseData),
	)
}

// GetResource gets resource.
func (c *HTTPAPIClient) GetResource(t *testing.T, path string, responseData interface{}) {
	r, err := c.Read(context.Background(), path, nil, &responseData)
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
	var responseData interface{}
	r, err := c.Delete(context.Background(), path, &responseData)
	assert.NoError(t, err, "deleting resource failed\n response: %+v\n responseData: %+v", r, responseData)
}

// CheckResourceDoesNotExist checks that there is no resource with given path.
func (c *HTTPAPIClient) CheckResourceDoesNotExist(t *testing.T, path string) {
	var responseData interface{}
	r, err := c.Do(context.Background(), echo.GET, path, nil, nil, &responseData, []int{http.StatusNotFound})
	assert.NoError(t, err, "getting resource failed\n response: %+v\n responseData: %+v", r, responseData)
}
