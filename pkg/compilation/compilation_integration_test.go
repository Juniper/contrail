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
	initialState               = "testdata/initial_state.yml"
	stateWithProject           = "testdata/state_with_project.yml"
	stateWithVirtualNetwork    = "testdata/state_with_virtual_network.yml"
	stateWithSubnetOne         = "testdata/state_with_subnet_one.yml"
	stateWithSubnetTwo         = "testdata/state_with_subnet_two.yml"
	stateWithVirtualMachineOne = "testdata/state_with_virtual_machine_one.yml"

	expectedEgressAccessControlList  = "testdata/egress_access_control_list.yml"
	expectedIngressAccessControlList = "testdata/ingress_access_control_list.yml"
	expectedApplicationPolicySet     = "testdata/application_policy_set.yml"
	requestedProjectBlue             = "testdata/requested_project_blue.yml"
	expectedProjectBlue              = "testdata/project_blue.yml"
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

			t.Run("create project", testCreateProject(hc, ec))
			t.Run("create virtual network", testCreateVirtualNetwork())
			t.Run("create subnet", testCreateSubnetOne())
			t.Run("create subnet", testCreateSubnetTwo())
			t.Run("create vm one", testCreateVMOne())
			t.Run("create vm two", testCreateVMTwo())
		})
	}
}

func testCreateProject(hc *integration.HTTPAPIClient, ec *integration.EtcdClient) func(t *testing.T) {
	return func(t *testing.T) {
		t.Skip("Not implemented") // TODO: implement compilation service functionality

		loadDBSnapshot(t, initialState)
		// TODO: defer cleanupDB()

		project := readProject(t, requestedProjectBlue) // TODO: verify requestedProjectBlue data
		projectWatch, projectCtx, cancelProjectCtx := ec.WatchResource(integration.ProjectSchemaID, project.UUID)
		defer cancelProjectCtx()

		// TODO: spawn watch on application_policy_set

		// TODO: spawn watch on security_group

		// TODO: spawn watch on ingress acl

		// TODO: spawn watch on egress acl

		hc.CreateProject(t, project)
		defer hc.DeleteProject(t, project.UUID)
		defer ec.DeleteProject(t, project.UUID)

		projectEvent := integration.RetrieveCreateEvent(projectCtx, t, projectWatch)
		testutil.AssertEqual(t, readResource(t, expectedProjectBlue), decodeJSON(t, projectEvent.Kv.Value))

		// TODO: check application_policy_set in etcd

		// TODO: check security_group in etcd

		// TODO: check ingress acl in etcd

		// TODO: check egress acl in etcd
	}
}

func testCreateVirtualNetwork() func(t *testing.T) {
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

func testCreateSubnetOne() func(t *testing.T) {
	return func(t *testing.T) {
		t.Skip("Not implemented") // TODO: implement compilation service functionality

		loadDBSnapshot(t, stateWithVirtualNetwork)
		// TODO: defer cleanupDB()

		// TODO: spawn watch on subnet1-related resources

		// TODO: create subnet1-related resources

		// TODO: check all resources in etcd after subnet1-related requests
	}
}

func testCreateSubnetTwo() func(t *testing.T) {
	return func(t *testing.T) {
		t.Skip("Not implemented") // TODO: implement compilation service functionality

		loadDBSnapshot(t, stateWithSubnetOne)
		// TODO: defer cleanupDB()

		// TODO: spawn watch on subnet2-related resources

		// TODO: create subnet2-related resources

		// TODO: check all resources in etcd after subnet2-related requests
	}
}

func testCreateVMOne() func(t *testing.T) {
	return func(t *testing.T) {
		t.Skip("Not implemented") // TODO: implement compilation service functionality

		loadDBSnapshot(t, stateWithSubnetTwo)
		// TODO: defer cleanupDB()

		// TODO: spawn watch on VM1-related resources

		// TODO: create VM1-related resources

		// TODO: check all resources in etcd after VM1-related requests
	}
}

func testCreateVMTwo() func(t *testing.T) {
	return func(t *testing.T) {
		t.Skip("Not implemented") // TODO: implement compilation service functionality

		loadDBSnapshot(t, stateWithVirtualMachineOne)
		// TODO: defer cleanupDB()

		// TODO: spawn watch on VM2-related resources

		// TODO: create VM1-related resources

		// TODO: check all resources in etcd after VM1-related requests
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
