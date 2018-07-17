package compilation_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

const (
	expectedEgressAccessControlList  = "testdata/egress_access_control_list.yml"
	expectedIngressAccessControlList = "testdata/ingress_access_control_list.yml"
	expectedApplicationPolicySet     = "testdata/application_policy_set.yml"
	requestedProjectBlue             = "testdata/requested_project_blue.yml"
	expectedProjectBlue              = "testdata/project_blue.yml"
	expectedSecurityGroup            = "testdata/security_group.yml"
)

func TestIntentCompilationServiceProcessesBasicResourcesCreateEvents(t *testing.T) {
	t.Skip("Not implemented") // TODO: implement compilation service functionality

	closeSync := integration.RunSyncService(t)
	defer closeSync()
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
			s := integration.NewRunningAPIServer(t, "../..", tt.dbDriver)
			defer s.Close(t)
			hc := integration.NewHTTPAPIClient(t, s.URL())

			// Create project section
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

			// Create virtual network with reference to network ipam section
			// TODO: spawn watch on virtual network

			// TODO: spawn watch on routing instance

			// TODO: spawn watch on route target

			// TODO: create virtual network

			// TODO: check virtual network in etcd

			// TODO: check routing instance in etcd

			// TODO: check route target in etcd

			// Create VM-related resources section
			// TODO: create VM-related resources

			// TODO: check all resources in etcd after VM-related creates
		})
	}
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
