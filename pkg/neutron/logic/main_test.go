package logic_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/format"
	"github.com/Juniper/contrail/pkg/neutron"
	"github.com/Juniper/contrail/pkg/neutron/logic"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

var server *integration.APIServer

func TestMain(m *testing.M) {
	integration.TestMain(m, &server)
}

func loadRequestFromJSONFile(t *testing.T, path string) *neutron.Request {
	var rawJSON map[string]json.RawMessage
	require.NoError(t, fileutil.LoadFile(path, &rawJSON))

	request := &neutron.Request{}
	err := logic.ParseField(rawJSON, "context", &request.Context)
	require.NoError(t, err, "failed to load request. invalid context")

	resource, err := logic.GetResource(request.Context.Type)
	require.NoError(t, err)

	request.Data.Resource = resource
	err = logic.ParseField(rawJSON, "data", &request.Data)
	require.NoError(t, err, "failed to load request. invalid data")

	return request
}

func assertEqual(t *testing.T, expected, actual interface{}) {
	testutil.AssertEqual(t, format.MustYAML(expected), format.MustYAML(actual))
}
