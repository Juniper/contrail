package apisrv

import (
	"context"
	"crypto/tls"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twinj/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

func TestKVStore(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, "TestKVStore")
	RunTest(t, "./test_data/test_kv_store.yml")
}

func TestFloatingIP(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, "TestAPIServer")
	RunTest(t, "./test_data/test_floating_ip.yml")
}

func TestNetworkIpam(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, "TestAPIServer")
	RunTest(t, "./test_data/test_network_ipam.yml")
}

func TestVirtualNetwork(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, "TestAPIServer")
	RunTest(t, "./test_data/test_virtual_network.yml")
}

func TestSecurityGroup(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, t.Name())
	RunTest(t, "./test_data/test_security_group.yml")
}

func TestQuotaChecking(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, "TestAPIServer")
	RunTest(t, "./test_data/test_quota_checking.yml")
}

func TestSync(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, "TestSync")
	RunTest(t, "./test_data/test_sync.yml")
}

func TestValidation(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, "TestValidation")
	RunTest(t, "./test_data/test_validation.yml")
}

func TestNameUnique(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, t.Name())
	RunTest(t, "./test_data/test_name_unique.yml")
}

func TestBaseProperties(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, t.Name())
	RunTest(t, "./test_data/test_base_properties.yml")
}

func TestBasePropsTwoParents(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, t.Name())
	RunTest(t, "./test_data/test_base_props_two_parents.yml")
}

func TestBaseWithConfigRootInParents(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, t.Name())
	RunTest(t, "./test_data/test_base_config_root_parent.yml")
}

func TestProject(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, t.Name())
	RunTest(t, "./test_data/test_project.yml")
}

func TestProjectConflict(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, t.Name())
	RunTest(t, "./test_data/test_project_conflict.yml")
}

func TestInstanceIP(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, t.Name())
	RunTest(t, "./test_data/test_instance_ip.yml")
}

func TestVirtualMachineInterface(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, t.Name())
	RunTest(t, "./test_data/test_virtual_machine_interface.yml")
}

func TestLogicalRouter(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, t.Name())
	RunTest(t, "./test_data/test_logical_router.yml")
}

func TestFQNameToID(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, t.Name())
	RunTest(t, "./test_data/test_fqname_to_id.yml")
}

func TestRefRead(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, t.Name())
	RunTest(t, "./test_data/test_ref_read.yml")
}

func TestRefUpdate(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, t.Name())
	RunTest(t, "./test_data/test_ref_update.yml")
}

func TestRefRelaxForDelete(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, t.Name())
	RunTest(t, "./test_data/test_ref_relax.yml")
}

func TestPropCollectionUpdate(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, t.Name())
	RunTest(t, "./test_data/test_prop_collection_update.yml")
}

func TestSetTag(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, t.Name())
	RunTest(t, "./test_data/test_set_tag.yml")
}

func TestK8sInstanceIP(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, t.Name())
	RunTest(t, "./test_data/test_k8s_instance_ip_alloc.yml")
}

func TestSanitizing(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, t.Name())
	RunTest(t, "./test_data/test_sanitizing.yml")
}

func TestProvisioning(t *testing.T) {
	AddKeystoneProjectAndUser(APIServer, t.Name())
	RunTest(t, "./test_data/test_provisioning.yml")
}

func TestGRPC(t *testing.T) {
	ctx := context.Background()
	AddKeystoneProjectAndUser(APIServer, "TestGRPC")
	restClient := client.NewHTTP(
		TestServer.URL,
		TestServer.URL+"/keystone/v3",
		"TestGRPC",
		"TestGRPC",
		"default",
		true,
		&keystone.Scope{
			Project: &keystone.Project{
				Name: "TestGRPC",
			},
		},
	)
	restClient.InSecure = true
	restClient.Init()
	err := restClient.Login(ctx)
	assert.NoError(t, err)
	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})
	dial := strings.TrimPrefix(TestServer.URL, "https://")
	conn, err := grpc.Dial(dial, grpc.WithTransportCredentials(creds))
	assert.NoError(t, err)
	defer LogFatalIfErr(conn.Close)
	md := metadata.Pairs("X-Auth-Token", restClient.AuthToken)
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

func TestRESTClient(t *testing.T) {
	ctx := context.Background()
	testName := "TestRESTClient"
	AddKeystoneProjectAndUser(APIServer, testName)
	restClient := client.NewHTTP(
		TestServer.URL,
		TestServer.URL+"/keystone/v3",
		testName,
		testName,
		"default",
		true,
		&keystone.Scope{
			Project: &keystone.Project{
				Name: testName,
			},
		},
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
