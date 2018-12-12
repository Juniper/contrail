package logic_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/format"
	"github.com/Juniper/contrail/pkg/neutron/logic"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

const (
	testDataPath = "./test_data/"
)

func runTest(t *testing.T, test func(*testing.T, *integration.HTTPAPIClient)) {
	for _, driver := range []string{basedb.DriverMySQL, basedb.DriverPostgreSQL} {
		func() {
			s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{
				DBDriver:     driver,
				RepoRootPath: "../../..",
			})
			defer func() {
				assert.NoError(t, s.Close())
			}()

			test(t, integration.NewTestingHTTPClient(t, s.URL()))
		}()
	}
}

func loadRequestFromJSONFile(t *testing.T, path string) *logic.Request {
	r := &logic.Request{}
	require.NoError(t, fileutil.LoadFile(path, r))
	return r
}

func assertEqual(t *testing.T, expected, actual interface{}) {
	testutil.AssertEqual(t, format.MustYAML(expected), format.MustYAML(actual))
}
