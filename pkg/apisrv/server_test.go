package apisrv

import (
	"context"
	"crypto/tls"
	"strings"
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/generated/services"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func TestServer(t *testing.T) {
	err := RunTest("./test_data/test_virtual_network.yml")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGRPC(t *testing.T) {
	common.UseTable(server.DB, "metadata")
	defer common.ClearTable(server.DB, "metadata")
	common.UseTable(server.DB, "project")
	defer common.ClearTable(server.DB, "project")

	restClient := NewClient(
		testServer.URL,
		testServer.URL+"/v3",
		"alice",
		"alice_password",
		&keystone.Scope{
			Project: &keystone.Project{
				ID: "admin",
			},
		},
	)
	restClient.InSecure = true
	restClient.Init()
	err := restClient.Login()
	assert.NoError(t, err)
	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})
	dial := strings.TrimPrefix(testServer.URL, "https://")
	conn, err := grpc.Dial(dial, grpc.WithTransportCredentials(creds))
	assert.NoError(t, err)
	defer conn.Close()
	md := metadata.Pairs("X-Auth-Token", restClient.AuthToken)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	// Contact the server and print out its response.
	c := services.NewContrailServiceClient(conn)
	assert.NoError(t, err)
	project := models.MakeProject()
	project.UUID = "test_project"
	project.FQName = []string{"default-domain", "project", "test"}
	_, err = c.CreateProject(ctx, &models.CreateProjectRequest{
		Project: project,
	})
	assert.NoError(t, err)
	response, err := c.ListProject(ctx, &models.ListProjectRequest{
		Spec: &models.ListSpec{
			Limit: 1,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(response.Projects))

	getResponse, err := c.GetProject(ctx, &models.GetProjectRequest{
		ID: project.UUID,
	})
	assert.NoError(t, err)
	assert.NotNil(t, getResponse.Project)

	_, err = c.DeleteProject(ctx, &models.DeleteProjectRequest{
		ID: project.UUID,
	})
	assert.NoError(t, err)
}
