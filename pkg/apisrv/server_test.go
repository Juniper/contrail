package apisrv_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/twinj/uuid"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

func TestTypeLogic(t *testing.T) {
	for _, testName := range []string{
		"alarm",
		"alias_ip",
		"bgpaas",
		"domain",
		"floating_ip",
		"firewall_policy",
		"firewall_rule",
		"forwarding_class",
		"instance_ip",
		"k8s_instance_ip_alloc",
		"logical_interface",
		"logical_router",
		"logical_router_vxlan_id",
		"network_ipam",
		"network_policy",
		"physical_interface",
		"project",
		"provisioning",
		"qos",
		"qos_config",
		"security_group",
		"service_template",
		"set_tag",
		"tag",
		"tag_type",
		"virtual_machine_interface",
		"virtual_network",
		"virtual_network_multi_chain",
		"virtual_network_vxlan_id",
		"virtual_router",
	} {
		t.Run(testName, func(t *testing.T) {
			RunTest(t, "./test_data/test_"+testName+".yml")
		})
	}
}

func TestServer(t *testing.T) {
	for _, test := range []string{
		"base_config_root_parent",
		"base_properties",
		"base_props_two_parents",
		"chown",
		"derived_relations",
		"fqname_to_id",
		"id_to_fqname",
		"int_pool",
		"keystone",
		"kv_store",
		"name_unique",
		"obj_perms",
		"parse_id_perms_uuid",
		"project_conflict",
		"prop_collection_update",
		"quota_checking",
		"ref_read",
		"ref_relax",
		"ref_relax_invalid_input",
		"ref_update",
		"sanitizing",
		"sync",
		"sync_sort",
		"user_visible",
		"validation",
	} {
		t.Run(test, func(t *testing.T) {
			RunTest(t, "./test_data/test_"+test+".yml")
		})
	}
}

func TestRESTClient(t *testing.T) {
	restClient, err := integration.NewHTTPClient(server.URL())
	require.NoError(t, err)

	project := models.MakeProject()
	project.UUID = uuid.NewV4().String()
	project.FQName = []string{"default-domain", "project", project.UUID}
	project.ParentType = "domain"
	project.ParentUUID = integration.DefaultDomainUUID
	project.ConfigurationVersion = 1
	project.IDPerms = &models.IdPermsType{UserVisible: true}
	ctx := context.Background()
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

	RunTestTemplate(t, "./test_data/test_pagination.tmpl", context)
}

func TestUploadFileHTTPEndpoint(t *testing.T) {
	const outputDirectory = "test_data/test_upload"
	wd, err := os.Getwd()
	require.NoError(t, err)

	for _, tt := range []struct {
		name       string
		path       string
		content    string
		statusCode int
	}{
		{
			name:       "valid relative path",
			path:       path.Join(outputDirectory, "file"),
			content:    "test content",
			statusCode: http.StatusOK,
		},
		{
			name:       "valid absolute path",
			path:       path.Join(wd, outputDirectory, "file"),
			content:    "test content",
			statusCode: http.StatusOK,
		},
		{
			name:       "relative path with nonexistent directories",
			path:       path.Join(outputDirectory, "bar", "baz", "file"),
			content:    "test content",
			statusCode: http.StatusOK,
		},
		{
			name:       "absolute path with nonexistent directories",
			path:       path.Join(wd, outputDirectory, "bar", "baz", "file"),
			content:    "test content",
			statusCode: http.StatusOK,
		},
		{
			name:       "empty path",
			path:       "",
			content:    "test content",
			statusCode: http.StatusBadRequest,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			err = os.RemoveAll(outputDirectory)
			require.NoError(t, err)

			err = os.MkdirAll(outputDirectory, 0755)
			require.NoError(t, err)

			defer func() {
				if err = os.RemoveAll(outputDirectory); err != nil {
					fmt.Println("Failed to clean up", outputDirectory, ":", err)
				}
			}()

			hc, err := integration.NewHTTPClient(server.URL())
			require.NoError(t, err)

			r, err := hc.UploadFile(context.Background(), tt.path, tt.content)

			assert.Equal(t, tt.statusCode, r.StatusCode)

			if tt.statusCode != http.StatusOK {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)

			b, err := ioutil.ReadFile(tt.path)
			assert.NoError(t, err)
			assert.Equal(t, tt.content, string(b))
		})
	}
}
