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

// CreateRequiredNetworkIPAM creates NetworkIPAM and stops test execution on error.
func (c *HTTPAPIClient) CreateRequiredNetworkIPAM(t *testing.T, ni *models.NetworkIpam) {
	c.createRequired(t, NetworkIpamPluralPath, &services.CreateNetworkIpamRequest{NetworkIpam: ni})
}

// RemoveNetworkIPAM deletes NetworkIPAM resource.
func (c *HTTPAPIClient) RemoveNetworkIPAM(t *testing.T, uuid string) {
	c.deleteOptional(t, path.Join(NetworkIpamSingularPath, uuid))
}

// CreateRequiredProject creates Project and stops test execution on error.
func (c *HTTPAPIClient) CreateRequiredProject(t *testing.T, p *models.Project) {
	c.createRequired(t, ProjectPluralPath, &services.CreateProjectRequest{Project: p})
}

// UpdateRequiredProject updates Project and stops test execution on error.
func (c *HTTPAPIClient) UpdateRequiredProject(t *testing.T, uuid string, requestData interface{}) {
	c.updateRequired(t, path.Join(ProjectSingularPath, uuid), requestData)
}

// RemoveProject deletes Project resource.
func (c *HTTPAPIClient) RemoveProject(t *testing.T, uuid string) {
	c.deleteOptional(t, path.Join(ProjectSingularPath, uuid))
}

// CreateRequiredSecurityGroup creates SecurityGroup and stops test execution on error.
func (c *HTTPAPIClient) CreateRequiredSecurityGroup(t *testing.T, p *models.SecurityGroup) {
	c.createRequired(t, SecurityGroupPluralPath, &services.CreateSecurityGroupRequest{SecurityGroup: p})
}

// RemoveSecurityGroup deletes SecurityGroup resource.
func (c *HTTPAPIClient) RemoveSecurityGroup(t *testing.T, uuid string) {
	c.deleteOptional(t, path.Join(SecurityGroupSingularPath, uuid))
}

// CreateRequiredVirtualNetwork creates VirtualNetwork and stops test execution on error.
func (c *HTTPAPIClient) CreateRequiredVirtualNetwork(t *testing.T, vn *models.VirtualNetwork) {
	c.createRequired(t, VirtualNetworkPluralPath, &services.CreateVirtualNetworkRequest{VirtualNetwork: vn})
}

// FetchVirtualNetwork gets VirtualNetwork resource.
func (c *HTTPAPIClient) FetchVirtualNetwork(t *testing.T, uuid string) *models.VirtualNetwork {
	var responseData services.GetVirtualNetworkResponse
	c.fetchOptional(t, path.Join(VirtualNetworkSingularPath, uuid), &responseData)
	return responseData.VirtualNetwork
}

// RemoveVirtualNetwork deletes VirtualNetwork resource.
func (c *HTTPAPIClient) RemoveVirtualNetwork(t *testing.T, uuid string) {
	c.deleteOptional(t, path.Join(VirtualNetworkSingularPath, uuid))
}

func (c *HTTPAPIClient) createRequired(t *testing.T, path string, requestData interface{}) {
	var responseData interface{}
	r, err := c.Create(context.Background(), path, requestData, &responseData)
	require.NoError(
		t,
		err,
		fmt.Sprintf("creating resource failed\n requestData: %+v\n "+
			"response: %+v\n responseData: %+v", requestData, r, responseData),
	)
}

func (c *HTTPAPIClient) fetchOptional(t *testing.T, path string, responseData interface{}) {
	r, err := c.Read(context.Background(), path, &responseData)
	assert.NoError(
		t,
		err,
		fmt.Sprintf("getting resource failed\n response: %+v\n responseData: %+v", r, responseData),
	)
}

func (c *HTTPAPIClient) updateRequired(t *testing.T, path string, requestData interface{}) {
	var responseData interface{}
	r, err := c.Update(context.Background(), path, requestData, &responseData)
	require.NoError(
		t,
		err,
		fmt.Sprintf("updating resource failed\n requestData: %+v\n "+
			"response: %+v\n responseData: %+v", requestData, r, responseData),
	)
}

func (c *HTTPAPIClient) deleteOptional(t *testing.T, path string) {
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
