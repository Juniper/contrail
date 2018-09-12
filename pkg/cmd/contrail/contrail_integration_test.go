package contrail_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/coreos/etcd/clientv3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/Juniper/contrail/pkg/testutil/integration/etcd"
)

const (
	expectedEgressAccessControlListPath  = "testdata/egress_access_control_list.yml"
	expectedIngressAccessControlListPath = "testdata/ingress_access_control_list.yml"
	expectedApplicationPolicySetPath     = "testdata/application_policy_set.yml"
	demoProjectRequestPath               = "testdata/demo_project_request.yml"
	demoProjectQuotaUpdatePath           = "testdata/demo_project_quota_update.json"
	expectedDemoProjectPath              = "testdata/demo_project.yml"
	securityGroupRequestPath             = "testdata/security_group_request.yml"
	expectedSecurityGroupPath            = "testdata/security_group.yml"
)

func TestCreateCoreResources(t *testing.T) {
	t.Skip("Not implemented") // TODO: implement API Server and Compilation Service functionality

	cacheDB, cancelEtcdEventProducer := integration.RunCacheDB(t)
	defer cancelEtcdEventProducer()

	closeIntentCompilation := integration.RunIntentCompilationService(t)
	defer closeIntentCompilation()

	ec := integrationetcd.NewEtcdClient(t)
	defer ec.Close(t)

	tests := []struct {
		dbDriver string
	}{
		{dbDriver: basedb.DriverMySQL},
		{dbDriver: basedb.DriverPostgreSQL},
	}

	for _, tt := range tests {
		t.Run(tt.dbDriver, func(t *testing.T) {
			s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{
				DBDriver:           tt.dbDriver,
				EnableEtcdNotifier: true,
				RepoRootPath:       "../../..",
				CacheDB:            cacheDB,
			})
			defer s.CloseT(t)

			hc := integration.NewTestingHTTPClient(t, s.URL())

			t.Run("create Project and Security Group", testCreateProjectAndSecurityGroup(hc, ec))
		})
	}
}

func testCreateProjectAndSecurityGroup(hc *integration.HTTPAPIClient, ec *integrationetcd.EtcdClient) func(t *testing.T) {
	return func(t *testing.T) {
		t.Skip("Not implemented") // TODO: implement API Server and Compilation Service functionality

		defaultDomainUUID := hc.FQNameToID(t, []string{"default-domain", "default-project"}, integration.ProjectType)

		apsWatch, apsCtx, cancelAPSCtx := ec.WatchResource(integrationetcd.ApplicationPolicySetSchemaID, "",
			clientv3.WithPrefix())
		defer cancelAPSCtx()

		aclWatch, aclCtx, cancelACLCtx := ec.WatchResource(integrationetcd.AccessControlListSchemaID, "",
			clientv3.WithPrefix())
		defer cancelACLCtx()

		project := loadProject(t, demoProjectRequestPath)
		projectWatch, projectCtx, cancelProjectCtx := ec.WatchResource(integrationetcd.ProjectSchemaID, project.UUID)
		defer cancelProjectCtx()

		// TODO: creating project fails with following message:
		// "Validation failed for resource with UUID 73a4ad89-3455-46f5-b37c-394aeb1f7c87:
		// quota property is missing for resource project"
		hc.CreateProject(t, project)
		defer hc.DeleteProject(t, project.UUID)
		defer ec.DeleteProject(t, project.UUID)

		sg := loadSecurityGroup(t, securityGroupRequestPath)
		sgWatch, sgCtx, cancelSGCtx := ec.WatchResource(integrationetcd.SecurityGroupSchemaID, sg.UUID)
		defer cancelSGCtx()

		// TODO: creating security group fails with following message:
		// "Please provide correct FQName or ParentUUID"
		hc.CreateSecurityGroup(t, sg)
		defer hc.DeleteSecurityGroup(t, sg.UUID)
		defer ec.DeleteSecurityGroup(t, sg.UUID)

		// TODO: Chown endpoint fails with following message: "Not Found"
		hc.Chown(t, project.UUID, sg.UUID)

		hc.UpdateProject(t, project.UUID, loadResourceJSON(t, demoProjectQuotaUpdatePath))

		apsEvent := integrationetcd.RetrieveCreateEvent(apsCtx, t, apsWatch)
		if apsEvent != nil {
			aps := decodeJSON(t, apsEvent.Kv.Value)
			testutil.AssertEqual(t, loadResourceYAML(t, expectedApplicationPolicySetPath), aps)

			projectEvents := integrationetcd.RetrieveWatchEvents(projectCtx, t, projectWatch)
			if len(projectEvents) > 0 {
				checkCreatedProject(t, defaultDomainUUID, aps["uuid"].(string), projectEvents[len(projectEvents)-1])
			}
		}

		retrieveAndCheckCreatedSecurityGroup(sgCtx, t, sgWatch)
		retrieveAndCheckCreatedACLs(aclCtx, t, aclWatch)
	}
}

func checkCreatedProject(t *testing.T, defaultDomainUUID, apsUUID string, projectEvent *clientv3.Event) {
	expectedDemoProject := loadResourceYAML(t, expectedDemoProjectPath)
	expectedDemoProject["parent_uuid"] = defaultDomainUUID
	expectedDemoProject["application_policy_set_refs"].([]interface{})[0].(map[string]interface{})["uuid"] = apsUUID
	testutil.AssertEqual(t, expectedDemoProject, decodeJSON(t, projectEvent.Kv.Value))
}

func retrieveAndCheckCreatedSecurityGroup(sgCtx context.Context, t *testing.T, sgWatch clientv3.WatchChan) {
	sgEvent := integrationetcd.RetrieveCreateEvent(sgCtx, t, sgWatch)
	if sgEvent != nil {
		testutil.AssertEqual(t, loadResourceYAML(t, expectedSecurityGroupPath), decodeJSON(t, sgEvent.Kv.Value))
	}
}

func retrieveAndCheckCreatedACLs(aclCtx context.Context, t *testing.T, aclWatch clientv3.WatchChan) {
	aclEvents := integrationetcd.RetrieveWatchEvents(aclCtx, t, aclWatch)
	if !assert.Equal(t, 2, len(aclEvents)) {
		return
	}

	aclOne := decodeJSON(t, aclEvents[0].Kv.Value)
	aclTwo := decodeJSON(t, aclEvents[1].Kv.Value)
	if aclOne["display_name"] == "ingress-access-control-list" {
		testutil.AssertEqual(t, expectedIngressAccessControlListPath, aclOne)
		testutil.AssertEqual(t, expectedEgressAccessControlListPath, aclTwo)
	} else if aclOne["display_name"] == "egress-access-control-list" {
		testutil.AssertEqual(t, expectedEgressAccessControlListPath, aclOne)
		testutil.AssertEqual(t, expectedIngressAccessControlListPath, aclTwo)
	} else {
		assert.Fail(t, "unexpected ACL display_name: %+v", aclOne)
	}
}

func loadResourceJSON(t *testing.T, filePath string) interface{} {
	var r map[string]interface{}
	readJSONFile(t, filePath, &r)
	return r
}

func loadResourceYAML(t *testing.T, filePath string) map[string]interface{} {
	var r map[string]interface{}
	readYAMLFile(t, filePath, &r)
	return r
}

func loadProject(t *testing.T, filePath string) *models.Project {
	var p models.Project
	readYAMLFile(t, filePath, &p)
	return &p
}

func loadSecurityGroup(t *testing.T, filePath string) *models.SecurityGroup {
	var sg models.SecurityGroup
	readYAMLFile(t, filePath, &sg)
	return &sg
}

func readJSONFile(t *testing.T, filePath string, data interface{}) {
	err := json.Unmarshal(readFile(t, filePath), data)
	require.NoError(t, err)
}

func readYAMLFile(t *testing.T, filePath string, data interface{}) {
	err := yaml.Unmarshal(readFile(t, filePath), data)
	require.NoError(t, err)
}

func readFile(t *testing.T, filePath string) []byte {
	bytes, err := ioutil.ReadFile(filePath)
	require.NoError(t, err)
	return bytes
}

func decodeJSON(t *testing.T, bytes []byte) map[string]interface{} {
	var r map[string]interface{}
	assert.NoError(t, json.Unmarshal(bytes, &r))
	return r
}
