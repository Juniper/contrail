package apiserver_test

import (
	"context"
	"testing"

	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	uuid "github.com/satori/go.uuid"
)

/////////////////////////
// Type-agnostic tests //
/////////////////////////

func TestServer(t *testing.T) {
	for _, test := range []string{
		"auth_skip",
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
			integration.RunTest(t, "./test_data/test_"+test+".yml", server)
		})
	}
}

func TestHTTPCRUD(t *testing.T) {
	hc := integration.NewTestingHTTPClient(t, server.URL(), integration.BobUserID)

	project := models.MakeProject()
	project.UUID = uuid.NewV4().String()
	project.FQName = []string{"default-domain", "project", project.UUID}
	project.ParentType = "domain"
	project.ParentUUID = integration.DefaultDomainUUID
	project.ConfigurationVersion = 1
	project.IDPerms = &models.IdPermsType{UserVisible: true}

	ctx := context.Background()
	_, err := hc.CreateProject(ctx, &services.CreateProjectRequest{
		Project: project,
	})
	assert.NoError(t, err)

	response, err := hc.ListProject(ctx, &services.ListProjectRequest{
		Spec: &baseservices.ListSpec{
			Limit: 1,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(response.Projects))

	getResponse, err := hc.GetProject(ctx, &services.GetProjectRequest{
		ID: project.UUID,
	})
	assert.NoError(t, err)
	assert.NotNil(t, getResponse.Project)
	assert.Equal(t, project.UUID, getResponse.Project.UUID)

	_, err = hc.DeleteProject(ctx, &services.DeleteProjectRequest{
		ID: project.UUID,
	})
	assert.NoError(t, err)
}

func TestHTTPPagination(t *testing.T) {
	integration.RunTestTemplate(
		t,
		"./test_data/test_pagination.tmpl",
		server,
		map[string]interface{}{
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
		},
	)
}

func TestCreateRefMethod(t *testing.T) {
	hc := integration.NewTestingHTTPClient(t, server.URL(), integration.BobUserID)

	testID := "test"
	projectUUID := testID + "_project"
	vnUUID := testID + "_virtual-network"
	niUUID := testID + "_network-ipam"

	integration.CreateProject(
		t,
		hc,
		&models.Project{
			UUID:       projectUUID,
			ParentType: integration.DomainType,
			ParentUUID: integration.DefaultDomainUUID,
			Name:       "testProject",
			Quota:      &models.QuotaType{},
		},
	)
	defer integration.DeleteProject(t, hc, projectUUID)

	integration.CreateNetworkIpam(
		t,
		hc,
		&models.NetworkIpam{
			UUID:       niUUID,
			ParentType: integration.ProjectType,
			ParentUUID: projectUUID,
			Name:       "testIpam",
		},
	)
	defer integration.DeleteNetworkIpam(t, hc, niUUID)

	integration.CreateVirtualNetwork(
		t,
		hc,
		&models.VirtualNetwork{
			UUID:       vnUUID,
			ParentType: integration.ProjectType,
			ParentUUID: projectUUID,
			Name:       "testVN",
			NetworkIpamRefs: []*models.VirtualNetworkNetworkIpamRef{
				{
					UUID: niUUID,
				},
			},
		},
	)

	defer integration.DeleteVirtualNetwork(t, hc, vnUUID)

	// After creating VirtualNetwork it is already connected to networkIpam
	vn := integration.GetVirtualNetwork(t, hc, vnUUID)

	assert.Len(t, vn.NetworkIpamRefs, 1)

	_, err := hc.DeleteVirtualNetworkNetworkIpamRef(
		context.Background(),
		&services.DeleteVirtualNetworkNetworkIpamRefRequest{
			ID: vnUUID,
			VirtualNetworkNetworkIpamRef: &models.VirtualNetworkNetworkIpamRef{
				UUID: niUUID,
			},
		},
	)
	assert.NoError(t, err)

	vn = integration.GetVirtualNetwork(t, hc, vnUUID)

	assert.Len(t, vn.NetworkIpamRefs, 0)

	_, err = hc.CreateVirtualNetworkNetworkIpamRef(
		context.Background(),
		&services.CreateVirtualNetworkNetworkIpamRefRequest{
			ID: vnUUID,
			VirtualNetworkNetworkIpamRef: &models.VirtualNetworkNetworkIpamRef{
				UUID: niUUID,
			},
		},
	)
	assert.NoError(t, err)

	vn = integration.GetVirtualNetwork(t, hc, vnUUID)
	assert.Len(t, vn.NetworkIpamRefs, 1)
}

func TestRemoteIntPoolMethods(t *testing.T) {
	hc := integration.NewTestingHTTPClient(t, server.URL(), integration.AdminUserID)

	err := hc.Login(context.Background())
	require.NoError(t, err)

	err = hc.CreateIntPool(context.Background(), "test_int_pool_806f099f3", 8000100, 8000200)
	require.NoError(t, err)
	defer func() {
		err = hc.DeleteIntPool(context.Background(), "test_int_pool_806f099f3")
		assert.NoError(t, err)
	}()

	val, err := hc.AllocateInt(context.Background(), "test_int_pool_806f099f3", "test_owner_806f099f3")
	defer func() {
		err = hc.DeallocateInt(context.Background(), "test_int_pool_806f099f3", val)
		assert.NoError(t, err)
	}()
	assert.NoError(t, err)
	assert.True(t, val > 8000099)

	owner, err := hc.GetIntOwner(context.Background(), "test_int_pool_806f099f3", val)
	assert.NoError(t, err)
	assert.Equal(t, "test_owner_806f099f3", owner)

	err = hc.SetInt(context.Background(), "test_int_pool_806f099f3", val+1, db.EmptyIntOwner)
	defer func() {
		err = hc.DeallocateInt(context.Background(), "test_int_pool_806f099f3", val+1)
		assert.NoError(t, err)
	}()

	assert.NoError(t, err)
}

/////////////////////////
// Type-specific tests //
/////////////////////////

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
			integration.RunTest(t, "./test_data/test_"+testName+".yml", server)
		})
	}
}

func TestPredefinedTagTypes(t *testing.T) {
	c := integration.NewTestingHTTPClient(t, server.URL(), integration.AdminUserID)

	predefinedTags := []struct {
		fqName    []string
		tagTypeID string
	}{
		{
			fqName:    []string{"label"},
			tagTypeID: "0x0000",
		},
		{
			fqName:    []string{"application"},
			tagTypeID: "0x0001",
		},
		{
			fqName:    []string{"tier"},
			tagTypeID: "0x0002",
		},
		{
			fqName:    []string{"deployment"},
			tagTypeID: "0x0003",
		},
		{
			fqName:    []string{"site"},
			tagTypeID: "0x0004",
		},
	}
	for _, tag := range predefinedTags {
		uuid := c.FQNameToID(t, tag.fqName, "tag-type")
		assert.NotEmpty(t, uuid)

		resp, err := c.GetTagType(context.Background(), &services.GetTagTypeRequest{ID: uuid})
		assert.NoError(t, err)
		assert.Equal(t, tag.tagTypeID, resp.GetTagType().GetTagTypeID())
	}
}
