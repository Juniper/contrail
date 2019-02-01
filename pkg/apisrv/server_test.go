package apisrv_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/twinj/uuid"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

func TestKVStore(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_kv_store.yml")
}

func TestBGPAsAService(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_bgpaas.yml")
}

func TestFloatingIP(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_floating_ip.yml")
}

func TestForwardingClass(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, "TestAPIServer")
	RunTest(t, "./test_data/test_forwarding_class.yml")
}

func TestNetworkIpam(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_network_ipam.yml")
}

func TestNetworkPolicy(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_network_policy.yml")
}

func TestVirtualNetwork(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_virtual_network.yml")
}

func TestVirtualNetworkMultiChain(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_virtual_network_multi_chain.yml")
}

func TestSecurityGroup(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_security_group.yml")
}

func TestQoSConfig(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_qos_config.yml")
}

func TestQuotaChecking(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_quota_checking.yml")
}

func TestSync(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_sync.yml")
}

func TestTagType(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_tag_type.yml")
}

func TestValidation(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_validation.yml")
}

func TestNameUnique(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_name_unique.yml")
}

func TestBaseProperties(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_base_properties.yml")
}

func TestBasePropsTwoParents(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_base_props_two_parents.yml")
}

func TestBaseWithConfigRootInParents(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_base_config_root_parent.yml")
}

func TestProject(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_project.yml")
}

func TestProjectConflict(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_project_conflict.yml")
}

func TestInstanceIP(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_instance_ip.yml")
}

func TestVirtualMachineInterface(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_virtual_machine_interface.yml")
}

func TestLogicalRouter(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_logical_router.yml")
}

func TestVirtualRouter(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_virtual_router.yml")
}

func TestFirewallPolicy(t *testing.T) {
	RunTest(t, "./test_data/test_firewall_policy.yml")
}

func TestFirewallRule(t *testing.T) {
	RunTest(t, "./test_data/test_firewall_rule.yml")
}

func TestFQNameToID(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_fqname_to_id.yml")
}

func TestRefRead(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_ref_read.yml")
}

func TestRefUpdate(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_ref_update.yml")
}

func TestRefRelaxForDelete(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_ref_relax.yml")
}

func TestRefRelaxForDeleteInvalidInput(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_ref_relax_invalid_input.yml")
}

func TestPropCollectionUpdate(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_prop_collection_update.yml")
}

func TestServiceTemplate(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_service_template.yml")
}

func TestTag(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_tag.yml")
}

func TestSetTag(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_set_tag.yml")
}

func TestK8sInstanceIP(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_k8s_instance_ip_alloc.yml")
}

func TestSanitizing(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_sanitizing.yml")
}

func TestProvisioning(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_provisioning.yml")
}

func TestIntPool(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_int_pool.yml")
}

func TestDerivedRelations(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_derived_relations.yml")
}

func TestAlarm(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_alarm.yml")
}

func TestAliasIP(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_alias_ip.yml")
}

func TestPhysicalInterface(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_physical_interface.yml")
}

func TestLogicalInterface(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_logical_interface.yml")
}

func TestDomain(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_domain.yml")
}

func TestQoS(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_qos.yml")
}

func TestIsVisible(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_user_visible.yml")
}

func restLogin(ctx context.Context, t *testing.T) (authToken string) {
	restClient := client.NewHTTP(
		server.URL(),
		server.URL()+"/keystone/v3",
		t.Name(),
		t.Name(),
		true,
		client.GetKeystoneScope("", "default", "", t.Name()),
	)
	restClient.InSecure = true
	restClient.Init()
	err := restClient.Login(ctx)
	require.NoError(t, err)
	return restClient.AuthToken
}

func TestRESTClient(t *testing.T) {
	ctx := context.Background()
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	restClient := client.NewHTTP(
		server.URL(),
		server.URL()+"/keystone/v3",
		t.Name(),
		t.Name(),
		true,
		client.GetKeystoneScope("", "default", "", t.Name()),
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
	project.ParentUUID = integration.DefaultDomainUUID
	project.ConfigurationVersion = 1
	project.IDPerms = &models.IdPermsType{UserVisible: true}
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

func TestPagination(t *testing.T) {
	context := map[string]interface{}{
		"ids": []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		"lists": []struct {
			name        string
			marker      int
			limit       int
			expectedIds []int
		}{
			{
				name:        "show limited count of alarms",
				limit:       3,
				expectedIds: []int{0, 1, 2},
			},
			{
				name:        "show limited count of alarms starting form the marker",
				marker:      2,
				limit:       4,
				expectedIds: []int{3, 4, 5, 6},
			},
			{
				name:        "show the alarms starting from the marker",
				marker:      7,
				expectedIds: []int{8, 9},
			},
			{
				name:        "check if no alarms arter the last marker",
				marker:      9,
				expectedIds: []int{},
			},
		},
	}

	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTestTemplate(t, "./test_data/test_pagination.tmpl", context)
}
