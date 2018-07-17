package compilation_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/Juniper/contrail/pkg/convert"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

const (
	stateWithProject        = "testdata/state_with_project.yml"
	stateWithVirtualNetwork = "testdata/state_with_virtual_network.yml"

	expectedEgressAccessControlList  = "testdata/egress_access_control_list.yml"
	expectedIngressAccessControlList = "testdata/ingress_access_control_list.yml"
	expectedApplicationPolicySet     = "testdata/application_policy_set.yml"
	requestedProjectBlue             = "testdata/requested_project_blue.yml"
	expectedDemoProject              = "testdata/demo_project.yml"
	expectedSecurityGroup            = "testdata/security_group.yml"
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

		project := readProject(t, requestedProjectBlue)
		projectWatch, projectCtx, cancelProjectCtx := ec.WatchResource(integration.ProjectSchemaID, project.UUID)
		defer cancelProjectCtx()

		// TODO: spawn watch on application_policy_set

		// TODO: spawn watch on security_group

		// TODO: spawn watch on ingress acl

		// TODO: spawn watch on egress acl

		hc.CreateProject(t, project)
		defer hc.DeleteProject(t, project.UUID)
		defer ec.DeleteProject(t, project.UUID)

		// POST /security-groups
		// PUT /project/950b2912-a742-47c8-acdb-ab361f17386 quota update for subnet
		// PUT /project/950b2912-a742-47c8-acdb-ab361f173867 quota update for vn
		// PUT /project/950b2912-a742-47c8-acdb-ab361f173867 quota update for floating ip
		// PUT /project/950b2912-a742-47c8-acdb-ab361f173867 quota update for security group rule
		// PUT /project/950b2912-a742-47c8-acdb-ab361f173867 quota update for security group
		// PUT /project/950b2912-a742-47c8-acdb-ab361f173867 quota update for logical router
		// PUT /project/950b2912-a742-47c8-acdb-ab361f173867 quota update for VMI

		projectEvent := integration.RetrieveCreateEvent(projectCtx, t, projectWatch)
		// TODO: add defaultDomainUUID to expectedDemoProject
		testutil.AssertEqual(t, readResource(t, expectedDemoProject), decodeJSON(t, projectEvent.Kv.Value))

		// TODO: check application_policy_set in etcd

		// TODO: check security_group in etcd

		// TODO: check ingress acl in etcd

		// TODO: check egress acl in etcd
	}
}

func testCreateVirtualNetworkWithSubnet() func(t *testing.T) {
	return func(t *testing.T) {
		t.Skip("Not implemented") // TODO: implement compilation service functionality

		loadDBSnapshot(t, stateWithProject)
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

		loadDBSnapshot(t, stateWithSubnetTwo)
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

func readResource(t *testing.T, yamlFilePath string) map[string]interface{} {
	var r map[string]interface{}
	err := yaml.Unmarshal(readFile(t, yamlFilePath), &r)
	require.NoError(t, err)
	return r
}

func readProject(t *testing.T, yamlFilePath string) *models.Project {
	var p *models.Project
	err := yaml.Unmarshal(readFile(t, yamlFilePath), p)
	require.NoError(t, err)
	return p
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
