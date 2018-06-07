package apisrv

import (
	"context"
	"crypto/tls"
	"strings"
	"testing"

	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/stretchr/testify/assert"
	"github.com/twinj/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func TestAPIServer(t *testing.T) {
	CreateTestProject(APIServer, "TestAPIServer")
	RunTest(t, "./test_data/test_virtual_network.yml")
	RunTest(t, "./test_data/test_floating_ip.yml")
}

func TestSync(t *testing.T) {
	CreateTestProject(APIServer, "TestSync")
	RunTest(t, "./test_data/test_sync.yml")
}

func TestGRPC(t *testing.T) {
	CreateTestProject(APIServer, "TestGRPC")
	restClient := NewClient(
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
	err := restClient.Login()
	assert.NoError(t, err)
	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})
	dial := strings.TrimPrefix(TestServer.URL, "https://")
	conn, err := grpc.Dial(dial, grpc.WithTransportCredentials(creds))
	assert.NoError(t, err)
	defer LogFatalIfErr(conn.Close)
	md := metadata.Pairs("X-Auth-Token", restClient.AuthToken)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	// Contact the server and print out its response.
	c := services.NewContrailServiceClient(conn)
	assert.NoError(t, err)
	project := models.MakeProject()
	project.UUID = uuid.NewV4().String()
	project.FQName = []string{"default-domain", "project", project.UUID}
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
