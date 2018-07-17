package compilation_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/coreos/etcd/clientv3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/convert"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

const (
	stateWithProjectPath        = "testdata/state_with_project.yml"
	stateWithVirtualNetworkPath = "testdata/state_with_virtual_network.yml"

	expectedEgressAccessControlListPath  = "testdata/egress_access_control_list.yml"
	expectedIngressAccessControlListPath = "testdata/ingress_access_control_list.yml"
	expectedApplicationPolicySetPath     = "testdata/application_policy_set.yml"
	requestedDemoProjectPath             = "testdata/requested_demo_project.yml"
	expectedDemoProjectPath              = "testdata/demo_project.yml"
	requestedSecurityGroupPath           = "testdata/requested_security_group.yml"
	expectedSecurityGroupPath            = "testdata/security_group.yml"
)

func TestIntentCompilationServiceProcessesEvents(t *testing.T) {
	closeIntentCompilation := integration.RunIntentCompilationService(t)
	defer closeIntentCompilation()
	ec := integration.NewEtcdClient(t)
	defer ec.Close(t)

	tests := []struct {
		dbDriver string
	}{
		{dbDriver: db.DriverMySQL},
		{dbDriver: db.DriverPostgreSQL},
	}

	for _, tt := range tests {
		t.Run(tt.dbDriver, func(t *testing.T) {
			// TODO: run DB cache

			s := integration.NewRunningAPIServer(t, "../..", tt.dbDriver, true)
			defer s.Close(t)
			hc := integration.NewHTTPAPIClient(t, s.URL())

			t.Run("create Project", testCreateProject(hc, ec))
			t.Run("create Virtual Network with Subnet", testCreateVirtualNetworkWithSubnet())
			t.Run("create Virtual Machines", testCreateVMs())
		})
	}
}

func testCreateProject(hc *integration.HTTPAPIClient, ec *integration.EtcdClient) func(t *testing.T) {
	return func(t *testing.T) {
		t.Skip("Not implemented") // TODO: implement compilation service functionality

		defaultDomainUUID := hc.FQNameToID(t, []string{"default-domain", "default-project"}, integration.ProjectType)

		expectedAPS := loadResource(t, expectedApplicationPolicySetPath)
		apsWatch, apsCtx, cancelAPSCtx := ec.WatchResource(
			integration.ApplicationPolicySetSchemaID,
			expectedAPS["uuid"].(string),
		)
		defer cancelAPSCtx()

		aclWatch, aclCtx, cancelACLCtx := ec.WatchResource(
			integration.AccessControlListSchemaID,
			"",
			clientv3.WithPrefix(),
		)
		defer cancelACLCtx()

		project := loadProject(t, requestedDemoProjectPath)
		projectWatch, projectCtx, cancelProjectCtx := ec.WatchResource(integration.ProjectSchemaID, project.UUID)
		defer cancelProjectCtx()

		hc.CreateProject(t, project)
		defer hc.DeleteProject(t, project.UUID)
		defer ec.DeleteProject(t, project.UUID)

		sg := loadSecurityGroup(t, requestedSecurityGroupPath)
		sgWatch, sgCtx, cancelSGCtx := ec.WatchResource(integration.SecurityGroupSchemaID, sg.UUID)
		defer cancelSGCtx()

		hc.CreateSecurityGroup(t, sg)
		defer hc.DeleteSecurityGroup(t, sg.UUID)
		defer ec.DeleteSecurityGroup(t, sg.UUID)

		// TODO: quota updates
		// PUT /project/950b2912-a742-47c8-acdb-ab361f17386 quota update for subnet
		// PUT /project/950b2912-a742-47c8-acdb-ab361f173867 quota update for vn
		// PUT /project/950b2912-a742-47c8-acdb-ab361f173867 quota update for floating ip
		// PUT /project/950b2912-a742-47c8-acdb-ab361f173867 quota update for security group rule
		// PUT /project/950b2912-a742-47c8-acdb-ab361f173867 quota update for security group
		// PUT /project/950b2912-a742-47c8-acdb-ab361f173867 quota update for logical router
		// PUT /project/950b2912-a742-47c8-acdb-ab361f173867 quota update for VMI

		projectEvent := integration.RetrieveCreateEvent(projectCtx, t, projectWatch)
		expectedDemoProject := loadResource(t, expectedDemoProjectPath)
		expectedDemoProject["parent_uuid"] = defaultDomainUUID
		// TODO: verify application_policy_set_refs is uuid of default-application-policy-set
		testutil.AssertEqual(t, expectedDemoProject, decodeJSON(t, projectEvent.Kv.Value))

		apsEvent := integration.RetrieveCreateEvent(apsCtx, t, apsWatch)
		testutil.AssertEqual(t, expectedAPS, decodeJSON(t, apsEvent.Kv.Value))

		sgEvent := integration.RetrieveCreateEvent(sgCtx, t, sgWatch)
		testutil.AssertEqual(t, loadResource(t, expectedSecurityGroupPath), decodeJSON(t, sgEvent.Kv.Value))

		checkCreatedACLs(aclCtx, t, aclWatch)
	}
}

func checkCreatedACLs(aclCtx context.Context, t *testing.T, aclWatch clientv3.WatchChan) {
	aclEvents := integration.RetrieveWatchEvents(aclCtx, t, aclWatch)

	if assert.Equal(t, 2, len(aclEvents)) {
		aclOne := decodeJSON(t, aclEvents[0].Kv.Value)
		aclTwo := decodeJSON(t, aclEvents[1].Kv.Value)
		if aclOne["display_name"] == "ingress-access-control-list" {
			testutil.AssertEqual(t, expectedIngressAccessControlListPath, aclOne)
			testutil.AssertEqual(t, expectedEgressAccessControlListPath, aclTwo)
		} else if aclOne["display_name"] == "egress-access-control-list" {
			testutil.AssertEqual(t, expectedEgressAccessControlListPath, aclOne)
			testutil.AssertEqual(t, expectedIngressAccessControlListPath, aclTwo)
		} else {
			assert.Fail(t, "unexpected ACL: %+v", aclOne)
		}
	}
}

func testCreateVirtualNetworkWithSubnet() func(t *testing.T) {
	return func(t *testing.T) {
		t.Skip("Not implemented") // TODO: implement compilation service functionality

		loadDBSnapshot(t, stateWithProjectPath)
		// TODO: defer cleanupDB()

		// TODO: spawn watch on virtual network
		// TODO: spawn watch on routing instance
		// TODO: spawn watch on route target

		// TODO: create virtual network

		// TODO: check virtual network in etcd
		// TODO: check routing instance in etcd
		// TODO: check route target in etcd
	}
}

func testCreateVMs() func(t *testing.T) {
	return func(t *testing.T) {
		t.Skip("Not implemented") // TODO: implement compilation service functionality

		loadDBSnapshot(t, stateWithVirtualNetworkPath)
		// TODO: defer cleanupDB()

		// TODO: spawn watch on VM-related resources

		// TODO: create VM-related resources

		// TODO: check all resources in etcd after VM-related requests
	}
}

func loadDBSnapshot(t *testing.T, sourceFile string) {
	err := convert.Convert(&convert.Config{
		InType:  convert.YAMLType,
		InFile:  sourceFile,
		OutType: convert.RDBMSType,
	})
	assert.NoError(t, err, "could not load initial database snapshot")
}

func loadResource(t *testing.T, filePath string) map[string]interface{} {
	var r map[string]interface{}
	readYAMLFile(t, filePath, r)
	return r
}

func loadProject(t *testing.T, filePath string) *models.Project {
	var p *models.Project
	readYAMLFile(t, filePath, p)
	return p
}

func loadSecurityGroup(t *testing.T, filePath string) *models.SecurityGroup {
	var sg *models.SecurityGroup
	readYAMLFile(t, filePath, sg)
	return sg
}

func readYAMLFile(t *testing.T, filePath string, data interface{}) {
	bytes, err := ioutil.ReadFile(filePath)
	require.NoError(t, err)

	err = yaml.Unmarshal(bytes, data)
	require.NoError(t, err)
}

func decodeJSON(t *testing.T, bytes []byte) map[string]interface{} {
	var r map[string]interface{}
	assert.NoError(t, json.Unmarshal(bytes, &r))
	return r
}
