package apisrv_test

import (
	"context"
	"crypto/tls"
	"fmt"
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
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/Juniper/contrail/pkg/types"
)

func TestKVStore(t *testing.T) {
	integration.AddKeystoneProjectAndUser(server.APIServer, t.Name())
	RunTest(t, "./test_data/test_kv_store.yml")
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

func TestGRPC(t *testing.T) {
	testGRPCServer(t, t.Name(), func(ctx context.Context, conn *grpc.ClientConn) {
		c := services.NewContrailServiceClient(conn)
		project := models.Project{
			UUID:                 uuid.NewV4().String(),
			FQName:               []string{"default-domain", "project", "my-project"},
			ParentType:           "domain",
			ParentUUID:           integration.DefaultDomainUUID,
			ConfigurationVersion: 1,
		}
		_, err := c.CreateProject(ctx, &services.CreateProjectRequest{
			Project: &project,
		})
		assert.NoError(t, err)

		t.Run("create and delete namespace and project_namespace_ref", testNamespaceRef(ctx, c, project.UUID))
		t.Run("list and get project", testProjectRead(ctx, c, project.UUID))

		_, err = c.DeleteProject(ctx, &services.DeleteProjectRequest{
			ID: project.UUID,
		})
		assert.NoError(t, err)
	})
}
func testNamespaceRef(ctx context.Context, c services.ContrailServiceClient, projectUUID string) func(*testing.T) {
	return func(t *testing.T) {
		ns := models.Namespace{
			UUID:       uuid.NewV4().String(),
			ParentType: "domain",
			ParentUUID: integration.DefaultDomainUUID,
			Name:       "my-namespace",
		}
		_, err := c.CreateNamespace(ctx, &services.CreateNamespaceRequest{
			Namespace: &ns,
		})
		assert.NoError(t, err)
		_, err = c.CreateProjectNamespaceRef(ctx, &services.CreateProjectNamespaceRefRequest{
			ID:                  projectUUID,
			ProjectNamespaceRef: &models.ProjectNamespaceRef{UUID: ns.UUID},
		})
		assert.NoError(t, err)

		_, err = c.DeleteProjectNamespaceRef(ctx, &services.DeleteProjectNamespaceRefRequest{
			ID:                  projectUUID,
			ProjectNamespaceRef: &models.ProjectNamespaceRef{UUID: ns.UUID},
		})
		assert.NoError(t, err)

		_, err = c.DeleteNamespace(ctx, &services.DeleteNamespaceRequest{
			ID: ns.UUID,
		})
		assert.NoError(t, err)
	}
}

func testProjectRead(ctx context.Context, c services.ContrailServiceClient, projectUUID string) func(*testing.T) {
	return func(t *testing.T) {
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
			ID: projectUUID,
		})
		assert.NoError(t, err)
		assert.NotNil(t, getResponse.Project)
	}
}

func testGRPCServer(t *testing.T, testName string, testBody func(ctx context.Context, conn *grpc.ClientConn)) {
	ctx := context.Background()
	integration.AddKeystoneProjectAndUser(server.APIServer, testName)
	authToken := restLogin(ctx, t)

	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})
	dial := strings.TrimPrefix(server.URL(), "https://")

	conn, err := grpc.Dial(dial, grpc.WithTransportCredentials(creds))
	assert.NoError(t, err)
	defer testutil.LogFatalIfError(conn.Close)
	md := metadata.Pairs("X-Auth-Token", authToken)
	ctx = metadata.NewOutgoingContext(ctx, md)

	testBody(ctx, conn)
}

func TestFQNameToIDGRPC(t *testing.T) {
	testGRPCServer(t, t.Name(),
		func(ctx context.Context, conn *grpc.ClientConn) {
			c := services.NewFQNameToIDClient(conn)
			resp, err := c.FQNameToID(ctx, &services.FQNameToIDRequest{
				FQName: []string{"default-domain"},
				Type:   "domain",
			})
			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, integration.DefaultDomainUUID, resp.UUID)
		})
}

func TestChownGRPC(t *testing.T) {
	testGRPCServer(t, t.Name(),
		func(ctx context.Context, conn *grpc.ClientConn) {
			c := services.NewChownClient(conn)

			// This is only a stub implementation
			_, err := c.Chown(ctx, &services.ChownRequest{})
			assert.NoError(t, err)
		})
}

func TestIPAMGRPC(t *testing.T) {
	testGRPCServer(t, t.Name(),
		func(ctx context.Context, conn *grpc.ClientConn) {
			c := services.NewIPAMClient(conn)
			allocateResp, err := c.AllocateInt(ctx, &services.AllocateIntRequest{Pool: types.VirtualNetworkIDPoolKey})
			assert.NoError(t, err)

			_, err = c.DeallocateInt(ctx, &services.DeallocateIntRequest{
				Pool:  types.VirtualNetworkIDPoolKey,
				Value: allocateResp.GetValue(),
			})
			assert.NoError(t, err)
		})
}

func TestRefRelaxGRPC(t *testing.T) {
	testGRPCServer(t, t.Name(), func(ctx context.Context, conn *grpc.ClientConn) {
		c := client.NewGRPC(services.NewContrailServiceClient(conn))

		projectResponse, err := c.CreateProject(ctx, &services.CreateProjectRequest{
			Project: &models.Project{
				Name:       fmt.Sprintf("%s_project", t.Name()),
				ParentType: "domain",
				ParentUUID: integration.DefaultDomainUUID,
			},
		})
		require.NoError(t, err)
		project := projectResponse.GetProject()

		defer func() {
			_, err = c.DeleteProject(ctx, &services.DeleteProjectRequest{
				ID: project.GetUUID(),
			})
			assert.NoError(t, err)
		}()

		policyResponse, err := c.CreateNetworkPolicy(ctx, &services.CreateNetworkPolicyRequest{
			NetworkPolicy: &models.NetworkPolicy{
				ParentType: models.KindProject,
				ParentUUID: project.GetUUID(),
			},
		})
		require.NoError(t, err)
		policy := policyResponse.GetNetworkPolicy()

		defer func() {
			if policy == nil {
				return
			}
			_, err = c.DeleteNetworkPolicy(ctx, &services.DeleteNetworkPolicyRequest{
				ID: policy.GetUUID(),
			})
			assert.NoError(t, err)
		}()

		networkResponse, err := c.CreateVirtualNetwork(ctx, &services.CreateVirtualNetworkRequest{
			VirtualNetwork: &models.VirtualNetwork{
				ParentType: models.KindProject,
				ParentUUID: project.GetUUID(),
				NetworkPolicyRefs: []*models.VirtualNetworkNetworkPolicyRef{
					{
						UUID: policy.GetUUID(),
					},
				},
			},
		})
		require.NoError(t, err)
		network := networkResponse.GetVirtualNetwork()

		defer func() {
			_, err = c.DeleteVirtualNetwork(ctx, &services.DeleteVirtualNetworkRequest{
				ID: network.GetUUID(),
			})
			assert.NoError(t, err)
		}()

		r := services.NewRefRelaxClient(conn)
		response, err := r.RelaxRef(ctx, &services.RelaxRefRequest{
			UUID:    network.GetUUID(),
			RefUUID: policy.GetUUID(),
		})
		assert.NoError(t, err)
		assert.Equal(t, &services.RelaxRefResponse{
			UUID: network.GetUUID(),
		}, response)

		_, err = c.DeleteNetworkPolicy(ctx, &services.DeleteNetworkPolicyRequest{
			ID: policy.GetUUID(),
		})
		if assert.NoError(t, err) {
			policy = nil
		}
	})
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
