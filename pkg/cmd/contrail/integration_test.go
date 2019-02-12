package contrail_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	yaml "gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	integrationetcd "github.com/Juniper/contrail/pkg/testutil/integration/etcd"
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
	expectedProjectCount                 = 3
	expectedAPSCount                     = 1
	expectedSGCount                      = 1
	expectedACLCount                     = 2
)

func TestFQNameCleanup(t *testing.T) {
	runDirtyTest(t, t.Name())
	integration.RunTest(t, t.Name(), server)
}

func TestProject(t *testing.T) {
	integration.RunTest(t, t.Name(), server)
}

func TestSecurityGroup(t *testing.T) {
	integration.RunTest(t, t.Name(), server)
}

func TestLogicalRouterPing(t *testing.T) {
	t.Skip("Intent compiler methods must be done in transaction otherwise this test is flaky.")
	integration.RunTest(t, t.Name(), server)
}

func TestWaiter(t *testing.T) {
	integration.RunTest(t, t.Name(), server)
}

func TestReferredSecurityGroups(t *testing.T) {
	integration.RunTest(t, t.Name(), server)
}

func TestCreateCoreResources(t *testing.T) {
	if viper.GetBool("sync.enabled") {
		t.Skip("test incompatible with sync process")
	}
	ec := integrationetcd.NewEtcdClient(t)
	defer ec.Close(t)

	closeIntentCompilation := integration.RunIntentCompilationService(t, server.URL())
	defer closeIntentCompilation()

	hc := integration.NewTestingHTTPClient(t, server.URL())

	t.Run("create Project and Security Group", testCreateProjectAndSecurityGroup(hc, ec))
}

func testCreateProjectAndSecurityGroup(
	hc *integration.HTTPAPIClient, ec *integrationetcd.EtcdClient,
) func(t *testing.T) {
	return func(t *testing.T) {
		wTime := 1 * time.Second

		defaultDomainUUID := hc.FQNameToID(t, []string{"default-domain"}, integration.DomainType)

		collectAPSEvs := ec.WatchKeyN(
			integrationetcd.JSONEtcdKey(integrationetcd.ApplicationPolicySetSchemaID, ""),
			expectedAPSCount,
			wTime,
			clientv3.WithPrefix(),
		)
		defer collectAPSEvs()

		collectACLEvs := ec.WatchKeyN(
			integrationetcd.JSONEtcdKey(integrationetcd.AccessControlListSchemaID, ""),
			expectedACLCount,
			wTime,
			clientv3.WithPrefix(),
		)
		defer collectACLEvs()

		project := loadProject(t, demoProjectRequestPath)

		collectProjectEvs := ec.WatchKeyN(
			integrationetcd.JSONEtcdKey(integrationetcd.ProjectSchemaID, project.UUID), expectedProjectCount, wTime,
		)

		defer collectProjectEvs()

		ctx := context.Background()

		_, err := hc.CreateProject(ctx, &services.CreateProjectRequest{Project: project})
		require.NoError(t, err)
		defer integration.DeleteProject(t, hc, project.UUID)

		sg := loadSecurityGroup(t, securityGroupRequestPath)
		collectSGEvs := ec.WatchKeyN(
			integrationetcd.JSONEtcdKey(integrationetcd.SecurityGroupSchemaID, ""),
			expectedSGCount,
			wTime,
			clientv3.WithPrefix(),
		)
		defer collectSGEvs()

		sgResp, err := hc.CreateSecurityGroup(ctx, &services.CreateSecurityGroupRequest{SecurityGroup: sg})
		require.NoError(t, err)
		defer integration.DeleteSecurityGroup(t, hc, sgResp.SecurityGroup.UUID)

		hc.Chown(t, project.UUID, sgResp.SecurityGroup.UUID)

		req := &services.UpdateProjectRequest{}
		readJSONFile(t, demoProjectQuotaUpdatePath, &req)
		req.Project.UUID = project.UUID
		_, err = hc.UpdateProject(ctx, req)
		require.NoError(t, err)

		apsEvents := collectAPSEvs()
		require.Len(t, apsEvents, 1)

		apsOne := decodeJSON(t, []byte(apsEvents[0]))
		testutil.AssertEqual(t, loadResourceYAML(t, expectedApplicationPolicySetPath), apsOne)

		projectEvents := collectProjectEvs()
		if len(projectEvents) > 0 {
			checkCreatedProject(t, defaultDomainUUID, apsOne["uuid"].(string), projectEvents[len(projectEvents)-1])
		}

		checkCreatedSecurityGroup(t, collectSGEvs())
		aclEvents := collectACLEvs()
		for _, ev := range aclEvents {
			acl := decodeJSON(t, []byte(ev))
			defer integration.DeleteAccessControlList(t, hc, acl["uuid"].(string))
		}
		if assert.Equal(t, expectedACLCount, len(aclEvents)) {
			checkCreatedACLs(t, aclEvents)
		}
	}
}

func checkCreatedProject(t *testing.T, defaultDomainUUID, apsUUID string, projectEvent string) {
	expectedDemoProject := loadResourceYAML(t, expectedDemoProjectPath)
	expectedDemoProject["parent_uuid"] = defaultDomainUUID
	expectedDemoProject["application_policy_set_refs"].([]interface{})[0].(map[interface{}]interface{})["uuid"] = apsUUID
	testutil.AssertEqual(t, expectedDemoProject, decodeJSON(t, []byte(projectEvent)))
}

func checkCreatedSecurityGroup(t *testing.T, sgEvents []string) {
	if !assert.Equal(t, expectedSGCount, len(sgEvents)) {
		return
	}
	testutil.AssertEqual(t, loadResourceYAML(t, expectedSecurityGroupPath), decodeJSON(t, []byte(sgEvents[0])))
}

func checkCreatedACLs(t *testing.T, aclEvents []string) {

	aclOne := decodeJSON(t, []byte(aclEvents[0]))
	aclTwo := decodeJSON(t, []byte(aclEvents[1]))

	if aclOne["display_name"] == "ingress-access-control-list" {
		testutil.AssertEqual(t, loadResourceYAML(t, expectedIngressAccessControlListPath), aclOne)
		testutil.AssertEqual(t, loadResourceYAML(t, expectedEgressAccessControlListPath), aclTwo)
	} else if aclOne["display_name"] == "egress-access-control-list" {
		testutil.AssertEqual(t, loadResourceYAML(t, expectedEgressAccessControlListPath), aclOne)
		testutil.AssertEqual(t, loadResourceYAML(t, expectedIngressAccessControlListPath), aclTwo)
	} else {
		assert.Fail(t, "unexpected ACL display_name: %+v", aclOne)
	}
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
