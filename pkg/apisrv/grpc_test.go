package apisrv_test

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"
	"testing"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/Juniper/contrail/pkg/types"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	asfclient "github.com/Juniper/asf/pkg/client"
	kstypes "github.com/Juniper/asf/pkg/keystone"
	protocodec "github.com/gogo/protobuf/codec"
	uuid "github.com/satori/go.uuid"
)

/////////////////////////
// Type-agnostic tests //
/////////////////////////

func TestFQNameToIDGRPC(t *testing.T) {
	testGRPCServer(t, t.Name(), func(ctx context.Context, conn *grpc.ClientConn) {
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

func TestIDToFQNameGRPC(t *testing.T) {
	testGRPCServer(t, t.Name(), func(ctx context.Context, conn *grpc.ClientConn) {
		c := services.NewIDToFQNameClient(conn)
		resp, err := c.IDToFQName(ctx, &services.IDToFQNameRequest{
			UUID: integration.DefaultDomainUUID,
		})
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, models.KindDomain, resp.Type)
	})
}

func TestChownGRPC(t *testing.T) {
	firstProjectName := uuid.NewV4().String()
	testGRPCServer(t, firstProjectName, func(firstProjectCTX context.Context, conn *grpc.ClientConn) {
		c := client.NewGRPC(services.NewContrailServiceClient(conn))

		project, cleanup := createProjectWithName(firstProjectCTX, t, c, firstProjectName)
		defer cleanup(t)

		otherProjectName := uuid.NewV4().String()
		addKeystoneProjectAndUser(server.APIServer, otherProjectName)
		otherProjectCTX := metadata.NewOutgoingContext(firstProjectCTX,
			metadata.Pairs("X-Auth-Token", restLogin(firstProjectCTX, t, otherProjectName)))

		_, cleanup = createProjectWithName(otherProjectCTX, t, c, otherProjectName)
		defer cleanup(t)

		networkResponse, err := c.CreateVirtualNetwork(firstProjectCTX, &services.CreateVirtualNetworkRequest{
			VirtualNetwork: &models.VirtualNetwork{
				ParentType: models.KindProject,
				ParentUUID: project.GetUUID(),
			},
		})
		require.NoError(t, err, "creating network failed")
		network := networkResponse.GetVirtualNetwork()

		defer func() {
			_, err = c.DeleteVirtualNetwork(firstProjectCTX, &services.DeleteVirtualNetworkRequest{
				ID: network.GetUUID(),
			})
			assert.NoError(t, err, "deleting network failed")
		}()

		ch := services.NewChownClient(conn)

		_, err = ch.Chown(firstProjectCTX, &services.ChownRequest{})
		assert.True(t, errutil.IsBadRequest(err), "chown with an empty request should fail")

		_, err = ch.Chown(firstProjectCTX, &services.ChownRequest{
			Owner: otherProjectName,
		})
		assert.True(t, errutil.IsBadRequest(err), "chown with an empty UUID should fail")

		_, err = ch.Chown(firstProjectCTX, &services.ChownRequest{
			UUID: network.GetUUID(),
		})
		assert.True(t, errutil.IsBadRequest(err), "chown with an empty Owner should fail")

		_, err = ch.Chown(firstProjectCTX, &services.ChownRequest{
			UUID:  "not a UUID",
			Owner: otherProjectName,
		})
		assert.True(t, errutil.IsBadRequest(err), "chown with a non-UUID UUID should fail")

		_, err = ch.Chown(firstProjectCTX, &services.ChownRequest{
			UUID:  network.GetUUID(),
			Owner: "not a UUID",
		})
		assert.True(t, errutil.IsBadRequest(err), "chown with a non-UUID Owner should fail")

		_, err = ch.Chown(firstProjectCTX, &services.ChownRequest{
			UUID:  uuid.NewV4().String(),
			Owner: otherProjectName,
		})
		assert.Error(t, err, "chown of a nonexistent resource should fail")

		_, err = ch.Chown(firstProjectCTX, &services.ChownRequest{
			UUID:  network.GetUUID(),
			Owner: otherProjectName,
		})
		assert.NoError(t, err, "chown failed")

		_, err = c.GetVirtualNetwork(firstProjectCTX, &services.GetVirtualNetworkRequest{
			ID: network.GetUUID(),
		})
		assert.True(t, errutil.IsNotFound(err), "the old owner should not be able to get the network after chown")

		chownedNetworkResponse, err := c.GetVirtualNetwork(otherProjectCTX, &services.GetVirtualNetworkRequest{
			ID: network.GetUUID(),
		})
		assert.NoError(t, err, "the new owner should be able to get the network")
		assert.Equal(t, otherProjectName, chownedNetworkResponse.GetVirtualNetwork().GetPerms2().GetOwner())

		_, err = ch.Chown(otherProjectCTX, &services.ChownRequest{
			UUID:  network.GetUUID(),
			Owner: firstProjectName,
		})
		assert.NoError(t, err, "chown back to the original owner failed")

		chownedBackNetworkResponse, err := c.GetVirtualNetwork(firstProjectCTX, &services.GetVirtualNetworkRequest{
			ID: network.GetUUID(),
		})
		assert.NoError(t, err, "the original owner should be able to get the network")
		assert.Equal(t, firstProjectName, chownedBackNetworkResponse.GetVirtualNetwork().GetPerms2().GetOwner())
	})
}

func TestIPAMGRPC(t *testing.T) {
	testGRPCServer(t, t.Name(), func(ctx context.Context, conn *grpc.ClientConn) {
		c := services.NewIPAMClient(conn)
		allocateResp, err := c.AllocateInt(ctx, &services.AllocateIntRequest{
			Pool:  types.VirtualNetworkIDPoolKey,
			Owner: db.EmptyIntOwner,
		})
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

		project, cleanup := createProject(ctx, t, c)
		defer cleanup(t)

		policyResponse, err := c.CreateNetworkPolicy(ctx, &services.CreateNetworkPolicyRequest{
			NetworkPolicy: &models.NetworkPolicy{
				ParentType: models.KindProject,
				ParentUUID: project.GetUUID(),
				IDPerms:    &models.IdPermsType{UserVisible: true},
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
				IDPerms:    &models.IdPermsType{UserVisible: true},
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

func TestPropCollectionUpdateGRPC(t *testing.T) {
	testGRPCServer(t, t.Name(), func(ctx context.Context, conn *grpc.ClientConn) {
		gc := client.NewGRPC(services.NewContrailServiceClient(conn))
		project, cleanup := createProject(ctx, t, gc)
		defer cleanup(t)

		r := &services.PropCollectionUpdateRequest{
			UUID: project.UUID,
			Updates: []*services.PropCollectionChange{{
				Field:     models.DomainFieldAnnotations,
				Operation: basemodels.PropCollectionUpdateOperationSet,
				Value: &services.PropCollectionChange_KeyValuePairValue{
					KeyValuePairValue: &models.KeyValuePair{
						Key:   "some-key",
						Value: "some-value",
					},
				},
			}},
		}
		b, err := proto.Marshal(r)
		fmt.Println(b, len(b), err)
		r1 := &services.PropCollectionUpdateRequest{}
		err = proto.Unmarshal(b, r1)
		fmt.Println(r1, err)

		c := services.NewPropCollectionUpdateClient(conn)
		_, err = c.PropCollectionUpdate(ctx, &services.PropCollectionUpdateRequest{
			UUID: project.UUID,
			Updates: []*services.PropCollectionChange{{
				Field:     models.DomainFieldAnnotations,
				Operation: basemodels.PropCollectionUpdateOperationSet,
				Value: &services.PropCollectionChange_KeyValuePairValue{
					KeyValuePairValue: &models.KeyValuePair{
						Key:   "some-key",
						Value: "some-value",
					},
				},
			}},
		})
		assert.NoError(t, err)

		_, err = c.PropCollectionUpdate(ctx, &services.PropCollectionUpdateRequest{
			UUID: project.UUID,
			Updates: []*services.PropCollectionChange{{
				Field:     models.DomainFieldAnnotations,
				Operation: basemodels.PropCollectionUpdateOperationDelete,
				Position: &services.PropCollectionChange_PositionString{
					PositionString: "some-key",
				},
			}},
		})
		assert.NoError(t, err)
	})
}

func createProject(
	ctx context.Context, t *testing.T, c *client.GRPC,
) (project *models.Project, cleanup func(t *testing.T)) {
	return createProjectWithName(ctx, t, c, fmt.Sprintf("%s_project", t.Name()))
}

func createProjectWithName(
	ctx context.Context, t *testing.T, c *client.GRPC, name string,
) (project *models.Project, cleanup func(t *testing.T)) {
	r, err := c.CreateProject(ctx, &services.CreateProjectRequest{
		Project: &models.Project{
			Name:       name,
			ParentType: "domain",
			ParentUUID: integration.DefaultDomainUUID,
			IDPerms:    &models.IdPermsType{UserVisible: true},
		},
	})
	require.NoError(t, err)

	return r.GetProject(), func(t *testing.T) {
		_, err := c.DeleteProject(ctx, &services.DeleteProjectRequest{
			ID: project.GetUUID(),
		})
		assert.NoError(t, err)
	}
}

/////////////////////////
// Type-specific tests //
/////////////////////////

func TestGRPC(t *testing.T) {
	testGRPCServer(t, t.Name(), func(ctx context.Context, conn *grpc.ClientConn) {
		c := services.NewContrailServiceClient(conn)
		project := models.Project{
			UUID:                 uuid.NewV4().String(),
			FQName:               []string{"default-domain", "project", "my-project"},
			ParentType:           "domain",
			ParentUUID:           integration.DefaultDomainUUID,
			ConfigurationVersion: 1,
			IDPerms:              &models.IdPermsType{UserVisible: true},
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
			IDPerms:    &models.IdPermsType{UserVisible: true},
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
		assert.NotNil(t, getResponse.GetProject())
	}
}

// Common

func testGRPCServer(t *testing.T, testName string, testBody func(ctx context.Context, conn *grpc.ClientConn)) {
	ctx := context.Background()
	addKeystoneProjectAndUser(server.APIServer, testName)
	authToken := restLogin(ctx, t, testName)

	conn, err := grpc.Dial(
		strings.TrimPrefix(server.URL(), "https://"),
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})),
		grpc.WithDefaultCallOptions(grpc.CallCustomCodec(protocodec.New(0))),
	)
	assert.NoError(t, err)
	defer testutil.LogFatalIfError(conn.Close)

	testBody(
		metadata.NewOutgoingContext(context.Background(), metadata.Pairs("X-Auth-Token", authToken)),
		conn,
	)
}

// addKeystoneProjectAndUser adds Keystone project and user in Server internal state.
// TODO: Remove that, because it modifies internal state of SUT.
// TODO: Use pre-created Server's keystone assignment.
func addKeystoneProjectAndUser(s *apisrv.Server, testID string) {
	assignment := s.Keystone.Assignment.(*keystone.StaticAssignment) // nolint: errcheck
	assignment.Projects[testID] = &kstypes.Project{
		Domain: assignment.Domains[integration.DefaultDomainID],
		ID:     testID,
		Name:   testID,
	}

	assignment.Users[testID] = &kstypes.User{
		Domain:   assignment.Domains[integration.DefaultDomainID],
		ID:       testID,
		Name:     testID,
		Password: testID,
		Roles: []*kstypes.Role{
			{
				ID:      "member",
				Name:    "Member",
				Project: assignment.Projects[testID],
			},
		},
	}
}

func restLogin(ctx context.Context, t *testing.T, projectName string) (authToken string) {
	c := client.NewHTTP(&asfclient.HTTPConfig{
		ID:       projectName,
		Password: projectName,
		Endpoint: server.URL(),
		AuthURL:  server.URL() + keystone.LocalAuthPath,
		Scope:    kstypes.NewScope("", "default", "", projectName),
		Insecure: true,
	})

	err := c.Login(ctx)
	require.NoError(t, err)

	return c.AuthToken
}
