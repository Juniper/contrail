package logic_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/db/basedb"
	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/format"
	"github.com/Juniper/contrail/pkg/neutron/logic"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

func newHTTPClient(t *testing.T) (*integration.HTTPAPIClient, func() error) {
	s := integration.NewRunningAPIServer(t, &integration.APIServerConfig{
		DBDriver:           basedb.DriverPostgreSQL,
		RepoRootPath:       "../../..",
		EnableEtcdNotifier: false,
	})
	return integration.NewTestingHTTPClient(t, s.URL()), s.Close
}

func loadRequestFromJSONFile(t *testing.T, path string) *logic.Request {
	r := &logic.Request{}
	require.NoError(t, fileutil.LoadFile(path, r))
	return r
}

func assertEqual(t *testing.T, expected, actual interface{}) {
	testutil.AssertEqual(t, format.MustYAML(expected), format.MustYAML(actual))
}
