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

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

const (
	expectedEgressAccessControlListPath  = "testdata/egress_access_control_list.yml"
	expectedIngressAccessControlListPath = "testdata/ingress_access_control_list.yml"
	expectedApplicationPolicySetPath     = "testdata/application_policy_set.yml"
	demoProjectRequestPath               = "testdata/demo_project_request.yml"
	demoProjectQuotaUpdatePath           = "testdata/demo_project_quota_update.yml"
	expectedDemoProjectPath              = "testdata/demo_project.yml"
	securityGroupRequestPath             = "testdata/security_group_request.yml"
	expectedSecurityGroupPath            = "testdata/security_group.yml"
)

func TestCreateCoreResources(t *testing.T) {
	cacheDB, cancelEtcdEventProducer := integration.RunCacheDB(t)
	defer cancelEtcdEventProducer()

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
			s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{
				DBDriver:           tt.dbDriver,
				EnableEtcdNotifier: true,
				RepoRootPath:       "../../..",
				CacheDB:            cacheDB,
			})
			defer s.Close(t)

			hc := integration.NewHTTPAPIClient(t, s.URL())

			t.Run("create Project and Security Group", testCreateProjectAndSecurityGroup(hc, ec))
		})
	}
}

func testCreateProjectAndSecurityGroup(hc *integration.HTTPAPIClient, ec *integration.EtcdClient) func(t *testing.T) {
	return func(t *testing.T) {
		t.Skip("Not implemented") // TODO: implement API Server and Compilation Service functionality

		defaultDomainUUID := hc.FQNameToID(t, []string{"default-domain", "default-project"}, integration.ProjectType)

		apsWatch, apsCtx, cancelAPSCtx := ec.WatchResource(integration.ApplicationPolicySetSchemaID, "",
			clientv3.WithPrefix())
		defer cancelAPSCtx()

		aclWatch, aclCtx, cancelACLCtx := ec.WatchResource(integration.AccessControlListSchemaID, "",
			clientv3.WithPrefix())
		defer cancelACLCtx()

		project := loadProject(t, demoProjectRequestPath)
		projectWatch, projectCtx, cancelProjectCtx := ec.WatchResource(integration.ProjectSchemaID, project.UUID)
		defer cancelProjectCtx()

		hc.CreateProject(t, project)
		defer hc.DeleteProject(t, project.UUID)
		defer ec.DeleteProject(t, project.UUID)

		sg := loadSecurityGroup(t, securityGroupRequestPath)
		sgWatch, sgCtx, cancelSGCtx := ec.WatchResource(integration.SecurityGroupSchemaID, sg.UUID)
		defer cancelSGCtx()

		hc.CreateSecurityGroup(t, sg)
		defer hc.DeleteSecurityGroup(t, sg.UUID)
		defer ec.DeleteSecurityGroup(t, sg.UUID)

		hc.Chown(t, project.UUID, sg.UUID)

		hc.UpdateProject(t, project.UUID, loadResource(t, demoProjectQuotaUpdatePath))

		apsEvent := integration.RetrieveCreateEvent(apsCtx, t, apsWatch)
		if apsEvent != nil {
			aps := decodeJSON(t, apsEvent.Kv.Value)
			testutil.AssertEqual(t, loadResource(t, expectedApplicationPolicySetPath), aps)

			projectEvents := integration.RetrieveWatchEvents(projectCtx, t, projectWatch)
			checkCreatedProject(t, defaultDomainUUID, aps["uuid"].(string), projectEvents[len(projectEvents)-1])
		}

		sgEvent := integration.RetrieveCreateEvent(sgCtx, t, sgWatch)
		testutil.AssertEqual(t, loadResource(t, expectedSecurityGroupPath), decodeJSON(t, sgEvent.Kv.Value))

		checkCreatedACLs(aclCtx, t, aclWatch)
	}
}

func checkCreatedProject(t *testing.T, defaultDomainUUID, apsUUID string, projectEvent *clientv3.Event) {
	expectedDemoProject := loadResource(t, expectedDemoProjectPath)
	expectedDemoProject["parent_uuid"] = defaultDomainUUID
	expectedDemoProject["application_policy_set_refs"].([]map[string]interface{})[0]["uuid"] = apsUUID
	testutil.AssertEqual(t, expectedDemoProject, decodeJSON(t, projectEvent.Kv.Value))
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

func loadResource(t *testing.T, filePath string) map[string]interface{} {
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
