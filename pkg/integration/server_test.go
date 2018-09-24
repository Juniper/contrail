package integration

import (
	"context"
	"crypto/tls"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/twinj/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/types"
)

func TestKVStore(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, "TestKVStore")
	RunTest(t, "./test_data/test_kv_store.yml")
}

func TestFloatingIP(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, "TestAPIServer")
	RunTest(t, "./test_data/test_floating_ip.yml")
}

func TestNetworkIpam(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, "TestAPIServer")
	RunTest(t, "./test_data/test_network_ipam.yml")
}

func TestVirtualNetwork(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, "TestAPIServer")
	RunTest(t, "./test_data/test_virtual_network.yml")
}

func TestSecurityGroup(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, t.Name())
	RunTest(t, "./test_data/test_security_group.yml")
}

func TestQuotaChecking(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, "TestAPIServer")
	RunTest(t, "./test_data/test_quota_checking.yml")
}

func TestSync(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, "TestSync")
	RunTest(t, "./test_data/test_sync.yml")
}

func TestValidation(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, "TestValidation")
	RunTest(t, "./test_data/test_validation.yml")
}

func TestNameUnique(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, t.Name())
	RunTest(t, "./test_data/test_name_unique.yml")
}

func TestBaseProperties(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, t.Name())
	RunTest(t, "./test_data/test_base_properties.yml")
}

func TestBasePropsTwoParents(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, t.Name())
	RunTest(t, "./test_data/test_base_props_two_parents.yml")
}

func TestBaseWithConfigRootInParents(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, t.Name())
	RunTest(t, "./test_data/test_base_config_root_parent.yml")
}

func TestProject(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, t.Name())
	RunTest(t, "./test_data/test_project.yml")
}

func TestProjectConflict(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, t.Name())
	RunTest(t, "./test_data/test_project_conflict.yml")
}

func TestInstanceIP(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, t.Name())
	RunTest(t, "./test_data/test_instance_ip.yml")
}

func TestVirtualMachineInterface(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, t.Name())
	RunTest(t, "./test_data/test_virtual_machine_interface.yml")
}

func TestLogicalRouter(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, t.Name())
	RunTest(t, "./test_data/test_logical_router.yml")
}

func TestVirtualRouter(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, t.Name())
	RunTest(t, "./test_data/test_virtual_router.yml")
}

func TestFQNameToID(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, t.Name())
	RunTest(t, "./test_data/test_fqname_to_id.yml")
}

func TestRefRead(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, t.Name())
	RunTest(t, "./test_data/test_ref_read.yml")
}

func TestRefUpdate(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, t.Name())
	RunTest(t, "./test_data/test_ref_update.yml")
}

func TestRefRelaxForDelete(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, t.Name())
	RunTest(t, "./test_data/test_ref_relax.yml")
}

func TestPropCollectionUpdate(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, t.Name())
	RunTest(t, "./test_data/test_prop_collection_update.yml")
}

func TestSetTag(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, t.Name())
	RunTest(t, "./test_data/test_set_tag.yml")
}

func TestK8sInstanceIP(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, t.Name())
	RunTest(t, "./test_data/test_k8s_instance_ip_alloc.yml")
}

func TestSanitizing(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, t.Name())
	RunTest(t, "./test_data/test_sanitizing.yml")
}

func TestProvisioning(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, t.Name())
	RunTest(t, "./test_data/test_provisioning.yml")
}

func TestIntPool(t *testing.T) {
	AddKeystoneProjectAndUser(runningServer.apiServer, t.Name())
	RunTest(t, "./test_data/test_int_pool.yml")
}

func restLogin(ctx context.Context, t *testing.T) (authToken string) {
	restClient := client.NewHTTP(
		runningServer.testServer.URL,
		runningServer.testServer.URL+"/keystone/v3",
		"TestGRPC",
		"TestGRPC",
		true,
		client.GetKeystoneScope("", "default", "", "TestGRPC"),
	)
	restClient.InSecure = true
	restClient.Init()
	err := restClient.Login(ctx)
	require.NoError(t, err)
	return restClient.AuthToken
}

func TestGRPC(t *testing.T) {
	ctx := context.Background()
	AddKeystoneProjectAndUser(runningServer.apiServer, "TestGRPC")

	authToken := restLogin(ctx, t)

	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})
	dial := strings.TrimPrefix(runningServer.testServer.URL, "https://")
	conn, err := grpc.Dial(dial, grpc.WithTransportCredentials(creds))
	assert.NoError(t, err)
	defer LogFatalIfErr(conn.Close)
	md := metadata.Pairs("X-Auth-Token", authToken)
	ctx = metadata.NewOutgoingContext(ctx, md)
	// Contact the server and print out its response.
	c := services.NewContrailServiceClient(conn)
	assert.NoError(t, err)
	project := models.MakeProject()
	project.UUID = uuid.NewV4().String()
	project.FQName = []string{"default-domain", "project", project.UUID}
	project.ParentType = "domain"
	project.ParentUUID = "beefbeef-beef-beef-beef-beefbeef0002"
	project.ConfigurationVersion = 1
	_, err = c.CreateProject(ctx, &services.CreateProjectRequest{
		Project: project,
	})
	assert.NoError(t, err)
	response, err := c.ListProject(ctx, &services.ListProjectRequest{
		Spec: &baseservices.ListSpec{
			Limit: 1,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(response.Projects))

	getResponse, err := c.GetProject(ctx, &services.GetProjectRequest{
		ID: project.UUID,
	})
	assert.NoError(t, err)
	assert.NotNil(t, getResponse.Project)

	_, err = c.DeleteProject(ctx, &services.DeleteProjectRequest{
		ID: project.UUID,
	})
	assert.NoError(t, err)
}

func TestIPAMGRPC(t *testing.T) {
	ctx := context.Background()
	AddKeystoneProjectAndUser(runningServer.apiServer, "TestIPAMGRPC")
	authToken := restLogin(ctx, t)

	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})
	dial := strings.TrimPrefix(runningServer.testServer.URL, "https://")

	conn, err := grpc.Dial(dial, grpc.WithTransportCredentials(creds))
	assert.NoError(t, err)
	defer LogFatalIfErr(conn.Close)
	md := metadata.Pairs("X-Auth-Token", authToken)
	ctx = metadata.NewOutgoingContext(ctx, md)
	// Contact the server and print out its response.
	c := services.NewIPAMClient(conn)
	assert.NoError(t, err)

	allocateResp, err := c.AllocateInt(ctx, &services.AllocateIntRequest{Pool: types.VirtualNetworkIDPoolKey})
	assert.NoError(t, err)

	_, err = c.DeallocateInt(ctx, &services.DeallocateIntRequest{
		Pool:  types.VirtualNetworkIDPoolKey,
		Value: allocateResp.GetValue(),
	})
	assert.NoError(t, err)
}

func TestRESTClient(t *testing.T) {
	ctx := context.Background()
	testName := "TestRESTClient"
	AddKeystoneProjectAndUser(runningServer.apiServer, testName)
	restClient := client.NewHTTP(
		runningServer.testServer.URL,
		runningServer.testServer.URL+"/keystone/v3",
		testName,
		testName,
		true,
		client.GetKeystoneScope("", "default", "", testName),
	)
	restClient.InSecure = true
	restClient.Init()
	err := restClient.Login(ctx)
	// Contact the server and print out its response.
	assert.NoError(t, err)
	project := models.MakeProject()
	project.UUID = uuid.NewV4().String()
	project.FQName = []string{"default-domain", "project", project.UUID}
	project.ParentType = "domain"
	project.ParentUUID = "beefbeef-beef-beef-beef-beefbeef0002"
	project.ConfigurationVersion = 1
	_, err = restClient.CreateProject(ctx, &services.CreateProjectRequest{
		Project: project,
	})
	assert.NoError(t, err)
	response, err := restClient.ListProject(ctx, &services.ListProjectRequest{
		Spec: &baseservices.ListSpec{
			Limit: 1,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(response.Projects))

	getResponse, err := restClient.GetProject(ctx, &services.GetProjectRequest{
		ID: project.UUID,
	})
	assert.NoError(t, err)
	assert.NotNil(t, getResponse.Project)
	assert.Equal(t, project.UUID, getResponse.Project.UUID)

	_, err = restClient.DeleteProject(ctx, &services.DeleteProjectRequest{
		ID: project.UUID,
	})
	assert.NoError(t, err)
}
